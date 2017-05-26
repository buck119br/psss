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
)

func main() {
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
	if *flagTCP && *flagProcess {
		testfunc()
	}
	if *flagTCP {
		if _, err := GetTCPRecord(); err != nil {
			panic(err)
		}
		return
	}
}

func testfunc() {
	var (
		tcpRecord *TCPRecord
		ok        bool
	)
	tcpRecords, err := GetTCPRecord()
	if err != nil {
		fmt.Println(err)
	}
	procs, err := GetProcInfo()
	if err != nil {
		fmt.Println(err)
	}
	for _, proc := range procs {
		for _, fd := range proc.Fd {
			if fd.SysStat.Dev != 6 {
				continue
			}
			if tcpRecord, ok = tcpRecords[fd.SysStat.Ino]; !ok {
				continue
			}
			proc.TCPSockets = append(proc.TCPSockets, tcpRecord)
		}
	}
	for _, proc := range procs {
		fmt.Println(proc.Name)
		for _, socket := range proc.TCPSockets {
			fmt.Println("\t", socket.LocalAddr, socket.RemoteAddr, TCPState[int(socket.Status)])
		}
	}
}
