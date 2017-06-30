package main

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/sys/unix"
)

var localAddrs = make(map[string]bool)

func isHostLocal(host string) bool {
	for v := range localAddrs {
		if strings.Contains(host, v) {
			return true
		}
	}
	return false
}

type topology struct {
	employer map[string]bool
	employee map[string]bool
	ports    map[IP]bool
}

func newtopology() *topology {
	t := new(topology)
	t.employer = make(map[string]bool)
	t.employee = make(map[string]bool)
	t.ports = make(map[IP]bool)
	return t
}

type demand struct {
	Listen map[string]*topology
	Estab  map[string]map[bool]map[*GenericRecord]bool
}

func newdemand() *demand {
	d := new(demand)
	d.Listen = make(map[string]*topology)
	d.Estab = make(map[string]map[bool]map[*GenericRecord]bool)
	return d
}

func (d *demand) isPortListening(port string) (bool, string) {
	for name, topo := range d.Listen {
		for ip := range topo.ports {
			if port == ip.Port {
				return true, name
			}
		}
	}
	return false, ""
}

func (d *demand) data() {
	netAddrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range netAddrs {
		localAddrs[strings.Split(v.String(), "/")[0]] = true
	}
	localAddrs["127.0.1.1"] = true

	ssFilter = 1<<SsESTAB | 1<<SsLISTEN
	GlobalRecords = make(map[string]map[uint32]*GenericRecord)
	GlobalRecords["4"] = GenericRecordRead(ProtocalTCP, unix.AF_INET)
	GlobalRecords["6"] = GenericRecordRead(ProtocalTCP, unix.AF_INET6)
	SetUpRelation()

	var ok bool
	for _, records := range GlobalRecords {
		for _, record := range records {
			if record.Status == SsLISTEN {
				if _, ok = d.Listen[record.UserName]; !ok {
					d.Listen[record.UserName] = newtopology()
				}
				d.Listen[record.UserName].ports[record.LocalAddr] = true
			}
		}
	}

	var (
		name          string
		isRemoteLocal bool
	)
	for _, records := range GlobalRecords {
		for _, record := range records {
			if record.Status == SsESTAB {
				if ok, _ = d.isPortListening(record.LocalAddr.Port); ok {
					for _, grecords := range GlobalRecords {
						for _, grecord := range grecords {
							if grecord.LocalAddr.Port == record.RemoteAddr.Port {
								d.Listen[record.UserName].employee[grecord.UserName] = true
								continue
							}
							goto next
						}
					}
					d.Listen[record.UserName].employee[record.RemoteAddr.String()] = true
				next:
					continue
				}
				if isRemoteLocal = isHostLocal(record.RemoteAddr.Host); isRemoteLocal {
					if ok, name = d.isPortListening(record.RemoteAddr.Port); ok {
						d.Listen[name].employer[record.UserName] = true
						continue
					}
				}
				if _, ok = d.Estab[record.UserName]; !ok {
					d.Estab[record.UserName] = make(map[bool]map[*GenericRecord]bool)
				}
				if _, ok = d.Estab[record.UserName][isRemoteLocal]; !ok {
					d.Estab[record.UserName][isRemoteLocal] = make(map[*GenericRecord]bool)
				}
				d.Estab[record.UserName][isRemoteLocal][record] = true
			}
		}
	}
}

func (d *demand) show() {
	var ok bool
	d.data()
	fmt.Println("Listen")
	for name, ipmap := range d.Listen {
		fmt.Println("\t", name)
		fmt.Println("\t\tPorts")
		for ip := range ipmap.ports {
			fmt.Println("\t\t\t", ip.String())
		}
		if len(ipmap.employer) > 0 {
			fmt.Println("\t\tEmployers")
			for v := range ipmap.employer {
				fmt.Println("\t\t\t", v)
			}
		}
		if len(ipmap.employee) > 0 {
			fmt.Println("\t\tEmployees")
			serviceSet := make(map[string]bool)
			for v := range ipmap.employee {
				if _, ok = serviceSet[v]; ok {
					continue
				}
				serviceSet[v] = true
				fmt.Println("\t\t\t", v)
			}
		}
	}
	fmt.Println("Estab")
	for name, procmap := range d.Estab {
		fmt.Println("\t", name)
		for isLocal, records := range procmap {
			if isLocal {
				fmt.Println("\t\tLocal", len(records))
			} else {
				fmt.Println("\t\tRemote", len(records))
			}
			serviceSet := make(map[string]bool)
			for record := range records {
				if _, ok = serviceSet[record.RemoteAddr.String()]; ok {
					continue
				}
				serviceSet[record.RemoteAddr.String()] = true
				fmt.Println("\t\t\t", record.RemoteAddr.String())
			}
		}
	}
}
