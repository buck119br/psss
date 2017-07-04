package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"strings"

	"github.com/buck119br/psss/psss"
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

type DemandProcInfo struct {
	State     string  `json:"state"`
	StartTime int64   `json:"starttime"`
	LoadAvg   float64 `json:"loadavg"`
	VmSize    uint64  `json:"vmsize"`
	VmRSS     uint64  `json:"vmrss"`
}

type IPset map[psss.IP]bool

func NewIPset() IPset { return make(map[psss.IP]bool) }

func (i IPset) MarshalJSON() ([]byte, error) {
	temp := make([]string, 0)
	for ip := range i {
		temp = append(temp, ip.String())
	}
	return json.Marshal(temp)
}

type IPCounter map[psss.IP]int

func NewIPCounter() IPCounter { return make(map[psss.IP]int) }

func (i IPCounter) MarshalJSON() ([]byte, error) {
	tempset := make(map[string]int)
	for ip, count := range i {
		tempset[ip.String()] = count
	}
	return json.Marshal(tempset)
}

type StringSet map[string]bool

func NewStringSet() StringSet { return make(map[string]bool) }

func (s StringSet) MarshalJSON() ([]byte, error) {
	temp := make([]string, 0)
	for str := range s {
		temp = append(temp, str)
	}
	return json.Marshal(temp)
}

type ListenTopology struct {
	ProcInfo map[int]DemandProcInfo `json:"procinfo"`
	Clients  StringSet              `json:"clients,omitempty"`
	Upstream StringSet              `json:"upstream,omitempty"`
	Addrs    IPset                  `json:"addrs"`
}

func newListenTopology() *ListenTopology {
	t := new(ListenTopology)
	t.ProcInfo = make(map[int]DemandProcInfo)
	t.Clients = NewStringSet()
	t.Upstream = NewStringSet()
	t.Addrs = NewIPset()
	return t
}

type EstabTopology struct {
	ProcInfo map[int]DemandProcInfo `json:"procinfo"`
	Addrs    IPCounter              `json:"upstream"`
}

func newEstabTopology() *EstabTopology {
	t := new(EstabTopology)
	t.ProcInfo = make(map[int]DemandProcInfo)
	t.Addrs = NewIPCounter()
	return t
}

type demand struct {
	Listen map[string]*ListenTopology `json:"listen"`
	Estab  map[string]*EstabTopology  `json:"estab"`
}

func newdemand() *demand {
	d := new(demand)
	d.Listen = make(map[string]*ListenTopology)
	d.Estab = make(map[string]*EstabTopology)
	return d
}

func (d *demand) isPortListening(port string) (bool, string) {
	for name, topo := range d.Listen {
		for ip := range topo.Addrs {
			if port == ip.Port {
				return true, name
			}
		}
	}
	return false, ""
}

