package main

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/sys/unix"
)

type demand struct {
	Listen map[string]map[IP]bool
	Estab  map[string]map[bool]map[*GenericRecord]bool
}

func newdemand() *demand {
	d := new(demand)
	d.Listen = make(map[string]map[IP]bool)
	d.Estab = make(map[string]map[bool]map[*GenericRecord]bool)
	return d
}

func (d *demand) data() {
	netAddrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddrs := make(map[string]bool)
	localAddrs["127.0.1.1"] = true
	localAddrs["0.0.0.0"] = true
	localAddrs["::0.0.0.0"] = true
	for _, v := range netAddrs {
		localAddrs[strings.Split(v.String(), "/")[0]] = true
	}

	GlocalRecords = make(map[string]map[uint32]*GenericRecord)
	GlobalRecords["4"] = GenericRecordRead(ProtocalTCP, unix.AF_INET)
	GlobalRecords["6"] = GenericRecordRead(ProtocalTCP, unix.AF_INET)
	SetUpRelation()

	var ok, isLocal bool
	for key, records := range GlobalRecords {
		for ino, record := range records {
			switch record.Status {
			case SsLISTEN:
				if _, ok = d.Listen[record.UserName]; !ok {
					d.Listen[record.UserName] = make(map[IP]bool)
				}
				d.Listen[record.UserName][record.LocalAddr] = true
			case SsESTAB:
				if _, ok = d.Estab[record.UserName]; !ok {
					d.Estab[record.UserName] = make(map[bool]map[*GenericRecord]bool)
				}
				_, isLocal = localAddrs[record.RemoteAddr.Host]
				if _, ok = d.Estab[record.UserName][isLocal]; !ok {
					d.Estab[record.UserName][isLocal] = make(map[*GenericRecord]bool)
				}
				d.Estab[record.UserName][isLocal][record] = true
			}
		}
	}
}

func (d *demand) show() {
	d.data()
	fmt.Println("Listen")
	for name, ipmap := range d.Listen {
		fmt.Println("\t", name)
		for ip := range ipmap {
			fmt.Println("\t\t", ip.String())
		}
	}
	fmt.Println("Estab")
	var ok bool
	for name, procmap := range d.Estab {
		fmt.Println("\t", name)
		localService := make(map[string]bool)
		for isLocal, records := range procmap {
			if isLocal {
				fmt.Println("\t\tLocal")
				for record := range records {
					if _, ok = localService[record.UserName]; !ok {
						fmt.Println("\t\t\t", record.UserName)
						localService[record.UserName] = true
					}
				}
			} else {
				fmt.Println("\t\tRemote")
				for record := range records {
					fmt.Println("\t\t\t", record.RemoteAddr.String())
				}
			}
		}
	}
}
