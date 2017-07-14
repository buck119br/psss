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

type GenericRecord struct {
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

func NewGenericRecord() *GenericRecord {
	t := new(GenericRecord)
	return t
}

func (record *GenericRecord) Reset() {
	record.LocalAddr.Host = ""
	record.LocalAddr.Port = ""
	record.RemoteAddr.Host = ""
	record.RemoteAddr.Port = ""
	record.Status = 0
	record.TxQueue = 0
	record.RxQueue = 0
	record.Timer = 0
	record.Timeout = 0
	record.Retransmit = 0
	record.UID = 0
	record.Probes = 0
	record.Inode = 0
	record.RefCount = 0
	record.SK = 0
	record.RTO = 0
	record.ATO = 0
	record.QACK = 0
	record.CongestionWindow = 0
	record.SlowStartThreshold = 0
	record.Opt = nil
	record.TCPInfo = nil
	record.VegasInfo = nil
	record.CONG = nil
	record.Drops = 0
	record.Type = 0
	record.Meminfo = nil
	record.UserName = ""
}

func (record *GenericRecord) SetUpRelation() {
	var (
		proc *ProcInfo
		i    int
	)
	for _, procMap := range globalProcInfo {
		for _, proc = range procMap {
			for i = range proc.Fds {
				if record.Inode == proc.Fds[i].Inode {
					record.UserName = proc.Stat.Name
				}
			}
		}
	}
}

func (record *GenericRecord) GenericInfoPrint() {
	if len(Sstate[record.Status]) >= 8 {
		fmt.Printf("%s\t", Sstate[record.Status])
	} else {
		fmt.Printf("%s\t\t", Sstate[record.Status])
	}
	fmt.Printf("%d\t%d\t", record.RxQueue, record.TxQueue)
	fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, record.LocalAddr.String(), MaxRemoteAddrLength, record.RemoteAddr.String())
}

func (record *GenericRecord) ProcInfoPrint() {
	var i int
	fmt.Printf(`["%s":`, record.UserName)
	for pid, proc := range globalProcInfo[record.UserName] {
		for i = range proc.Fds {
			if record.Inode == proc.Fds[i].Inode {
				fmt.Printf(`(pid=%d,fd=%s)`, pid, proc.Fds[i].Name)
			}
		}
	}
	fmt.Printf("]")
}

func (record *GenericRecord) TimerInfoPrint() {
	fmt.Printf("[timer:(%s,%dsec,", TimerState[record.Timer], record.Timeout)
	if record.Timer != 1 {
		fmt.Printf("%d)]    ", record.Probes)
	} else {
		fmt.Printf("%d)]    ", record.Retransmit)
	}
}

func (record *GenericRecord) ExtendInfoPrint() {
	fmt.Printf("[detail:(")
	if record.UID != 0 {
		fmt.Printf("uid:%d,", record.UID)
	}
	fmt.Printf("ino:%d,sk:%x", record.Inode, record.SK)
	if len(record.Opt) > 0 {
		fmt.Printf(",opt:%v", record.Opt)
	}
	fmt.Printf(")]    ")
}

func (record *GenericRecord) MeminfoPrint() {
	fmt.Printf("[skmem:(r:%d,rb:%d,t:%d,tb:%d,f:%d,w:%d,o:%d,bl:%d)]    ",
		record.Meminfo[SK_MEMINFO_RMEM_ALLOC],
		record.Meminfo[SK_MEMINFO_RCVBUF],
		record.Meminfo[SK_MEMINFO_WMEM_ALLOC],
		record.Meminfo[SK_MEMINFO_SNDBUF],
		record.Meminfo[SK_MEMINFO_FWD_ALLOC],
		record.Meminfo[SK_MEMINFO_WMEM_QUEUED],
		record.Meminfo[SK_MEMINFO_OPTMEM],
		record.Meminfo[SK_MEMINFO_BACKLOG])
}

