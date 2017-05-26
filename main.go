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
	if *flagTCP {
		if _, err := GetTCPRecord(); err != nil {
			panic(err)
		}
		return
	}
}
