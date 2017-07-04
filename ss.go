package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"encoding/json"
	"golang.org/x/sys/unix"
)

const (
	GlobalRecordsKey = "GlobalRecords"

	UserHz = 3

	IPv4String = "IPv4"
	IPv6String = "IPv6"
)

var (
	Summary   map[string]map[string]int
	SummaryPF = []string{
		"TCP",
		"UDP",
		"UDPLITE",
		"RAW",
		"FRAG",
	}

	procFilePath = map[string]string{
		"sockstat4": "/proc/net/sockstat",
		"sockstat6": "/proc/net/sockstat6",
		"TCP4":      "/proc/net/tcp",
		"TCP6":      "/proc/net/tcp6",
		"UDP4":      "/proc/net/udp",
		"UDP6":      "/proc/net/udp6",
		"RAW4":      "/proc/net/raw",
		"RAW6":      "/proc/net/raw6",
		"Unix":      "/proc/net/unix",
	}

	TimerName = []string{
		"OFF",
		"ON",
		"KEEPALIVE",
		"TIMEWAIT",
		"PERSIST",
		"UNKNOWN",
	}

	Colons = []string{
		":::::::",
		"::::::",
		":::::",
		"::::",
		":::",
	}

	GlobalRecords map[string]map[uint32]*GenericRecord
)

type IP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (i IP) String() (str string) {
	return i.Host + ":" + i.Port
}

type IPset map[IP]bool

func NewIPset() IPset {
	return make(map[IP]bool)
}

func (i IPset) MarshalJSON() ([]byte, error) {
	temp := make([]string, 0)
	for ip := range i {
		temp = append(temp, ip.String())
	}
	return json.Marshal(temp)
}

type IPCounter map[IP]int

func NewIPCounter() IPCounter {
	return make(map[IP]int)
}

func (i IPCounter) MarshalJSON() ([]byte, error) {
	tempset := make(map[string]int)
	for ip, count := range i {
		tempset[ip.String()] = count
	}
	return json.Marshal(tempset)
}

func IPv4HexToString(ipHex string) (ip string, err error) {
	var tempInt int64
	if len(ipHex) != 8 {
		fmt.Printf("invalid input:[%s]\n", ipHex)
		return ip, fmt.Errorf("invalid input:[%s]", ipHex)
	}
	for i := 3; i > 0; i-- {
		if tempInt, err = strconv.ParseInt(ipHex[i*2:(i+1)*2], 16, 64); err != nil {
			fmt.Println(err)
			return "", err
		}
		ip += fmt.Sprintf("%d", tempInt) + "."
	}
	if tempInt, err = strconv.ParseInt(ipHex[0:2], 16, 64); err != nil {
		fmt.Println(err)
		return "", err
	}
	ip += fmt.Sprintf("%d", tempInt)
	return ip, nil
}