func (d *demand) idUserListening(user string) bool {
	for name := range d.Listen {
		if name == user {
			return true
		}
	}
	return false
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

	psss.SsFilter = 1<<psss.SsESTAB | 1<<psss.SsLISTEN
	psss.GlobalRecords = make(map[string]map[uint32]*psss.GenericRecord)
	psss.GlobalRecords["4"] = psss.GenericRecordRead(psss.ProtocalTCP, unix.AF_INET)
	psss.GlobalRecords["6"] = psss.GenericRecordRead(psss.ProtocalTCP, unix.AF_INET6)
	psss.SetUpRelation()

	var ok bool
	for _, records := range psss.GlobalRecords {
		for _, record := range records {
			if record.Status == psss.SsLISTEN {
				if _, ok = d.Listen[record.UserName]; !ok {
					d.Listen[record.UserName] = newListenTopology()
				}
				d.Listen[record.UserName].Addrs[record.LocalAddr] = true
			}
		}
	}

	var isLocalListening bool
	for _, records := range psss.GlobalRecords {
		for _, record := range records {
			if record.Status == psss.SsESTAB {
				if d.idUserListening(record.UserName) {
					isLocalListening, _ = d.isPortListening(record.LocalAddr.Port)
					for _, grecords := range psss.GlobalRecords {
						for _, grecord := range grecords {
							if grecord.LocalAddr.Port == record.RemoteAddr.Port && isHostLocal(record.RemoteAddr.Host) {
								if isLocalListening {
									d.Listen[record.UserName].Clients[grecord.UserName] = true
								} else {
									d.Listen[record.UserName].Upstream[grecord.UserName] = true
								}
								goto next
							}
						}
					}
					if isLocalListening {
						d.Listen[record.UserName].Clients[record.RemoteAddr.String()] = true
					} else {
						d.Listen[record.UserName].Upstream[record.RemoteAddr.String()] = true
					}
					goto next
				}

				if isHostLocal(record.RemoteAddr.Host) {
					if ok, _ = d.isPortListening(record.RemoteAddr.Port); ok {
						continue
					}
				}
				if _, ok = d.Estab[record.UserName]; !ok {
					d.Estab[record.UserName] = newEstabTopology()
				}
				d.Estab[record.UserName].Addrs[record.RemoteAddr] = d.Estab[record.UserName].Addrs[record.RemoteAddr] + 1
			next:
			}
		}
	}

	for name, topo := range d.Listen {
		for pid, proc := range psss.GlobalProcInfo[name] {
			topo.ProcInfo[pid] = DemandProcInfo{
				State:     psss.ProcState[proc.Stat.State],
				StartTime: int64(psss.GlobalSystemInfo.Stat.Btime + proc.Stat.Starttime/psss.SC_CLK_TCK),
				LoadAvg:   math.Trunc(float64(proc.Stat.Utime+proc.Stat.Stime)/float64(psss.GlobalSystemInfo.Stat.CPUTimes[math.MaxInt16].Total)*10000) / 10000,
				VmSize:    proc.Stat.Vsize,
				VmRSS:     uint64(proc.Stat.Rss) * uint64(os.Getpagesize()),
			}
		}
	}
	for name, topo := range d.Estab {
		for pid, proc := range psss.GlobalProcInfo[name] {
			topo.ProcInfo[pid] = DemandProcInfo{
				State:     psss.ProcState[proc.Stat.State],
				StartTime: int64(psss.GlobalSystemInfo.Stat.Btime + proc.Stat.Starttime/psss.SC_CLK_TCK),
				LoadAvg:   math.Trunc(float64(proc.Stat.Utime+proc.Stat.Stime)/float64(psss.GlobalSystemInfo.Stat.CPUTimes[math.MaxInt16].Total)*100000) / 100000,
				VmSize:    proc.Stat.Vsize,
				VmRSS:     uint64(proc.Stat.Rss) * uint64(os.Getpagesize()),
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
		fmt.Println("\t\tProcInfo")
		for _, proc := range psss.GlobalProcInfo[name] {
			fmt.Printf("\t\t\tPid:%d\n", proc.Stat.Pid)
			fmt.Printf("\t\t\t\t")
			proc.Stat.GenericInfoPrint()
			fmt.Printf("\t\t\t\t")
			proc.Stat.LoadInfoPrint()
			fmt.Printf("\t\t\t\t")
			proc.Stat.MeminfoPrint()
		}
		fmt.Println("\t\tAddrs")
		for ip := range ipmap.Addrs {
			fmt.Println("\t\t\t", ip.String())
		}
		if len(ipmap.Clients) > 0 {
			fmt.Println("\t\tClients")
			for v := range ipmap.Clients {
				fmt.Println("\t\t\t", v)
			}
		}
		if len(ipmap.Upstream) > 0 {
			fmt.Println("\t\tUpstream")
			serviceSet := make(map[string]bool)
			for v := range ipmap.Upstream {
				if _, ok = serviceSet[v]; ok {
					continue
				}
				serviceSet[v] = true
				fmt.Println("\t\t\t", v)
			}
		}
	}
	fmt.Println("Estab")
	for name, topo := range d.Estab {
		fmt.Println("\t", name)
		fmt.Println("\t\tProcInfo")
		for _, proc := range psss.GlobalProcInfo[name] {
			fmt.Printf("\t\t\tPid:%d\n", proc.Stat.Pid)
			fmt.Printf("\t\t\t\t")
			proc.Stat.GenericInfoPrint()
			fmt.Printf("\t\t\t\t")
			proc.Stat.LoadInfoPrint()
			fmt.Printf("\t\t\t\t")
			proc.Stat.MeminfoPrint()
		}
		fmt.Println("\t\tRemote")
		for ip, count := range topo.Addrs {
			fmt.Printf("\t\t\t%s (count:%d)\n", ip.String(), count)
		}
	}

	buf, err := json.Marshal(*d)
	fmt.Println(string(buf), err)
}
