package main

import (
	"fmt"

	"github.com/buck119br/psss/psss"
	"golang.org/x/sys/unix"
)

func ShowSummary() {
	var format string
	summary, err := psss.GenericReadSockstat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Transport\t Total\t IPv4\t IPv6\t")
	for _, pf := range psss.SummaryPF {
		if len(pf) >= 8 {
			format = "%s\t %d\t %d\t %d\t\n"
		} else {
			format = "%s\t\t %d\t %d\t %d\t\n"
		}
		fmt.Printf(format, pf, summary[pf][psss.IPv4String]+summary[pf][psss.IPv6String], summary[pf][psss.IPv4String], summary[pf][psss.IPv6String])
	}
}

func SocketShow() {
	if psss.ProtocalFilter&psss.ProtocalUnix != 0 {
		psss.AddrLengthInit()
		sis, _ = psss.GenericUnixRead()
		GenericShow(psss.ProtocalUnix, unix.AF_UNIX)
	}
	if psss.ProtocalFilter&psss.ProtocalRAW != 0 && psss.AfFilter&(1<<unix.AF_INET) != 0 {
		psss.AddrLengthInit()
		sis, _ = psss.GenericInetRead(psss.ProtocalRAW, unix.AF_INET)
		GenericShow(psss.ProtocalRAW, unix.AF_INET)
	}
	if psss.ProtocalFilter&psss.ProtocalRAW != 0 && psss.AfFilter&(1<<unix.AF_INET6) != 0 {
		psss.AddrLengthInit()
		sis, _ = psss.GenericInetRead(psss.ProtocalRAW, unix.AF_INET6)
		GenericShow(psss.ProtocalRAW, unix.AF_INET6)
	}
	if psss.ProtocalFilter&psss.ProtocalUDP != 0 && psss.AfFilter&(1<<unix.AF_INET) != 0 {
		psss.AddrLengthInit()
		sis, _ = psss.GenericInetRead(psss.ProtocalUDP, unix.AF_INET)
		GenericShow(psss.ProtocalUDP, unix.AF_INET)
	}
	if psss.ProtocalFilter&psss.ProtocalUDP != 0 && psss.AfFilter&(1<<unix.AF_INET6) != 0 {
		psss.AddrLengthInit()
		sis, _ = psss.GenericInetRead(psss.ProtocalUDP, unix.AF_INET6)
		GenericShow(psss.ProtocalUDP, unix.AF_INET6)
	}
	if psss.ProtocalFilter&psss.ProtocalTCP != 0 && psss.AfFilter&(1<<unix.AF_INET) != 0 {
		psss.AddrLengthInit()
		sis, _ = psss.GenericInetRead(psss.ProtocalTCP, unix.AF_INET)
		GenericShow(psss.ProtocalTCP, unix.AF_INET)
	}
	if psss.ProtocalFilter&psss.ProtocalTCP != 0 && psss.AfFilter&(1<<unix.AF_INET6) != 0 {
		psss.AddrLengthInit()
		sis, _ = psss.GenericInetRead(psss.ProtocalTCP, unix.AF_INET6)
		GenericShow(psss.ProtocalTCP, unix.AF_INET6)
	}

	for i := range psss.GlobalProcFds {
		fmt.Println(i)
		for j := range psss.GlobalProcFds[i] {
			fmt.Println("\t", j, psss.GlobalProcFds[i][j])
		}
	}
}

func GenericShow(protocal, af int) {
	if len(sis) == 0 {
		return
	}
	fmt.Printf("Netid\tState\t\tRecv-Q\tSend-Q\t")
	fmt.Printf("%-*s\t%-*s\t", psss.MaxLocalAddrLength, "LocalAddress:Port", psss.MaxRemoteAddrLength, "RemoteAddress:Port")
	if *flagProcess {
		fmt.Printf("Users")
	}
	fmt.Printf("\n")
	var ok bool
	for _, si := range sis {
		switch protocal {
		case psss.ProtocalTCP:
			fmt.Printf("tcp")
		case psss.ProtocalUDP:
			fmt.Printf("udp")
		case psss.ProtocalRAW:
			fmt.Printf("raw")
		case psss.ProtocalUnix:
			if _, ok = psss.SocketType[si.Type]; !ok {
				fmt.Printf("dgr\t")
			} else {
				fmt.Printf("%s\t", psss.SocketType[si.Type])
			}
		}
		switch af {
		case unix.AF_INET:
			fmt.Printf("4\t")
		case unix.AF_INET6:
			fmt.Printf("6\t")
		}
		si.GenericInfoPrint()
		if *flagProcess && len(si.UserName) > 0 {
			si.ProcInfoPrint()
		}
		if newlineFlag {
			fmt.Printf("\n")
		}
		if protocal != psss.ProtocalUnix {
			if *flagOption && si.Timer != 0 {
				si.TimerInfoPrint()
			}
			if *flagExtended {
				si.ExtendInfoPrint()
			}
		}
		if *flagMemory && len(si.Meminfo) == 8 {
			si.MeminfoPrint()
		}
		if *flagInfo && protocal == psss.ProtocalTCP && si.TCPInfo != nil {
			si.TCPInfoPrint()
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
