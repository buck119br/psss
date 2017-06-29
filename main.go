package main

import (
	"flag"
	"fmt"

	"golang.org/x/sys/unix"
)

const (
	version = "ss utility, 0.0.1"
	usage   = "Usage:\tss [ OPTIONS ]\n" +
		"\tss [ OPTIONS ] [ FILTER ]\n"
)

const (
	ProtocalUnknown = 1 << iota
	ProtocalDCCP
	ProtocalNetlink
	ProtocalPacket
	ProtocalRAW
	ProtocalSCTP
	ProtocalTCP
	ProtocalUDP
	ProtocalUnix
	ProtocalMax
)

var (
	flagHelp    = flag.Bool("h", false, "help message")               // ok
	flagVersion = flag.Bool("v", false, "output version information") // ok

	flagAll    = flag.Bool("a", false, "display all sockets")       // ok
	flagListen = flag.Bool("l", false, "display listening sockets") // ok

	flagExtended   = flag.Bool("e", false, "show detailed socket information") // ok
	flagInfo       = flag.Bool("i", false, "show internal TCP information")    // ok
	flagMemory     = flag.Bool("m", false, "show socket memory usage")         // ok
	flagNotResolve = flag.Bool("n", false, "don't resolve service names")      // born to be
	flagOption     = flag.Bool("o", false, "show timer information")           // ok
	flagProcess    = flag.Bool("p", false, "show process using socket")        // ok
	flagResolve    = flag.Bool("r", false, "resolve host names")               //
	flagSummary    = flag.Bool("s", false, "show socket usage summary")        // ok

	flagIPv4   = flag.Bool("4", false, "display only IP version 4 sockets") // ok
	flagIPv6   = flag.Bool("6", false, "display only IP version 6 sockets") // ok
	flagPacket = flag.Bool("0", false, "display PACKET sockets")            //
	flagDCCP   = flag.Bool("d", false, "display only DCCP sockets")         //
	flagSCTP   = flag.Bool("S", false, "display only SCTP sockets")         //
	flagTCP    = flag.Bool("t", false, "display only TCP sockets")          // ok
	flagUDP    = flag.Bool("u", false, "display only UDP sockets")          // ok
	flagRAW    = flag.Bool("w", false, "display only RAW sockets")          // ok
	flagUnix   = flag.Bool("x", false, "display only Unix domain sockets")  // ok

	flagDemand = flag.Bool("demand", false, "my boss' demand") // ok

	afFilter       uint64
	protocalFilter uint64
	ssFilter       uint32
)

func init() {
	Summary = make(map[string]map[string]int)
	for _, pf := range SummaryPF {
		Summary[pf] = make(map[string]int)
	}

	GlobalRecords = make(map[string]map[uint32]*GenericRecord)
	GlobalProcInfo = make(map[string]map[int]*ProcInfo)
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 || *flagHelp {
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}
	if *flagVersion {
		fmt.Println(version)
		return
	}
	if *flagSummary {
		ShowSummary()
		return
	}
	// sock state
	if *flagAll {
		ssFilter = (1 << SsMAX) - 1
	}
	if *flagListen {
		ssFilter = 1<<SsLISTEN | 1<<SsUNCONN
	}
	if ssFilter == 0 {
		ssFilter = 1 << SsESTAB
	}

	if *flagIPv4 {
		afFilter |= 1 << unix.AF_INET
		if !*flagTCP && !*flagUDP && !*flagRAW {
			*flagTCP = true
			*flagUDP = true
			*flagRAW = true
		}
	}
	if *flagIPv6 {
		afFilter |= 1 << unix.AF_INET6
		if !*flagTCP && !*flagUDP && !*flagRAW {
			*flagTCP = true
			*flagUDP = true
			*flagRAW = true
		}
	}
	if afFilter == 0 {
		afFilter |= 1<<unix.AF_INET | 1<<unix.AF_INET6
	}

	if *flagTCP {
		protocalFilter |= ProtocalTCP
	}
	if *flagUDP {
		protocalFilter |= ProtocalUDP
	}
	if *flagRAW {
		protocalFilter |= ProtocalRAW
	}
	if *flagUnix {
		afFilter |= 1 << unix.AF_UNIX
		protocalFilter |= ProtocalUnix
	}
	if protocalFilter == 0 {
		protocalFilter |= ProtocalMax - 1
	}

	if *flagExtended || *flagOption || *flagMemory || *flagInfo {
		NewlineFlag = true
	}

	if *flagProcess {
		GetProcInfo()
	}

	if *flagDemand {
		// DemandShow()
		return
	}
	SocketShow()
}