func IPv6HexToString(ipHex string) (ip string, err error) {
	prefix := ipHex[:24]
	suffix := ipHex[24:]
	for i := 0; i < 6; i++ {
		if prefix[i:i+4] == "0000" {
			ip += ":"
			continue
		}
		ip += prefix[i:i+4] + ":"
	}
	for _, v := range Colons {
		ip = strings.Replace(ip, v, "::", -1)
	}
	if suffix, err = IPv4HexToString(suffix); err != nil {
		fmt.Println(err)
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
	TCPInfo   TCPInfo
	VegasInfo TCPVegasInfo
	CONG      []byte
	// Extended Info
	Drops   int   // Generic like UDP, RAW specific
	Type    uint8 // socket type
	Meminfo []uint32
	// Related processes
	Procs    map[*ProcInfo]bool
	UserName string
}

func NewGenericRecord() GenericRecord {
	t := new(GenericRecord)
	t.Procs = make(map[*ProcInfo]bool)
	return *t
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
	fmt.Printf(`["%s":`, record.UserName)
	for proc := range record.Procs {
		for _, fd := range proc.Fd {
			if fd.SysStat.Ino == uint64(record.Inode) {
				fmt.Printf(`(pid=%d,fd=%s)`, proc.Stat.Pid, fd.Name)
			}
		}
	}
	fmt.Printf("]")
}
func (record *GenericRecord) TimerInfoPrint() {
	fmt.Printf("[timer:(%s,%dsec,", TimerName[record.Timer], record.Timeout)
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

	rtt := record.TCPInfo.Rtt
	if record.VegasInfo.Enabled != 0 && record.VegasInfo.Rtt != 0 && record.VegasInfo.Rtt != 0x7fffffff {
		rtt = record.VegasInfo.Rtt
	}
	if rtt > 0 && record.TCPInfo.Snd_mss != 0 && record.TCPInfo.Snd_cwnd != 0 {
		fmt.Printf(" send:%sbps", BwToStr(float64(record.TCPInfo.Snd_cwnd)*float64(record.TCPInfo.Snd_mss)*8000000/float64(rtt)))
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

func UnixRecordRead() (records map[uint32]*GenericRecord) {
	skfd, err := SendUnixDiagMsg(ssFilter,
		UDIAG_SHOW_NAME|UDIAG_SHOW_VFS|UDIAG_SHOW_PEER|UDIAG_SHOW_ICONS|UDIAG_SHOW_RQLEN|UDIAG_SHOW_MEMINFO)
	if err != nil {
		goto readProc
	}
	defer unix.Close(skfd)
	return RecvUnixDiagMsgAll(skfd)

readProc:
	// In this way, so much information cannot get.
	var (
		line        string
		fields      []string
		fieldsIndex int
		tempInt64   int64
		flag        int64
	)
	records = make(map[uint32]*GenericRecord)
	file, err := os.Open(procFilePath["Unix"])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
		line = scanner.Text()
		fields = strings.Fields(line)
		if len(fields) < 7 {
			continue
		}
		if fields[0] == "Num" {
			continue
		}
		record := NewGenericRecord()
		// Num: the kernel table slot number.
		fieldsIndex = 0
		if record.SK, err = strconv.ParseUint(strings.Replace(fields[fieldsIndex], ":", "", -1), 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.RemoteAddr.Host = "*"
		record.RemoteAddr.Port = "Unknown"
		if MaxRemoteAddrLength < len(record.RemoteAddr.String()) {
			MaxRemoteAddrLength = len(record.RemoteAddr.String())
		}
		fieldsIndex++
		// RefCount: the number of users of the socket.
		record.RxQueue = 0
		fieldsIndex++
		// Protocol: currently always 0.
		fieldsIndex++
		// Flags: the internal kernel flags holding the status of the socket.
		if flag, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// Type: the socket type.
		// For SOCK_STREAM sockets, this is 0001; for SOCK_DGRAM sockets, it is 0002; and for SOCK_SEQPACKET sockets, it is 0005.
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.Type = uint8(tempInt64)
		fieldsIndex++
		// St: the internal state of the socket.
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		if flag&(1<<16) != 0 {
			record.Status = SsLISTEN
		} else {
			record.Status = UnixSstate[int(tempInt64)-1]
		}
		if ssFilter&(1<<record.Status) == 0 {
			continue
		}
		fieldsIndex++
		// Inode
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 10, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.Inode = uint32(tempInt64)
		record.LocalAddr.Port = fmt.Sprintf("%d", record.Inode)
		// Path: the bound path (if any) of the socket.
		// Sockets in the abstract namespace are included in the list, and are shown with a Path that commences with the character '@'.
		if len(fields) > 7 {
			fieldsIndex++
			record.LocalAddr.Host = fields[fieldsIndex]
		} else {
			record.LocalAddr.Host = "*"
		}
		if MaxLocalAddrLength < len(record.LocalAddr.String()) {
			MaxLocalAddrLength = len(record.LocalAddr.String())
		}
		records[record.Inode] = &record
	}
	return
}

func GenericRecordRead(protocal, af int) (records map[uint32]*GenericRecord) {
	var (
		ipproto uint8
		exts    uint8
		skfd    int
		err     error
	)
	switch protocal {
	case ProtocalTCP:
		ipproto = unix.IPPROTO_TCP
		if *flagInfo {
			exts |= 1 << (INET_DIAG_INFO - 1)
			exts |= 1 << (INET_DIAG_VEGASINFO - 1)
			exts |= 1 << (INET_DIAG_CONG - 1)
		}
	case ProtocalUDP:
		ipproto = unix.IPPROTO_UDP
	case ProtocalRAW:
		ipproto = unix.IPPROTO_RAW
	default:
		fmt.Println("invalid protocal")
		return
	}
	if *flagMemory {
		exts |= 1 << (INET_DIAG_SKMEMINFO - 1)
	}
	if skfd, err = SendInetDiagMsg(uint8(af), ipproto, exts, ssFilter); err != nil {
		goto readProc
	}
	defer unix.Close(skfd)
	return RecvInetDiagMsgAll(skfd)

readProc:
	var (
		procPath    string
		file        *os.File
		line        string
		fields      []string
		fieldsIndex int
		stringBuff  []string
		tempInt64   int64
	)
	records = make(map[uint32]*GenericRecord)

	switch protocal {
	case ProtocalTCP:
		procPath = "TCP"
	case ProtocalUDP:
		procPath = "UDP"
	case ProtocalRAW:
		procPath = "RAW"
	}
	switch af {
	case unix.AF_INET:
		procPath += "4"
	case unix.AF_INET6:
		procPath += "6"
	}
	if file, err = os.Open(procFilePath[procPath]); err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
		line = scanner.Text()
		fields = strings.Fields(line)
		if fields[0] == "sl" {
			continue
		}
		record := NewGenericRecord()
		// Local address
		fieldsIndex = 1
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		switch af {
		case unix.AF_INET:
			record.LocalAddr.Host, err = IPv4HexToString(stringBuff[0])
		case unix.AF_INET6:
			record.LocalAddr.Host, err = IPv6HexToString(stringBuff[0])
		}
		if err != nil {
			continue
		}
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.LocalAddr.Port = fmt.Sprintf("%d", tempInt64)
		if MaxLocalAddrLength < len(record.LocalAddr.String()) {
			MaxLocalAddrLength = len(record.LocalAddr.String())
		}
		fieldsIndex++
		// Remote address
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		switch af {
		case unix.AF_INET:
			record.RemoteAddr.Host, err = IPv4HexToString(stringBuff[0])
		case unix.AF_INET6:
			record.RemoteAddr.Host, err = IPv6HexToString(stringBuff[0])
		}
		if err != nil {
			continue
		}
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.RemoteAddr.Port = fmt.Sprintf("%d", tempInt64)
		if MaxRemoteAddrLength < len(record.RemoteAddr.String()) {
			MaxRemoteAddrLength = len(record.RemoteAddr.String())
		}
		fieldsIndex++
		// Status
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.Status = uint8(tempInt64)
		if ssFilter&(1<<record.Status) == 0 {
			continue
		}
		fieldsIndex++
		// TxQueue:RxQueue
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tempInt64, err = strconv.ParseInt(stringBuff[0], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.TxQueue = uint32(tempInt64)
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.RxQueue = uint32(tempInt64)
		fieldsIndex++
		// Timer:TmWhen
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tempInt64, err = strconv.ParseInt(stringBuff[0], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.Timer = int(tempInt64)
		if record.Timer > 4 {
			record.Timer = 5
		}
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.Timeout = int(tempInt64)
		record.Timeout = (record.Timeout*1000 + UserHz - 1) / UserHz
		fieldsIndex++
		// Retransmit
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.Retransmit = int(tempInt64)
		fieldsIndex++
		if record.UID, err = strconv.ParseUint(fields[fieldsIndex], 10, 64); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		if record.Probes, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 10, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.Inode = uint32(tempInt64)
		fieldsIndex++
		if record.RefCount, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		if record.SK, err = strconv.ParseUint(fields[fieldsIndex], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		switch protocal {
		case ProtocalTCP:
			if len(fields) > 12 {
				fieldsIndex++
				if record.RTO, err = strconv.ParseFloat(fields[fieldsIndex], 64); err != nil {
					fmt.Println(err)
					continue
				}
				fieldsIndex++
				if record.ATO, err = strconv.ParseFloat(fields[fieldsIndex], 64); err != nil {
					fmt.Println(err)
					continue
				}
				record.ATO = record.ATO / float64(UserHz)
				fieldsIndex++
				if record.QACK, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
					fmt.Println(err)
					continue
				}
				record.QACK = record.QACK / 2
				fieldsIndex++
				if record.CongestionWindow, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
					fmt.Println(err)
					continue
				}
				fieldsIndex++
				if record.SlowStartThreshold, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				record.RTO = 0
				record.ATO = 0
				record.QACK = 0
				record.CongestionWindow = 2
				record.SlowStartThreshold = -1
			}
			if record.SlowStartThreshold == -1 {
				record.SlowStartThreshold = 0
			}
			if record.RTO == float64(3*UserHz) {
				record.RTO = 0
			}
			if record.Timer != 1 {
				record.Retransmit = record.Probes
			}
		case ProtocalUDP, ProtocalRAW:
			fieldsIndex++
			if record.Drops, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				fmt.Println(err)
				continue
			}
		}
		if len(fields) > 17 {
			record.Opt = fields[17:]
		}
		records[record.Inode] = &record
	}
	return
}

func GetSocketCount(fields []string) (int, error) {
	for i := range fields {
		if fields[i] == "inuse" {
			return strconv.Atoi(fields[i+1])
		}
	}
	return 0, nil
}

// IPv6:versionFlag = true; IPv4:versionFlag = false
func GenericReadSockstat() {
	var (
		file   *os.File
		line   string
		fields []string
		err    error
	)

	for _, v := range []string{"sockstat4", "sockstat6"} {
		if file, err = os.Open(procFilePath[v]); err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if err = scanner.Err(); err != nil {
				fmt.Println(err)
				return
			}
			line = scanner.Text()
			fields = strings.Fields(line)
			switch fields[0] {
			case "sockets:":
				continue
			case "TCP:":
				Summary["TCP"][IPv4String], err = GetSocketCount(fields[1:])
			case "TCP6:":
				Summary["TCP"][IPv6String], err = GetSocketCount(fields[1:])
			case "UDP:":
				Summary["UDP"][IPv4String], err = GetSocketCount(fields[1:])
			case "UDP6:":
				Summary["UDP"][IPv6String], err = GetSocketCount(fields[1:])
			case "UDPLITE:":
				Summary["UDPLITE"][IPv4String], err = GetSocketCount(fields[1:])
			case "UDPLITE6:":
				Summary["UDPLITE"][IPv6String], err = GetSocketCount(fields[1:])
			case "RAW:":
				Summary["RAW"][IPv4String], err = GetSocketCount(fields[1:])
			case "RAW6:":
				Summary["RAW"][IPv6String], err = GetSocketCount(fields[1:])
			case "FRAG:":
				Summary["FRAG"][IPv4String], err = GetSocketCount(fields[1:])
			case "FRAG6:":
				Summary["FRAG"][IPv6String], err = GetSocketCount(fields[1:])
			default:
				continue
			}
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
