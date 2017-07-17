package psss

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	SsUNKNOWN uint8 = iota
	SsESTAB
	SsSYNSENT
	SsSYNRECV
	SsFINWAIT1
	SsFINWAIT2
	SsTIMEWAIT
	SsUNCONN
	SsCLOSEWAIT
	SsLASTACK
	SsLISTEN
	SsCLOSING
	SsMAX
)

const (
	SOCK_STREAM    = 1
	SOCK_DGRAM     = 2
	SOCK_RAW       = 3
	SOCK_RDM       = 4
	SOCK_SEQPACKET = 5
	SOCK_DCCP      = 6
	SOCK_PACKET    = 10
)

const (
	SK_MEMINFO_RMEM_ALLOC = iota
	SK_MEMINFO_RCVBUF
	SK_MEMINFO_WMEM_ALLOC
	SK_MEMINFO_SNDBUF
	SK_MEMINFO_FWD_ALLOC
	SK_MEMINFO_WMEM_QUEUED
	SK_MEMINFO_OPTMEM
	SK_MEMINFO_BACKLOG
	SK_MEMINFO_DROPS
	SK_MEMINFO_VARS
)

const (
	IPv4String = "IPv4"
	IPv6String = "IPv6"
)

var (
	Sstate = []string{
		"UNKNOWN",
		"ESTAB",
		"SYN-SENT",
		"SYN-RECV",
		"FIN-WAIT-1",
		"FIN-WAIT-2",
		"TIME-WAIT",
		"UNCONN",
		"CLOSE-WAIT",
		"LAST-ACK",
		"LISTEN",
		"CLOSING",
		"MAX",
	}

	SocketType = map[uint8]string{
		SOCK_STREAM:    "str",
		SOCK_DGRAM:     "dgr",
		SOCK_RAW:       "raw",
		SOCK_RDM:       "rdm",
		SOCK_SEQPACKET: "seq",
		SOCK_DCCP:      "dccp",
		SOCK_PACKET:    "pack",
	}

	TimerState = []string{
		"OFF",
		"ON",
		"KEEPALIVE",
		"TIMEWAIT",
		"PERSIST",
		"UNKNOWN",
	}

	SummaryPF = []string{
		"TCP",
		"UDP",
		"UDPLITE",
		"RAW",
		"FRAG",
	}

	Colons = []string{
		":::::::",
		"::::::",
		":::::",
		"::::",
		":::",
	}
)

type IP struct {
	Host string
	Port string
}

func (i IP) String() (str string) {
	return i.Host + ":" + i.Port
}

func IPv4HexToString(ipHex string) (ip string, err error) {
	if len(ipHex) != 8 {
		return ip, fmt.Errorf("invalid input:[%s]", ipHex)
	}
	for indexBuffer = 3; indexBuffer > 0; indexBuffer-- {
		if int64Buffer, err = strconv.ParseInt(ipHex[indexBuffer*2:(indexBuffer+1)*2], 16, 64); err != nil {
			return "", err
		}
		ip += fmt.Sprintf("%d", int64Buffer) + "."
	}
	if int64Buffer, err = strconv.ParseInt(ipHex[0:2], 16, 64); err != nil {
		return "", err
	}
	ip += fmt.Sprintf("%d", int64Buffer)
	return ip, nil
}

func IPv6HexToString(ipHex string) (ip string, err error) {
	prefix := ipHex[:24]
	suffix := ipHex[24:]
	for indexBuffer = 0; indexBuffer < 6; indexBuffer++ {
		if prefix[indexBuffer:indexBuffer+4] == "0000" {
			ip += ":"
			continue
		}
		ip += prefix[indexBuffer:indexBuffer+4] + ":"
	}
	for indexBuffer = range Colons {
		ip = strings.Replace(ip, Colons[indexBuffer], "::", -1)
	}
	if suffix, err = IPv4HexToString(suffix); err != nil {
		return "", err
	}
	ip += suffix
	return ip, nil
}

