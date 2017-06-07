package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UserHz = 3

	IPv4String = "IPv4"
	IPv6String = "IPv6"

	Sockstat4Path = "/proc/net/sockstat"
	Sockstat6Path = "/proc/net/sockstat6"
	TCPv4Path     = "/proc/net/tcp"
	TCPv6Path     = "/proc/net/tcp6"
	UDPv4Path     = "/proc/net/udp"
	UDPv6Path     = "/proc/net/udp6"
	RAWv4Path     = "/proc/net/raw"
	RAWv6Path     = "/proc/net/raw6"
	UnixPath      = "/proc/net/unix"

	TCPv4Str = "TCPv4"
	TCPv6Str = "TCPv6"
	UDPv4Str = "UDPv4"
	UDPv6Str = "UDPv6"
	RAWv4Str = "RAWv4"
	RAWv6Str = "RAWv6"
	UnixStr  = "Unix"
)

var (
	Summary map[string]map[string]int

	GlobalTCPv4Records map[uint64]*GenericRecord
	GlobalTCPv6Records map[uint64]*GenericRecord
	GlobalUDPv4Records map[uint64]*GenericRecord
	GlobalUDPv6Records map[uint64]*GenericRecord
	GlobalRAWv4Records map[uint64]*GenericRecord
	GlobalRAWv6Records map[uint64]*GenericRecord
	GlobalUnixRecords  map[uint64]*GenericRecord

	Protocal = []string{
		"RAW",
		"UDP",
		"TCP",
		"FRAG",
	}

	UnixSstate     = []int{7, 2, 1, 11}
	UnixSocketType = map[int]string{
		1: "u_str",
		2: "u_dgr",
		5: "u_seq",
	}

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
	}
	SstateActive = map[int]bool{
		0:  false,
		1:  true,
		2:  false,
		3:  false,
		4:  false,
		5:  false,
		6:  false,
		7:  false,
		8:  false,
		9:  false,
		10: true,
		11: false,
	}
	SstateListen = map[int]bool{
		0:  false,
		1:  false,
		2:  false,
		3:  false,
		4:  false,
		5:  false,
		6:  false,
		7:  true,
		8:  false,
		9:  false,
		10: true,
		11: false,
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
)

type IP struct {
	Host string
	Port string
}

func (i IP) String() (str string) {
	return i.Host + ":" + i.Port
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
	LocalAddr  IP
	RemoteAddr IP
	Status     int
	TxQueue    int
	RxQueue    int
	Timer      int
	Timeout    int
	Retransmit int
	UID        uint64
	Probes     int // unanswered 0-window probes
	Inode      uint64
	RefCount   int
	SK         uint64
	// TCP specific
	RTO                float64 // RetransmitTimeout
	ATO                float64 // Predicted tick of soft clock (delayed ACK control data)
	QACK               int     // (ack.quick<<1)|ack.pingpong
	CongestionWindow   int     // sending congestion window
	SlowStartThreshold int     // slow start size threshold, or -1 if the threshold is >= 0xFFFF
	// Generic like UDP, RAW specific
	Drops int
	// UNIX socket specific
	Type int
	// Option Info
	Opt []string
	// Related processes
	Procs    map[*ProcInfo]bool
	UserName string
}

func NewGenericRecord() *GenericRecord {
	t := new(GenericRecord)
	t.Procs = make(map[*ProcInfo]bool)
	return t
}

func UnixRecordRead() {
	var (
		line        string
		fields      []string
		fieldsIndex int
		tempInt64   int64
		flag        int64
	)
	file, err := os.Open(UnixPath)
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
		record.RemoteAddr.Port = fmt.Sprintf("%x", record.SK)
		fieldsIndex++
		// RefCount: the number of users of the socket.
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.RxQueue = int(tempInt64)
		fieldsIndex++
		// Protocol: currently always 0.
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.TxQueue = int(tempInt64)
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
		record.Type = int(tempInt64)
		fieldsIndex++
		// St: the internal state of the socket.
		if tempInt64, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		if flag&(1<<16) != 0 {
			record.Status = 10 // LISTEN
		} else {
			record.Status = UnixSstate[int(tempInt64)-1]
			if record.Type == 2 && record.Status == 7 && len(record.RemoteAddr.Port) != 0 {
				record.Status = 1
			}
		}
		fieldsIndex++
		// Inode
		if record.Inode, err = strconv.ParseUint(fields[fieldsIndex], 10, 64); err != nil {
			fmt.Println(err)
			continue
		}
		record.LocalAddr.Port = fmt.Sprintf("%d", record.Inode)
		// Path: the bound path (if any) of the socket.
		// Sockets in the abstract namespace are included in the list, and are shown with a Path that commences with the character '@'.
		if len(fields) > 7 {
			fieldsIndex++
			record.LocalAddr.Host = fields[fieldsIndex]

		} else {
			record.LocalAddr.Host = "*"
		}
		GlobalUnixRecords[record.Inode] = record
	}
}

