package main

import (
	"fmt"

	mynet "github.com/buck119br/psss/net"
	"golang.org/x/sys/unix"
)

var (
	MaxLocalAddrLength  int
	MaxRemoteAddrLength int

	NewlineFlag bool
)

func ShowSummary() {
	var format string
	GenericReadSockstat()
	fmt.Println("Transport\t Total\t IPv4\t IPv6\t")
	for _, pf := range SummaryPF {
		if len(pf) >= 8 {
			format = "%s\t %d\t %d\t %d\t\n"
		} else {
			format = "%s\t\t %d\t %d\t %d\t\n"
		}
		fmt.Printf(format, pf, Summary[pf][IPv4String]+Summary[pf][IPv6String], Summary[pf][IPv4String], Summary[pf][IPv6String])
	}
}

func SocketShowInit() {
	GlobalRecords[GlobalRecordsKey] = make(map[uint32]*GenericRecord)
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}

func SocketShow() {
	if protocalFilter&ProtocalUnix != 0 {
		SocketShowInit()
		UnixRecordRead()
		GenericShow(ProtocalUnix, unix.AF_UNIX)
	}
	if protocalFilter&ProtocalRAW != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		SocketShowInit()
		GenericRecordRead()
		GenericShow(ProtocalRAW, unix.AF_INET)
	}
	if protocalFilter&ProtocalRAW != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		SocketShowInit()
		GenericRecordRead()
		GenericShow(ProtocalRAW, unix.AF_INET6)
	}
	if protocalFilter&ProtocalUDP != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		SocketShowInit()
		GenericRecordRead()
		GenericShow(ProtocalUDP, unix.AF_INET)
	}
	if protocalFilter&ProtocalUDP != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		SocketShowInit()
		GenericRecordRead()
		GenericShow(ProtocalUDP, unix.AF_INET6)
	}
	if protocalFilter&ProtocalTCP != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		SocketShowInit()
		GenericRecordRead()
		GenericShow(ProtocalTCP, unix.AF_INET)
	}
	if protocalFilter&ProtocalTCP != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		SocketShowInit()
		GenericRecordRead()
		GenericShow(ProtocalTCP, unix.AF_INET6)
	}
}

func GenericShow(protocal int, af int) {
	var ok bool
	if *flagProcess {
		GetProcInfo()
	}
	fmt.Printf("Netid\tState\t\tRecv-Q\tSend-Q\t")
	fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, "LocalAddress:Port", MaxRemoteAddrLength, "RemoteAddress:Port")
	if *flagProcess {
		fmt.Printf("Users")
	}
	fmt.Printf("\n")
	for _, record := range GlobalRecords[GlobalRecordsKey] {
		if !*flagAll && !mynet.SstateActive[record.Status] {
			continue
		}
		if *flagListen && !mynet.SstateListen[record.Status] {
			continue
		}
		switch protocal {
		case ProtocalTCP:
			fmt.Printf("tcp\t")
		case ProtocalUDP:
			fmt.Printf("udp\t")
		case ProtocalRAW:
			fmt.Printf("raw\t")
		case ProtocalUnix:
			if _, ok = mynet.SocketType[record.Type]; !ok {
				fmt.Printf("dgr\t")
			} else {
				fmt.Printf("%s\t", mynet.SocketType[record.Type])
			}
		}
		if len(mynet.Sstate[record.Status]) >= 8 {
			fmt.Printf("%s\t", mynet.Sstate[record.Status])
		} else {
			fmt.Printf("%s\t\t", mynet.Sstate[record.Status])
		}
		fmt.Printf("%d\t%d\t", record.RxQueue, record.TxQueue)
		fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, record.LocalAddr.String(), MaxRemoteAddrLength, record.RemoteAddr.String())
		// Process Info
		if *flagProcess && len(record.UserName) > 0 {
			record.ProcInfoPrint()
		}
		if NewlineFlag {
			fmt.Printf("\n")
		}
		if protocal != ProtocalUnix {
			// Timer Info
			if *flagOption && record.Timer != 0 {
				record.TimerInfoPrint()
			}
			// Detailed Info
			if *flagExtended {
				record.ExtendInfoPrint()
			}
		}
		// Meminfo
		if *flagMemory && len(record.Meminfo) == 8 {
			record.MeminfoPrint()
		}
		// internal TCP info
		if *flagInfo && protocal == ProtocalUnix {
			record.TCPInfoPrint()
		}
		fmt.Printf("\n")
	}
}
