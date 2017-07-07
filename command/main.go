package main

import (
	"flag"
	"fmt"

	"github.com/buck119br/psss/psss"
	"golang.org/x/sys/unix"
)

const (
	version = "ss utility, 0.0.1"
	usage   = "Usage:\tss [ OPTIONS ]\n" +
		"\tss [ OPTIONS ] [ FILTER ]\n"
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

	NewlineFlag bool
)

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
		psss.SsFilter = (1 << psss.SsMAX) - 1
	}
	if *flagListen {
		psss.SsFilter = 1<<psss.SsLISTEN | 1<<psss.SsUNCONN
	}
	if psss.SsFilter == 0 {
		psss.SsFilter = 1 << psss.SsESTAB
	}

	if *flagIPv4 {
		psss.AfFilter |= 1 << unix.AF_INET
		if !*flagTCP && !*flagUDP && !*flagRAW {
			*flagTCP = true
			*flagUDP = true
			*flagRAW = true
		}
	}
	if *flagIPv6 {
		psss.AfFilter |= 1 << unix.AF_INET6
		if !*flagTCP && !*flagUDP && !*flagRAW {
			*flagTCP = true
			*flagUDP = true
			*flagRAW = true
		}
	}
	if psss.AfFilter == 0 {
		psss.AfFilter |= 1<<unix.AF_INET | 1<<unix.AF_INET6
	}

	if *flagTCP {
		psss.ProtocalFilter |= psss.ProtocalTCP
	}
	if *flagUDP {
		psss.ProtocalFilter |= psss.ProtocalUDP
	}
	if *flagRAW {
		psss.ProtocalFilter |= psss.ProtocalRAW
	}
	if *flagUnix {
		psss.AfFilter |= 1 << unix.AF_UNIX
		psss.ProtocalFilter |= psss.ProtocalUnix
	}
	if psss.ProtocalFilter == 0 {
		psss.ProtocalFilter |= psss.ProtocalMax - 1
	}

	if *flagInfo {
		psss.FlagInfo = true
	}

	if *flagMemory {
		psss.FlagMemory = true
	}

	if *flagExtended || *flagOption || *flagMemory || *flagInfo {
		NewlineFlag = true
	}

	if *flagProcess {
		var ok bool
		psss.GlobalProcInfo = make(map[string]map[int]*psss.ProcInfo)
		go psss.GetProcInfo()
		psss.ProcInfoInputChan <- psss.NewProcInfo()
		for proc := range psss.ProcInfoOutputChan {
			if proc == nil {
				break
			}
			if proc.Stat.Name == "NULL" {
				goto next
			}
			if _, ok = psss.GlobalProcInfo[proc.Stat.Name]; !ok {
				psss.GlobalProcInfo[proc.Stat.Name] = make(map[int]*psss.ProcInfo)
			}
			psss.GlobalProcInfo[proc.Stat.Name][proc.Stat.Pid] = proc
		next:
			psss.ProcInfoInputChan <- psss.NewProcInfo()
		}
	}

	SocketShow()
}
