package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TCPv4Path = "/proc/net/tcp"
	TCPv6Path = "/proc/net/tcp6"

	IPv4String = "IPv4"
	IPv6String = "IPv6"
)

var (
	GlobalTCPv4Records map[uint64]*TCPRecord
	GlobalTCPv6Records map[uint64]*TCPRecord

	Protocal = []string{
		"RAW",
		"UDP",
		"TCP",
		"INET",
		"FRAG",
	}

	TCPState = []string{
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
)

type IP struct {
	Host string
	Port string
}

func (i IP) String() (str string) {
	return i.Host + ":" + i.Port
}

type TCPRecord struct {
	LocalAddr         IP
	RemoteAddr        IP
	Status            int64
	TxQueue           int64
	RxQueue           int64
	Timer             int64
	TmWhen            int64
	Retransmit        int64
	UID               int
	Timeout           int
	Inode             uint64
	RefCount          int
	MemLocation       uint64
	RetransmitTimeout int
	PredictedTick     int
	ACK               int
	CongestionWindow  int
	SlowStartSize     int
	Procs             []*ProcInfo
}

func NewTCPRecord() *TCPRecord {
	t := new(TCPRecord)
	t.Procs = make([]*ProcInfo, 0, 0)
	return t
}

func IPv4HexToString(ipHex string) (ip string, err error) {
	var tempInt int64
	if len(ipHex) != 8 {
		return ip, fmt.Errorf("invalid input:[%s]", ipHex)
	}
	for i := 3; i > 0; i-- {
		if tempInt, err = strconv.ParseInt(ipHex[i*2:(i+1)*2], 16, 64); err != nil {
			return "", err
		}
		ip += fmt.Sprintf("%d", tempInt) + "."
	}
	if tempInt, err = strconv.ParseInt(ipHex[0:2], 16, 64); err != nil {
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
	if suffix, err = IPv4HexToString(suffix); err != nil {
		return "", err
	}
	ip += suffix
	return ip, nil
}

// IPv6:versionFlag = true; IPv4:versionFlag = false
func GetTCPRecord(versionFlag bool) (err error) {
	var (
		file        *os.File
		line        string
		fields      []string
		fieldsIndex int
		stringBuff  []string
		tempInt64   int64
	)
	if versionFlag {
		file, err = os.Open(TCPv6Path)
	} else {
		file, err = os.Open(TCPv4Path)
	}
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tcpRecord := NewTCPRecord()
		if err = scanner.Err(); err != nil {
			return
		}
		line = scanner.Text()
		fields = strings.Fields(line)
		if fields[0] == "sl" {
			continue
		}
		fieldsIndex = 1
		// Local address
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if versionFlag {
			tcpRecord.LocalAddr.Host, err = IPv6HexToString(stringBuff[0])
		} else {
			tcpRecord.LocalAddr.Host, err = IPv4HexToString(stringBuff[0])
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		tcpRecord.LocalAddr.Port = fmt.Sprintf("%d", tempInt64)
		fieldsIndex++
		// Remote address
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if versionFlag {
			tcpRecord.RemoteAddr.Host, err = IPv6HexToString(stringBuff[0])
		} else {
			tcpRecord.RemoteAddr.Host, err = IPv4HexToString(stringBuff[0])
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		tcpRecord.RemoteAddr.Port = fmt.Sprintf("%d", tempInt64)
		fieldsIndex++
		// Status
		if tcpRecord.Status, err = strconv.ParseInt(fields[fieldsIndex], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// TxQueue:RxQueue
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tcpRecord.TxQueue, err = strconv.ParseInt(stringBuff[0], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		if tcpRecord.RxQueue, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// Timer:TmWhen
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tcpRecord.Timer, err = strconv.ParseInt(stringBuff[0], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		if tcpRecord.TmWhen, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// Retransmit
		if tcpRecord.Retransmit, err = strconv.ParseInt(fields[fieldsIndex], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// UID
		if tcpRecord.UID, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// Timeout
		if tcpRecord.Timeout, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// Inode
		if tcpRecord.Inode, err = strconv.ParseUint(fields[fieldsIndex], 10, 64); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// Socket reference count
		if tcpRecord.RefCount, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			fmt.Println(err)
			continue
		}
		fieldsIndex++
		// Socket memory location
		if tcpRecord.MemLocation, err = strconv.ParseUint(fields[fieldsIndex], 16, 64); err != nil {
			fmt.Println(err)
			continue
		}
		if tcpRecord.Inode != 0 {
			fieldsIndex++
			// Retransmit timeout
			if tcpRecord.RetransmitTimeout, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				fmt.Println(err)
				continue
			}
			fieldsIndex++
			// Predicted tick of soft clock (delayed ACK control data)
			if tcpRecord.PredictedTick, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				fmt.Println(err)
				continue
			}
			fieldsIndex++
			// 	(ack.quick<<1)|ack.pingpong
			if tcpRecord.ACK, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				fmt.Println(err)
				continue
			}
			fieldsIndex++
			// 	sending congestion window
			if tcpRecord.CongestionWindow, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				fmt.Println(err)
				continue
			}
			fieldsIndex++
			// 	slow start size threshold, or -1 if the threshold is >= 0xFFFF
			if tcpRecord.SlowStartSize, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				fmt.Println(err)
				continue
			}
		}
		if versionFlag {
			GlobalTCPv6Records[tcpRecord.Inode] = tcpRecord
			continue
		}
		GlobalTCPv4Records[tcpRecord.Inode] = tcpRecord
	}
	return nil
}
