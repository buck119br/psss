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

const ()

var (
	flagHelp    = flag.Bool("h", false, "this message")
	flagVersion = flag.Bool("v", false, "output version information")

	flagAll        = flag.Bool("a", false, "display all sockets")
	flagExtended   = flag.Bool("e", false, "show detailed socket information")
	flagInfo       = flag.Bool("i", false, "show internal TCP information")
	flagListen     = flag.Bool("l", false, "display listening sockets")
	flagMemory     = flag.Bool("m", false, "show socket memory usage")
	flagNotResolve = flag.Bool("n", false, "don't resolve service names")
	flagOpetion    = flag.Bool("o", false, "show timer information")
	flagProcess    = flag.Bool("p", false, "show process using socket")
	flagResolve    = flag.Bool("r", false, "resolve host names")
	flagSummary    = flag.Bool("s", false, "show socket usage summary")

	flagIPv4   = flag.Bool("4", false, "display only IP version 4 sockets")
	flagIPv6   = flag.Bool("6", false, "display only IP version 6 sockets")
	flagPacket = flag.Bool("0", false, "display PACKET sockets")
	flagDCCP   = flag.Bool("d", false, "display only DCCP sockets")
	flagTCP    = flag.Bool("t", false, "display only TCP sockets")
	flagUDP    = flag.Bool("u", false, "display only UDP sockets")
	flagRAW    = flag.Bool("w", false, "display only RAW sockets")
	flagUNIX   = flag.Bool("x", false, "display only Unix domain sockets")

	flagDemand = flag.Bool("demand", false, "my boss' demand")
)

func init() {
	Summary = make(map[string]map[string]int)
	for _, v := range Protocal {
		Summary[v] = make(map[string]int)
	}

	GlobalTCPv4Records = make(map[uint64]*TCPRecord)
	GlobalTCPv6Records = make(map[uint64]*TCPRecord)
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
		if err = ShowUsageSummary(); err != nil {
			fmt.Println(err)
		}
		return
	}
	if *flagAll {
		*flagPacket = true
		*flagDCCP = true
		*flagTCP = true
		*flagUDP = true
		*flagRAW = true
		*flagUNIX = true
	}
	if *flagProcess {
		if err = GetProcInfo(); err != nil {
			fmt.Println(err)
			return
		}
		SetUpRelation()
	}
}