type SocketInfo struct {
	// Generic
	LocalAddr  IP
	RemoteAddr IP
	Status     uint8
	TxQueue    uint32
	RxQueue    uint32
	Timer      int
	Timeout    int
	Retransmit int
	UID        uint64
	Probes     int // unanswered 0-window probes
	Inode      uint32
	RefCount   int
	SK         uint64
	// /proc/net/tcp or /proc/net/tcp6 specific
	RTO                float64  // RetransmitTimeout
	ATO                float64  // Predicted tick of soft clock (delayed ACK control data)
	QACK               int      // (ack.quick<<1)|ack.pingpong
	CongestionWindow   int      // sending congestion window
	SlowStartThreshold int      // slow start size threshold, or -1 if the threshold is >= 0xFFFF
	Opt                []string // Option Info
	// Internal TCP information
	TCPInfo   *TCPInfo
	VegasInfo *TCPVegasInfo
	CONG      []byte
	// Extended Info
	Drops   int   // Generic like UDP, RAW specific
	Type    uint8 // socket type
	Meminfo []uint32
	// Related processes
	UserName string
}

func NewSocketInfo() *SocketInfo {
	t := new(SocketInfo)
	return t
}

func (si *SocketInfo) Reset() {
	si.LocalAddr.Host = ""
	si.LocalAddr.Port = ""
	si.RemoteAddr.Host = ""
	si.RemoteAddr.Port = ""
	si.Status = 0
	si.TxQueue = 0
	si.RxQueue = 0
	si.Timer = 0
	si.Timeout = 0
	si.Retransmit = 0
	si.UID = 0
	si.Probes = 0
	si.Inode = 0
	si.RefCount = 0
	si.SK = 0
	si.RTO = 0
	si.ATO = 0
	si.QACK = 0
	si.CongestionWindow = 0
	si.SlowStartThreshold = 0
	si.Opt = nil
	si.TCPInfo = nil
	si.VegasInfo = nil
	si.CONG = nil
	si.Drops = 0
	si.Type = 0
	si.Meminfo = nil
	si.UserName = ""
}

func (si *SocketInfo) SetUpRelation() {
	var (
		pid   int
		inode uint32
	)
	for name := range GlobalProcFds {
		for pid = range GlobalProcFds[name] {
			for inode = range GlobalProcFds[name][pid] {
				if si.Inode == inode {
					si.UserName = globalProcInfo[name][pid].Stat.Name
					return
				}
			}
		}
	}
}

func (si *SocketInfo) GenericInfoPrint() {
	if len(Sstate[si.Status]) >= 8 {
		fmt.Printf("%s\t", Sstate[si.Status])
	} else {
		fmt.Printf("%s\t\t", Sstate[si.Status])
	}
	fmt.Printf("%d\t%d\t", si.RxQueue, si.TxQueue)
	fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, si.LocalAddr.String(), MaxRemoteAddrLength, si.RemoteAddr.String())
}

func (si *SocketInfo) ProcInfoPrint() {
	var inode uint32
	fmt.Printf(`["%s":`, si.UserName)
	for pid, proc := range globalProcInfo[si.UserName] {
		for inode = range GlobalProcFds[si.UserName][pid] {
			if si.Inode == inode {
				fmt.Printf(`(pid=%d,fd=%s)`, pid, GlobalProcFds[si.UserName][pid][inode].Name)
			}
		}
	}
	fmt.Printf("]")
}

func (si *SocketInfo) TimerInfoPrint() {
	fmt.Printf("[timer:(%s,%dsec,", TimerState[si.Timer], si.Timeout)
	if si.Timer != 1 {
		fmt.Printf("%d)]    ", si.Probes)
	} else {
		fmt.Printf("%d)]    ", si.Retransmit)
	}
}

func (si *SocketInfo) ExtendInfoPrint() {
	fmt.Printf("[detail:(")
	if si.UID != 0 {
		fmt.Printf("uid:%d,", si.UID)
	}
	fmt.Printf("ino:%d,sk:%x", si.Inode, si.SK)
	if len(si.Opt) > 0 {
		fmt.Printf(",opt:%v", si.Opt)
	}
	fmt.Printf(")]    ")
}

func (si *SocketInfo) MeminfoPrint() {
	fmt.Printf("[skmem:(r:%d,rb:%d,t:%d,tb:%d,f:%d,w:%d,o:%d,bl:%d)]    ",
		si.Meminfo[SK_MEMINFO_RMEM_ALLOC],
		si.Meminfo[SK_MEMINFO_RCVBUF],
		si.Meminfo[SK_MEMINFO_WMEM_ALLOC],
		si.Meminfo[SK_MEMINFO_SNDBUF],
		si.Meminfo[SK_MEMINFO_FWD_ALLOC],
		si.Meminfo[SK_MEMINFO_WMEM_QUEUED],
		si.Meminfo[SK_MEMINFO_OPTMEM],
		si.Meminfo[SK_MEMINFO_BACKLOG])
}

