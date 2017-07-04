package main

import (
	"fmt"

	"github.com/buck119br/psss/psss"
	"golang.org/x/sys/unix"
)

func ShowSummary() {
	var format string
	psss.GenericReadSockstat()
	fmt.Println("Transport\t Total\t IPv4\t IPv6\t")
	for _, pf := range psss.SummaryPF {
		if len(pf) >= 8 {
			format = "%s\t %d\t %d\t %d\t\n"
		} else {
			format = "%s\t\t %d\t %d\t %d\t\n"
		}
		fmt.Printf(format, pf, psss.Summary[pf][psss.IPv4String]+psss.Summary[pf][psss.IPv6String], psss.Summary[pf][psss.IPv4String], psss.Summary[pf][psss.IPv6String])
	}
}

func SocketShow() {
	if protocalFilter&ProtocalUnix != 0 {
		psss.AddrLengthInit()
		psss.GlobalRecords[psss.GlobalRecordsKey] = psss.UnixRecordRead()
		GenericShow(ProtocalUnix, unix.AF_UNIX)
	}
	if protocalFilter&ProtocalRAW != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		psss.AddrLengthInit()
		psss.GlobalRecords[psss.GlobalRecordsKey] = psss.GenericRecordRead(ProtocalRAW, unix.AF_INET)
		GenericShow(ProtocalRAW, unix.AF_INET)
	}
	if protocalFilter&ProtocalRAW != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		psss.AddrLengthInit()
		psss.GlobalRecords[psss.GlobalRecordsKey] = psss.GenericRecordRead(ProtocalRAW, unix.AF_INET6)
		GenericShow(ProtocalRAW, unix.AF_INET6)
	}
	if protocalFilter&ProtocalUDP != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		psss.AddrLengthInit()
		psss.GlobalRecords[psss.GlobalRecordsKey] = psss.GenericRecordRead(ProtocalUDP, unix.AF_INET)
		GenericShow(ProtocalUDP, unix.AF_INET)
	}
	if protocalFilter&ProtocalUDP != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		psss.AddrLengthInit()
		psss.GlobalRecords[psss.GlobalRecordsKey] = psss.GenericRecordRead(ProtocalUDP, unix.AF_INET6)
		GenericShow(ProtocalUDP, unix.AF_INET6)
	}
	if protocalFilter&ProtocalTCP != 0 && afFilter&(1<<unix.AF_INET) != 0 {
		psss.AddrLengthInit()
		psss.GlobalRecords[psss.GlobalRecordsKey] = psss.GenericRecordRead(ProtocalTCP, unix.AF_INET)
		GenericShow(ProtocalTCP, unix.AF_INET)
	}
	if protocalFilter&ProtocalTCP != 0 && afFilter&(1<<unix.AF_INET6) != 0 {
		psss.AddrLengthInit()
		psss.GlobalRecords[psss.GlobalRecordsKey] = psss.GenericRecordRead(ProtocalTCP, unix.AF_INET6)
		GenericShow(ProtocalTCP, unix.AF_INET6)
	}
}

func GenericShow(protocal, af int) {
	if len(psss.GlobalRecords[psss.GlobalRecordsKey]) == 0 {
		return
	}
	if *flagProcess {
		SetUpRelation()
	}
	fmt.Printf("Netid\tState\t\tRecv-Q\tSend-Q\t")
	fmt.Printf("%-*s\t%-*s\t", psss.MaxLocalAddrLength, "LocalAddress:Port", psss.MaxRemoteAddrLength, "RemoteAddress:Port")
	if *flagProcess {
		fmt.Printf("Users")
	}
	fmt.Printf("\n")
	var ok bool
	for _, record := range psss.GlobalRecords[psss.GlobalRecordsKey] {
		switch protocal {
		case ProtocalTCP:
			fmt.Printf("tcp")
		case ProtocalUDP:
			fmt.Printf("udp")
		case ProtocalRAW:
			fmt.Printf("raw")
		case ProtocalUnix:
			if _, ok = SocketType[record.Type]; !ok {
				fmt.Printf("dgr\t")
			} else {
				fmt.Printf("%s\t", SocketType[record.Type])
			}
		}
		switch af {
		case unix.AF_INET:
			fmt.Printf("4\t")
		case unix.AF_INET6:
			fmt.Printf("6\t")
		}
		record.GenericInfoPrint()
		if *flagProcess && len(record.UserName) > 0 {
			record.ProcInfoPrint()
		}
		if NewlineFlag {
			fmt.Printf("\n")
		}
		if protocal != ProtocalUnix {
			if *flagOption && record.Timer != 0 {
				record.TimerInfoPrint()
			}
			if *flagExtended {
				record.ExtendInfoPrint()
			}
		}
		if *flagMemory && len(record.Meminfo) == 8 {
			record.MeminfoPrint()
		}
		if *flagInfo && protocal == ProtocalTCP {
			record.TCPInfoPrint()
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}