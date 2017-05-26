package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TCPPath = "/proc/net/tcp"
)

var TCPState = []string{
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

type TCPRecord struct {
	LocalAddr         string
	RemoteAddr        string
	Status            int64
	TxQueue           int64
	RxQueue           int64
	Timer             int64
	TmWhen            int64
	Retransmit        int64
	UID               int
	Timeout           int
	Inode             int
	RefCount          int
	MemLocation       uint64
	RetransmitTimeout int
	PredictedTick     int
	ACK               int
	CongestionWindow  int
	SlowStartSize     int
}

func GetTCPRecord() (tcpRecords map[int]*TCPRecord, err error) {
	var (
		line        string
		fields      []string
		fieldsIndex int
		stringBuff  []string
		tempInt64   int64
	)
	file, err := os.Open(TCPPath)
	if err != nil {
		return
	}
	defer file.Close()
	tcpRecords = make(map[int]*TCPRecord)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tcpRecord := new(TCPRecord)
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
		if tcpRecord.LocalAddr, err = HexToIP(stringBuff[0]); err != nil {
			continue
		}
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			continue
		}
		tcpRecord.LocalAddr += ":" + fmt.Sprintf("%d", tempInt64)
		fieldsIndex++
		// Remote address
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tcpRecord.RemoteAddr, err = HexToIP(stringBuff[0]); err != nil {
			continue
		}
		if tempInt64, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			continue
		}
		tcpRecord.RemoteAddr += ":" + fmt.Sprintf("%d", tempInt64)
		fieldsIndex++
		// Status
		if tcpRecord.Status, err = strconv.ParseInt(fields[fieldsIndex], 16, 64); err != nil {
			continue
		}
		fieldsIndex++
		// TxQueue:RxQueue
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tcpRecord.TxQueue, err = strconv.ParseInt(stringBuff[0], 16, 64); err != nil {
			continue
		}
		if tcpRecord.RxQueue, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			continue
		}
		fieldsIndex++
		// Timer:TmWhen
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if tcpRecord.Timer, err = strconv.ParseInt(stringBuff[0], 16, 64); err != nil {
			continue
		}
		if tcpRecord.TmWhen, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			continue
		}
		fieldsIndex++
		// Retransmit
		if tcpRecord.Retransmit, err = strconv.ParseInt(fields[fieldsIndex], 16, 64); err != nil {
			continue
		}
		fieldsIndex++
		// UID
		if tcpRecord.UID, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			continue
		}
		fieldsIndex++
		// Timeout
		if tcpRecord.Timeout, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			continue
		}
		fieldsIndex++
		// Inode
		if tcpRecord.Inode, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			continue
		}
		fieldsIndex++
		// Socket reference count
		if tcpRecord.RefCount, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
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
				continue
			}
			fieldsIndex++
			// Predicted tick of soft clock (delayed ACK control data)
			if tcpRecord.PredictedTick, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				continue
			}
			fieldsIndex++
			// 	(ack.quick<<1)|ack.pingpong
			if tcpRecord.ACK, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				continue
			}
			fieldsIndex++
			// 	sending congestion window
			if tcpRecord.CongestionWindow, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				continue
			}
			fieldsIndex++
			// 	slow start size threshold, or -1 if the threshold is >= 0xFFFF
			if tcpRecord.SlowStartSize, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				continue
			}
		}
		tcpRecords[tcpRecord.Inode] = tcpRecord
	}
	return tcpRecords, nil
}

func HexToIP(ipHex string) (ip string, err error) {
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
