package main

import (
	"flag"
	"fmt"
)

const (
	version = "ss utility, 0.0.1"
	usage   = "Usage:\tss [ OPTIONS ]\n" +
		"\tss [ OPTIONS ] [ FILTER ]\n"

	IPv4String = "IPv4"
	IPv6String = "IPv6"
)

const (
	FbTCPv4 = 1 << iota
	FbTCPv6
	FbUDPv4
	FbUDPv6
)

var (
	flagHelp    = flag.Bool("h", false, "this message")               // OK
	flagVersion = flag.Bool("v", false, "output version information") // OK

	flagAll        = flag.Bool("a", false, "display all sockets")
	flagExtended   = flag.Bool("e", false, "show detailed socket information")
	flagInfo       = flag.Bool("i", false, "show internal TCP information")
	flagListen     = flag.Bool("l", false, "display listening sockets")
	flagMemory     = flag.Bool("m", false, "show socket memory usage")
	flagNotResolve = flag.Bool("n", false, "don't resolve service names")
	flagOpetion    = flag.Bool("o", false, "show timer information")
	flagProcess    = flag.Bool("p", false, "show process using socket")
	flagResolve    = flag.Bool("r", false, "resolve host names")
	flagSummary    = flag.Bool("s", false, "show socket usage summary") // OK

	flagIPv4   = flag.Bool("4", false, "display only IP version 4 sockets")
	flagIPv6   = flag.Bool("6", false, "display only IP version 6 sockets")
	flagPacket = flag.Bool("0", false, "display PACKET sockets")
	flagDCCP   = flag.Bool("d", false, "display only DCCP sockets")
	flagTCP    = flag.Bool("t", false, "display only TCP sockets")
	flagUDP    = flag.Bool("u", false, "display only UDP sockets")
	flagRAW    = flag.Bool("w", false, "display only RAW sockets")
	flagUNIX   = flag.Bool("x", false, "display only Unix domain sockets")

	flagDemand = flag.Bool("demand", false, "my boss' demand")

	/* Family bitmap
	31 30 29 28 27 26 25 24 23 22 21 20 19 18 17 16 15 14 13 12 11 10 09 08 07 06 05 04 03 02 01 00
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  TCPv4
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  TCPv6
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  UDPv4
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  UDPv6
	*/
	Family int
)

func init() {
	Summary = make(map[string]map[string]int)
	for _, v := range Protocal {
		Summary[v] = make(map[string]int)
	}

	GlobalTCPv4Records = make(map[uint64]*TCPRecord)
	GlobalTCPv6Records = make(map[uint64]*TCPRecord)
}

func dataReader() (err error) {
	if Family|FbTCPv4 != 0 {
		if err = GenericReadTCP(false); err != nil {
			fmt.Println(err)
			return
		}
	}
	if Family|FbTCPv6 != 0 {
		if err = GenericReadTCP(true); err != nil {
			fmt.Println(err)
			return
		}
	}
	return nil
}

func main() {
	var err error
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
		if err = ShowSummary(); err != nil {
			fmt.Println(err)
		}
		return
	}
	if *flagAll {
		Family |= FbTCPv4 | FbTCPv6 | FbUDPv4 | FbUDPv6
	}
	if *flagIPv4 {
		Family |= FbTCPv4 | FbUDPv4
	}
	if *flagIPv6 {
		Family |= FbTCPv6 | FbUDPv6
	}
	if *flagTCP {
		Family |= FbTCPv4 | FbTCPv6
	}
	if *flagUDP {
		Family |= FbUDPv4 | FbUDPv6
	}
	// if *flagDemand {
	// 	*flagTCP = true
	// 	*flagProcess = true
	// }
	if err = dataReader(); err != nil {
		return
	}
	if *flagProcess {
		if err = GetProcInfo(); err != nil {
			fmt.Println(err)
			return
		}
		SetUpRelation()
	}
	SocketShow()
	if *flagDemand {
		Show()
		return
	}
}
