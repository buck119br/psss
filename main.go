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

func ShowUsageSummary() (err error) {
	var format string
	summary := make(map[string]map[string]int)
	for _, v := range Protocal {
		summary[v] = make(map[string]int)
		summary[v][IPv4String] = 0
	}
	summary["TCP"][IPv4String] = len(GlobalTCPv4Records)
	fmt.Println("Transport\t Total\t IPv4\t IPv6\t")
	for _, protocal := range Protocal {
		if len(protocal) >= 8 {
			format = "%s\t %d\t %d\t %d\t\n"
		} else {
			format = "%s\t\t %d\t %d\t %d\t\n"
		}
		fmt.Printf(format, protocal, summary[protocal][IPv4String]+0, summary[protocal][IPv4String], 0)

	}
	return nil
}

func Show() {
	var (
		records []*TCPRecord
		ok      bool
	)
	statusMap := make(map[string][]*TCPRecord)
	for _, v := range GlobalTCPv4Records {
		if records, ok = statusMap[TCPState[int(v.Status)]]; !ok {
			records = make([]*TCPRecord, 0, 0)
		}
		records = append(records, v)
		statusMap[TCPState[int(v.Status)]] = records
	}
	for status, records := range statusMap {
		fmt.Println(status)
		for _, record := range records {
			fmt.Printf("\t %s\t %s\t users:(", record.LocalAddr, record.RemoteAddr)
			procSet := make(map[string]bool)
			for _, v := range record.Procs {
				for _, fd := range v.Fd {
					if fd.SysStat.Ino == record.Inode {
						procSet[v.Name] = true
						break
					}
				}
			}
			for k := range procSet {
				fmt.Printf(`"%s"`, k)
			}
			fmt.Printf(")\n")
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
	if err = GetTCPv4Record(); err != nil {
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

func testfunc() {
}