func GenericRecordRead(family string) (err error) {
	var (
		file        *os.File
		line        string
		fields      []string
		fieldsIndex int
		stringBuff  []string
		tempInt64   int64
	)
	switch family {
	case TCPv4Str:
		file, err = os.Open(TCPv4Path)
	case TCPv6Str:
		file, err = os.Open(TCPv6Path)
	case UDPv4Str:
		file, err = os.Open(UDPv4Path)
	case UDPv6Str:
		file, err = os.Open(UDPv6Path)
	case RAWv4Str:
		file, err = os.Open(RAWv4Path)
	case RAWv6Str:
		file, err = os.Open(RAWv6Path)
	default:
		err = fmt.Errorf("invalid family string.")
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
			return err
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
		switch family {
		case TCPv4Str, UDPv4Str, RAWv4Str:
			record.LocalAddr.Host, err = IPv4HexToString(stringBuff[0])
		case TCPv6Str, UDPv6Str, RAWv6Str:
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
		switch family {
		case TCPv4Str, UDPv4Str, RAWv4Str:
			record.RemoteAddr.Host, err = IPv4HexToString(stringBuff[0])
		case TCPv6Str, UDPv6Str, RAWv6Str:
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
		record.Status = int(tempInt64)
		fieldsIndex++
		// TxQueue:RxQueue
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tempInt64, err = strconv.ParseInt(stringBuff[0], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.TxQueue = int(tempInt64)
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 32); err != nil {
			fmt.Println(err)
			continue
		}
		record.RxQueue = int(tempInt64)
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
		if record.Inode, err = strconv.ParseUint(fields[fieldsIndex], 10, 64); err != nil {
			fmt.Println(err)
			continue
		}
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
		switch family {
		case TCPv4Str, TCPv6Str:
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
		case UDPv4Str, UDPv6Str, RAWv4Str, RAWv6Str:
			fieldsIndex++
			if record.Drops, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				fmt.Println(err)
				continue
			}
		}
		if len(fields) > 17 {
			record.Opt = fields[17:]
		}
		switch family {
		case TCPv4Str:
			GlobalTCPv4Records[record.Inode] = record
		case TCPv6Str:
			GlobalTCPv6Records[record.Inode] = record
		case UDPv4Str:
			GlobalUDPv4Records[record.Inode] = record
		case UDPv6Str:
			GlobalUDPv6Records[record.Inode] = record
		case RAWv4Str:
			GlobalRAWv4Records[record.Inode] = record
		case RAWv6Str:
			GlobalRAWv6Records[record.Inode] = record
		}
	}
	return nil
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
func GenericReadSockstat(versionFlag bool) (err error) {
	var (
		file      *os.File
		line      string
		fields    []string
		tempCount int
	)
	if versionFlag {
		file, err = os.Open(Sockstat6Path)
	} else {
		file, err = os.Open(Sockstat4Path)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
			return err
		}
		line = scanner.Text()
		fields = strings.Fields(line)
		switch fields[0] {
		case "sockets:":
			continue
		case "TCP:":
			if Summary["TCP"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		case "TCP6:":
			if Summary["TCP"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		case "UDP:":
			if Summary["UDP"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		case "UDP6:":
			if Summary["UDP"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		case "UDPLITE:":
			if tempCount, err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
			Summary["UDP"][IPv4String] += tempCount
		case "UDPLITE6:":
			if tempCount, err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
			Summary["UDP"][IPv6String] += tempCount
		case "RAW:":
			if Summary["RAW"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		case "RAW6:":
			if Summary["RAW"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		case "FRAG:":
			if Summary["FRAG"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		case "FRAG6:":
			if Summary["FRAG"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				fmt.Println(err)
				return err
			}
		default:
			continue
		}
	}
	return nil
}
