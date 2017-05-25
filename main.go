package main

import (
	"flag"
	"fmt"
)

const (
	ProcRoot = "/proc"
	version  = "ss utility, 0.0.1"
	usage    = "Usage:\tss [ OPTIONS ]\n" +
		"\tss [ OPTIONS ] [ FILTER ]\n"
)

var (
	flagTCP     = flag.Bool("t", false, "display only TCP sockets")
	flagVersion = flag.Bool("v", false, "output version information")
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
	if *flagTCP {
		if _, err := TCPRecordRead(); err != nil {
			panic(err)
		}
		return
	}
}
