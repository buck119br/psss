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

const (
	FbTCPv4 = 1 << iota
	FbTCPv6
	FbUDPv4
	FbUDPv6
	FbRAWv4
	FbRAWv6
	FbUnix
)

var (
	flagHelp    = flag.Bool("h", false, "this message")               // OK
	flagVersion = flag.Bool("v", false, "output version information") // OK

	flagAll        = flag.Bool("a", false, "display all sockets")              // ok
	flagExtended   = flag.Bool("e", false, "show detailed socket information") // ok
	flagInfo       = flag.Bool("i", false, "show internal TCP information")    //
	flagListen     = flag.Bool("l", false, "display listening sockets")        // ok
	flagMemory     = flag.Bool("m", false, "show socket memory usage")         //
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
	flagUnix   = flag.Bool("x", false, "display only Unix domain sockets")  //

	flagDemand = flag.Bool("demand", false, "my boss' demand") // ok

	/* Family bitmap
	31 30 29 28 27 26 25 24 23 22 21 20 19 18 17 16 15 14 13 12 11 10 09 08 07 06 05 04 03 02 01 00
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  TCPv4
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  TCPv6
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  UDPv4
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  UDPv6
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  RAWv4
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  RAWv6
	|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  Unix
	*/
	Family int
)

func init() {
	Summary = make(map[string]map[string]int)
	for _, v := range Protocal {
		Summary[v] = make(map[string]int)
	}

	GlobalTCPv4Records = make(map[uint64]*GenericRecord)
	GlobalTCPv6Records = make(map[uint64]*GenericRecord)
	GlobalUDPv4Records = make(map[uint64]*GenericRecord)
	GlobalUDPv6Records = make(map[uint64]*GenericRecord)
	GlobalRAWv4Records = make(map[uint64]*GenericRecord)
	GlobalRAWv6Records = make(map[uint64]*GenericRecord)
	GlobalUnixRecords = make(map[uint64]*GenericRecord)
}

func dataReader() {
	if Family&FbTCPv4 != 0 {
		GenericRecordRead(TCPv4Str)
	}
	if Family&FbTCPv6 != 0 {
		GenericRecordRead(TCPv6Str)
	}
	if Family&FbUDPv4 != 0 {
		GenericRecordRead(UDPv4Str)
	}
	if Family&FbUDPv6 != 0 {
		GenericRecordRead(UDPv6Str)
	}
	if Family&FbRAWv4 != 0 {
		GenericRecordRead(RAWv4Str)
	}
	if Family&FbRAWv6 != 0 {
		GenericRecordRead(RAWv6Str)
	}
	if Family&FbUnix != 0 {
		UnixRecordRead()
	}
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
	if *flagRAW {
		Family |= FbRAWv4 | FbRAWv6
	}
	if *flagUnix {
		Family |= FbUnix
	}
	if Family == 0 && *flagAll {
		Family |= FbTCPv4 | FbTCPv6 | FbUDPv4 | FbUDPv6 | FbRAWv4 | FbRAWv6 | FbUnix
	}
	dataReader()
	if *flagProcess {
		GetProcInfo()
		SetUpRelation()
	}
	if *flagExtended || *flagOption {
		NewlineFlag = true
	}
	if *flagDemand {
		DemandShow()
		return
	}
	SocketShow()
}
