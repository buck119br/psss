package main

import (
	"flag"
	"fmt"
)

const (
	version = "ss utility, 0.0.1"
	usage   = "Usage:\tss [ OPTIONS ]\n" +
		"\tss [ OPTIONS ] [ FILTER ]\n"
)

var (
	flagVersion = flag.Bool("v", false, "output version information")
	flagTCP     = flag.Bool("t", false, "display only TCP sockets")
	flagProcess = flag.Bool("p", false, "show process using socket")
	flagSummary = flag.Bool("s", false, "show socket usage summary")
)

func init() {
	GlobalTCPv4Records = make(map[uint64]*TCPRecord)
	GlobalTCPv6Records = make(map[uint64]*TCPRecord)
}

func SetUpRelation() {
	var (
		tcpRecord *TCPRecord
		ok        bool
	)
	for _, proc := range GlobalProcInfo {
		for _, fd := range proc.Fd {
			if fd.SysStat.Dev != 6 {
				continue
			}
			if tcpRecord, ok = GlobalTCPv4Records[fd.SysStat.Ino]; !ok {
				continue
			}
			tcpRecord.Procs = append(tcpRecord.Procs, proc)
			GlobalTCPv4Records[fd.SysStat.Ino] = tcpRecord
		}
	}
}

func main() {
	var err error
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}
	if *flagVersion {
		fmt.Println(version)
		return
	}
	if err = GetTCPRecord(false); err != nil {
		fmt.Println(err)
		return
	}
	if err = GetTCPRecord(true); err != nil {
		fmt.Println(err)
		return
	}
	if err = GetProcInfo(); err != nil {
		fmt.Println(err)
		return
	}
	SetUpRelation()
	switch {
	case *flagSummary:
		ShowUsageSummary()
	case *flagTCP && *flagProcess:
		Show()
	}

}
