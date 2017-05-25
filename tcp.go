package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type TCPRecordRaw struct {
	LocalAddr  string
	RemoteAddr string
	Status     int
	TxQ        string
	RxQ        string
	tr         string
	tmWhen     string
	retrnsmt   string
	UID        int
	Timeout    int64
	Inode      int64
}

func TCPRecordRead() (tcpRecord []TCPRecordRaw, err error) {
	var (
		line   string
		fields []string
	)
	file, err := os.Open("/Users/liuchunxu/workspace/temp/tcp")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return
		}
		line = scanner.Text()
		fields = strings.Fields(line)
		fmt.Println(fields)
	}
	return tcpRecord, nil
}