func (si *SocketInfo) TCPInfoPrint() {
	fmt.Printf("[internal:(")
	if si.TCPInfo.Options&TCPI_OPT_TIMESTAMPS != 0 {
		fmt.Printf(" ts")
	}
	if si.TCPInfo.Options&TCPI_OPT_SACK != 0 {
		fmt.Printf(" sack")
	}
	if si.TCPInfo.Options&TCPI_OPT_ECN != 0 {
		fmt.Printf(" ecn")
	}
	if si.TCPInfo.Options&TCPI_OPT_ECN_SEEN != 0 {
		fmt.Printf(" ecnseen")
	}
	if si.TCPInfo.Options&TCPI_OPT_SYN_DATA != 0 {
		fmt.Printf(" fastopen")
	}
	if len(si.CONG) > 1 {
		fmt.Printf(" %s", string(si.CONG))
	}
	if si.TCPInfo.Options&TCPI_OPT_WSCALE != 0 {
		fmt.Printf(" wscale:%d,%d", si.TCPInfo.Pad_cgo_0[0]&0xf, si.TCPInfo.Pad_cgo_0[0]>>4)
	}
	if si.TCPInfo.Rto != 0 && si.TCPInfo.Rto != 3000000 {
		fmt.Printf(" rto:%.2f", float64(si.TCPInfo.Rto)/1000)
	}
	if si.TCPInfo.Backoff != 0 {
		fmt.Printf(" bakcoff:%d", si.TCPInfo.Backoff)
	}
	if si.TCPInfo.Rtt != 0 {
		fmt.Printf(" rtt:%.2f/%.2f", float64(si.TCPInfo.Rtt)/1000, float64(si.TCPInfo.Rttvar)/1000)
	}
	if si.TCPInfo.Ato != 0 {
		fmt.Printf(" ato:%.2f", float64(si.TCPInfo.Ato)/1000)
	}
	if si.QACK != 0 {
		fmt.Printf(" qack:%d", si.QACK)
	}
	if si.QACK&1 != 0 {
		fmt.Printf(" bidir")
	}
	if si.TCPInfo.Snd_mss != 0 {
		fmt.Printf(" mss:%d", si.TCPInfo.Snd_mss)
	}
	if si.TCPInfo.Rcv_mss != 0 {
		fmt.Printf(" rcvmss:%d", si.TCPInfo.Rcv_mss)
	}
	if si.TCPInfo.Advmss != 0 {
		fmt.Printf(" advmss:%d", si.TCPInfo.Advmss)
	}
	if si.TCPInfo.Snd_cwnd != 0 {
		fmt.Printf(" cwnd:%d", si.TCPInfo.Snd_cwnd)
	}
	if si.TCPInfo.Snd_ssthresh < 0xffff {
		fmt.Printf(" ssthresh:%d", si.TCPInfo.Snd_ssthresh)
	}
	if si.TCPInfo.Bytes_acked != 0 {
		fmt.Printf(" bytes_acked:%s", BwToStr(float64(si.TCPInfo.Bytes_acked)))
	}
	if si.TCPInfo.Bytes_received != 0 {
		fmt.Printf(" bytes_received:%s", BwToStr(float64(si.TCPInfo.Bytes_received)))
	}
	if si.TCPInfo.Segs_out != 0 {
		fmt.Printf(" segs_out:%d", si.TCPInfo.Segs_out)
	}
	if si.TCPInfo.Segs_in != 0 {
		fmt.Printf(" segs_in:%d", si.TCPInfo.Segs_in)
	}
	if si.TCPInfo.Data_segs_out != 0 {
		fmt.Printf(" data_segs_out:%d", si.TCPInfo.Data_segs_out)
	}
	if si.TCPInfo.Data_segs_in != 0 {
		fmt.Printf(" data_segs_in:%d", si.TCPInfo.Data_segs_in)
	}

	// DCTCP && BBRInfo

	if si.VegasInfo != nil {
		rtt := si.TCPInfo.Rtt
		if si.VegasInfo.Enabled != 0 && si.VegasInfo.Rtt != 0 && si.VegasInfo.Rtt != 0x7fffffff {
			rtt = si.VegasInfo.Rtt
		}
		if rtt > 0 && si.TCPInfo.Snd_mss != 0 && si.TCPInfo.Snd_cwnd != 0 {
			fmt.Printf(" send:%sbps", BwToStr(float64(si.TCPInfo.Snd_cwnd)*float64(si.TCPInfo.Snd_mss)*8000000/float64(rtt)))
		}
	}

	if si.TCPInfo.Last_data_sent != 0 {
		fmt.Printf(" lastsnd:%d", si.TCPInfo.Last_data_sent)
	}
	if si.TCPInfo.Last_data_recv != 0 {
		fmt.Printf(" lastrcv:%d", si.TCPInfo.Last_data_recv)
	}
	if si.TCPInfo.Last_ack_recv != 0 {
		fmt.Printf(" lastack:%d", si.TCPInfo.Last_ack_recv)
	}
	if si.TCPInfo.Pacing_rate != 0 {
		fmt.Printf(" pacing_rate:%sbps", BwToStr(float64(si.TCPInfo.Pacing_rate*8)))
		if si.TCPInfo.Max_pacing_rate != 0 {
			fmt.Printf("/%sbps", BwToStr(float64(si.TCPInfo.Max_pacing_rate*8)))
		}
	}
	if si.TCPInfo.Delivery_rate != 0 {
		fmt.Printf(" delivery_rate:%sbps", BwToStr(float64(si.TCPInfo.Delivery_rate*8)))
	}
	if si.TCPInfo.Pad_cgo_0[1] != 0 {
		fmt.Printf(" app_limited")
	}
	if si.TCPInfo.Busy_time != 0 {
		fmt.Printf(" busy:%sms", BwToStr(float64(si.TCPInfo.Busy_time/1000)))
	}
	if si.TCPInfo.Rwnd_limited != 0 {
		fmt.Printf(" rwnd_limited:%sms(%.2f%%)",
			BwToStr(float64(si.TCPInfo.Rwnd_limited/1000)),
			100.0*float64(si.TCPInfo.Rwnd_limited)/float64(si.TCPInfo.Busy_time))
	}
	if si.TCPInfo.Sndbuf_limited != 0 {
		fmt.Printf(" sndbuf_limited:%sms(%.2f%%)",
			BwToStr(float64(si.TCPInfo.Sndbuf_limited/1000)),
			100.0*float64(si.TCPInfo.Sndbuf_limited)/float64(si.TCPInfo.Busy_time))
	}
	if si.TCPInfo.Unacked != 0 {
		fmt.Printf(" unacked:%d", si.TCPInfo.Unacked)
	}
	if si.TCPInfo.Retrans != 0 || si.TCPInfo.Total_retrans != 0 {
		fmt.Printf(" retrans:%d/%d", si.TCPInfo.Retrans, si.TCPInfo.Total_retrans)
	}
	if si.TCPInfo.Lost != 0 {
		fmt.Printf(" lost:%d", si.TCPInfo.Lost)
	}
	if si.TCPInfo.Sacked != 0 && si.Status != SsLISTEN {
		fmt.Printf(" sacked:%d", si.TCPInfo.Sacked)
	}
	if si.TCPInfo.Fackets != 0 {
		fmt.Printf(" fackets:%d", si.TCPInfo.Fackets)
	}
	if si.TCPInfo.Reordering != 3 {
		fmt.Printf(" reordering:%d", si.TCPInfo.Reordering)
	}
	if si.TCPInfo.Rcv_rtt != 0 {
		fmt.Printf(" rcv_rtt:%.2f", float64(si.TCPInfo.Rcv_rtt)/1000)
	}
	if si.TCPInfo.Rcv_space != 0 {
		fmt.Printf(" rcv_space:%d", si.TCPInfo.Rcv_space)
	}
	if si.TCPInfo.Notsent_bytes != 0 {
		fmt.Printf(" notsent:%d", si.TCPInfo.Notsent_bytes)
	}
	if si.TCPInfo.Min_rtt != 0 && si.TCPInfo.Min_rtt != math.MaxUint32 {
		fmt.Printf(" minrtt:%s", BwToStr(float64(si.TCPInfo.Min_rtt)/1000))
	}
	fmt.Printf(" )]\n")
}