func (record *GenericRecord) TCPInfoPrint() {
	fmt.Printf("[internal:(")
	if record.TCPInfo.Options&TCPI_OPT_TIMESTAMPS != 0 {
		fmt.Printf(" ts")
	}
	if record.TCPInfo.Options&TCPI_OPT_SACK != 0 {
		fmt.Printf(" sack")
	}
	if record.TCPInfo.Options&TCPI_OPT_ECN != 0 {
		fmt.Printf(" ecn")
	}
	if record.TCPInfo.Options&TCPI_OPT_ECN_SEEN != 0 {
		fmt.Printf(" ecnseen")
	}
	if record.TCPInfo.Options&TCPI_OPT_SYN_DATA != 0 {
		fmt.Printf(" fastopen")
	}
	if len(record.CONG) > 1 {
		fmt.Printf(" %s", string(record.CONG))
	}
	if record.TCPInfo.Options&TCPI_OPT_WSCALE != 0 {
		fmt.Printf(" wscale:%d,%d", record.TCPInfo.Pad_cgo_0[0]&0xf, record.TCPInfo.Pad_cgo_0[0]>>4)
	}
	if record.TCPInfo.Rto != 0 && record.TCPInfo.Rto != 3000000 {
		fmt.Printf(" rto:%.2f", float64(record.TCPInfo.Rto)/1000)
	}
	if record.TCPInfo.Backoff != 0 {
		fmt.Printf(" bakcoff:%d", record.TCPInfo.Backoff)
	}
	if record.TCPInfo.Rtt != 0 {
		fmt.Printf(" rtt:%.2f/%.2f", float64(record.TCPInfo.Rtt)/1000, float64(record.TCPInfo.Rttvar)/1000)
	}
	if record.TCPInfo.Ato != 0 {
		fmt.Printf(" ato:%.2f", float64(record.TCPInfo.Ato)/1000)
	}
	if record.QACK != 0 {
		fmt.Printf(" qack:%d", record.QACK)
	}
	if record.QACK&1 != 0 {
		fmt.Printf(" bidir")
	}
	if record.TCPInfo.Snd_mss != 0 {
		fmt.Printf(" mss:%d", record.TCPInfo.Snd_mss)
	}
	if record.TCPInfo.Rcv_mss != 0 {
		fmt.Printf(" rcvmss:%d", record.TCPInfo.Rcv_mss)
	}
	if record.TCPInfo.Advmss != 0 {
		fmt.Printf(" advmss:%d", record.TCPInfo.Advmss)
	}
	if record.TCPInfo.Snd_cwnd != 0 {
		fmt.Printf(" cwnd:%d", record.TCPInfo.Snd_cwnd)
	}
	if record.TCPInfo.Snd_ssthresh < 0xffff {
		fmt.Printf(" ssthresh:%d", record.TCPInfo.Snd_ssthresh)
	}
	if record.TCPInfo.Bytes_acked != 0 {
		fmt.Printf(" bytes_acked:%s", BwToStr(float64(record.TCPInfo.Bytes_acked)))
	}
	if record.TCPInfo.Bytes_received != 0 {
		fmt.Printf(" bytes_received:%s", BwToStr(float64(record.TCPInfo.Bytes_received)))
	}
	if record.TCPInfo.Segs_out != 0 {
		fmt.Printf(" segs_out:%d", record.TCPInfo.Segs_out)
	}
	if record.TCPInfo.Segs_in != 0 {
		fmt.Printf(" segs_in:%d", record.TCPInfo.Segs_in)
	}
	if record.TCPInfo.Data_segs_out != 0 {
		fmt.Printf(" data_segs_out:%d", record.TCPInfo.Data_segs_out)
	}
	if record.TCPInfo.Data_segs_in != 0 {
		fmt.Printf(" data_segs_in:%d", record.TCPInfo.Data_segs_in)
	}

	// DCTCP && BBRInfo

	if record.VegasInfo != nil {
		rtt := record.TCPInfo.Rtt
		if record.VegasInfo.Enabled != 0 && record.VegasInfo.Rtt != 0 && record.VegasInfo.Rtt != 0x7fffffff {
			rtt = record.VegasInfo.Rtt
		}
		if rtt > 0 && record.TCPInfo.Snd_mss != 0 && record.TCPInfo.Snd_cwnd != 0 {
			fmt.Printf(" send:%sbps", BwToStr(float64(record.TCPInfo.Snd_cwnd)*float64(record.TCPInfo.Snd_mss)*8000000/float64(rtt)))
		}
	}

	if record.TCPInfo.Last_data_sent != 0 {
		fmt.Printf(" lastsnd:%d", record.TCPInfo.Last_data_sent)
	}
	if record.TCPInfo.Last_data_recv != 0 {
		fmt.Printf(" lastrcv:%d", record.TCPInfo.Last_data_recv)
	}
	if record.TCPInfo.Last_ack_recv != 0 {
		fmt.Printf(" lastack:%d", record.TCPInfo.Last_ack_recv)
	}
	if record.TCPInfo.Pacing_rate != 0 {
		fmt.Printf(" pacing_rate:%sbps", BwToStr(float64(record.TCPInfo.Pacing_rate*8)))
		if record.TCPInfo.Max_pacing_rate != 0 {
			fmt.Printf("/%sbps", BwToStr(float64(record.TCPInfo.Max_pacing_rate*8)))
		}
	}
	if record.TCPInfo.Delivery_rate != 0 {
		fmt.Printf(" delivery_rate:%sbps", BwToStr(float64(record.TCPInfo.Delivery_rate*8)))
	}
	if record.TCPInfo.Pad_cgo_0[1] != 0 {
		fmt.Printf(" app_limited")
	}
	if record.TCPInfo.Busy_time != 0 {
		fmt.Printf(" busy:%sms", BwToStr(float64(record.TCPInfo.Busy_time/1000)))
	}
	if record.TCPInfo.Rwnd_limited != 0 {
		fmt.Printf(" rwnd_limited:%sms(%.2f%%)",
			BwToStr(float64(record.TCPInfo.Rwnd_limited/1000)),
			100.0*float64(record.TCPInfo.Rwnd_limited)/float64(record.TCPInfo.Busy_time))
	}
	if record.TCPInfo.Sndbuf_limited != 0 {
		fmt.Printf(" sndbuf_limited:%sms(%.2f%%)",
			BwToStr(float64(record.TCPInfo.Sndbuf_limited/1000)),
			100.0*float64(record.TCPInfo.Sndbuf_limited)/float64(record.TCPInfo.Busy_time))
	}
	if record.TCPInfo.Unacked != 0 {
		fmt.Printf(" unacked:%d", record.TCPInfo.Unacked)
	}
	if record.TCPInfo.Retrans != 0 || record.TCPInfo.Total_retrans != 0 {
		fmt.Printf(" retrans:%d/%d", record.TCPInfo.Retrans, record.TCPInfo.Total_retrans)
	}
	if record.TCPInfo.Lost != 0 {
		fmt.Printf(" lost:%d", record.TCPInfo.Lost)
	}
	if record.TCPInfo.Sacked != 0 && record.Status != SsLISTEN {
		fmt.Printf(" sacked:%d", record.TCPInfo.Sacked)
	}
	if record.TCPInfo.Fackets != 0 {
		fmt.Printf(" fackets:%d", record.TCPInfo.Fackets)
	}
	if record.TCPInfo.Reordering != 3 {
		fmt.Printf(" reordering:%d", record.TCPInfo.Reordering)
	}
	if record.TCPInfo.Rcv_rtt != 0 {
		fmt.Printf(" rcv_rtt:%.2f", float64(record.TCPInfo.Rcv_rtt)/1000)
	}
	if record.TCPInfo.Rcv_space != 0 {
		fmt.Printf(" rcv_space:%d", record.TCPInfo.Rcv_space)
	}
	if record.TCPInfo.Notsent_bytes != 0 {
		fmt.Printf(" notsent:%d", record.TCPInfo.Notsent_bytes)
	}
	if record.TCPInfo.Min_rtt != 0 && record.TCPInfo.Min_rtt != math.MaxUint32 {
		fmt.Printf(" minrtt:%s", BwToStr(float64(record.TCPInfo.Min_rtt)/1000))
	}
	fmt.Printf(" )]\n")
}
