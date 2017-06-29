package main

import (
	"fmt"

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
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}

func SocketShow() {
	if protocalFilter&ProtocalUnix != 0 {
		SocketShowInit()
		GlobalRecords[GlobalRecordsKey] = UnixRecordRead()
		GenericShow(ProtocalUnix, unix.AF_UNIX)
	}
	if protocalFilter&ProtocalRAW != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		SocketShowInit()
		GlobalRecords[GlobalRecordsKey] = GenericRecordRead(ProtocalRAW, unix.AF_INET)
		GenericShow(ProtocalRAW, unix.AF_INET)
	}
	if protocalFilter&ProtocalRAW != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		SocketShowInit()
		GlobalRecords[GlobalRecordsKey] = GenericRecordRead(ProtocalRAW, unix.AF_INET6)
		GenericShow(ProtocalRAW, unix.AF_INET6)
	}
	if protocalFilter&ProtocalUDP != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		SocketShowInit()
		GlobalRecords[GlobalRecordsKey] = GenericRecordRead(ProtocalUDP, unix.AF_INET)
		GenericShow(ProtocalUDP, unix.AF_INET)
	}
	if protocalFilter&ProtocalUDP != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		SocketShowInit()
		GlobalRecords[GlobalRecordsKey] = GenericRecordRead(ProtocalUDP, unix.AF_INET6)
		GenericShow(ProtocalUDP, unix.AF_INET6)
	}
	if protocalFilter&ProtocalTCP != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		SocketShowInit()
		GlobalRecords[GlobalRecordsKey] = GenericRecordRead(ProtocalTCP, unix.AF_INET)
		GenericShow(ProtocalTCP, unix.AF_INET)
	}
	if protocalFilter&ProtocalTCP != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		SocketShowInit()
		GlobalRecords[GlobalRecordsKey] = GenericRecordRead(ProtocalTCP, unix.AF_INET6)
		GenericShow(ProtocalTCP, unix.AF_INET6)
	}
}

func GenericShow(protocal, af int) {
	if len(GlobalRecords[GlobalRecordsKey]) == 0 {
		return
	}
	if *flagProcess {
		SetUpRelation()
	}
	fmt.Printf("Netid\tState\t\tRecv-Q\tSend-Q\t")
	fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, "LocalAddress:Port", MaxRemoteAddrLength, "RemoteAddress:Port")
	if *flagProcess {
		fmt.Printf("Users")
	}
	fmt.Printf("\n")
	var ok bool
	for _, record := range GlobalRecords[GlobalRecordsKey] {
		switch protocal {
		case ProtocalTCP:
			fmt.Printf("tcp\t")
		case ProtocalUDP:
			fmt.Printf("udp\t")
		case ProtocalRAW:
			fmt.Printf("raw\t")
		case ProtocalUnix:
			if _, ok = SocketType[record.Type]; !ok {
				fmt.Printf("dgr\t")
			} else {
				fmt.Printf("%s\t", SocketType[record.Type])
			}
		}
		if len(Sstate[record.Status]) >= 8 {
			fmt.Printf("%s\t", Sstate[record.Status])
		} else {
			fmt.Printf("%s\t\t", Sstate[record.Status])
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
		if *flagInfo && protocal == ProtocalTCP {
			record.TCPInfoPrint()
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
