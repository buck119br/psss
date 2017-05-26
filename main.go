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
			proc.TCPRecords = append(proc.TCPRecords, tcpRecord)
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
		procRecords map[string][]*TCPRecord
		records     []*TCPRecord
		procName    string
		status      string
		ok          bool
	)
	statusMap := make(map[string]map[string][]*TCPRecord)
	for _, record := range GlobalTCPv4Records {
		status = TCPState[int(record.Status)]
		if procRecords, ok = statusMap[status]; !ok {
			procRecords = make(map[string][]*TCPRecord)
		}
		for _, proc := range record.Procs {
			for _, fd := range proc.Fd {
				if fd.SysStat.Ino == record.Inode {
					procName = proc.Name
					if records, ok = procRecords[procName]; !ok {
						records = make([]*TCPRecord, 0, 0)
					}
					break
				}
			}
		}
		records = append(records, record)
		procRecords[procName] = records
		statusMap[status] = procRecords
	}
	for status, procRecords = range statusMap {
		fmt.Println(status)
		for procName, records = range procRecords {
			fmt.Println("\t", procName)
			for _, v := range records {
				fmt.Printf("\t\t %s\t %s\n", v.LocalAddr, v.RemoteAddr)
			}
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
