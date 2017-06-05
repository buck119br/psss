package main

import (
	"fmt"
	"net"
	"sort"
	"strings"
)

var (
	MaxLocalAddrLength  = 17
	MaxRemoteAddrLength = 18
)

func ShowSummary() (err error) {
	// Read
	if err = GenericReadSockstat(false); err != nil {
		return err
	}
	if err = GenericReadSockstat(true); err != nil {
		return err
	}
	// Display
	var format string
	fmt.Println("Transport\t Total\t IPv4\t IPv6\t")
	for _, protocal := range Protocal {
		if len(protocal) >= 8 {
			format = "%s\t %d\t %d\t %d\t\n"
		} else {
			format = "%s\t\t %d\t %d\t %d\t\n"
		}
		fmt.Printf(format, protocal, Summary[protocal][IPv4String]+Summary[protocal][IPv6String], Summary[protocal][IPv4String], Summary[protocal][IPv6String])
	}
	return nil
}

func ShowTCP(records map[uint64]*TCPRecord) {
	for _, record := range records {
		fmt.Printf("tcp\t")
		if len(TCPState[record.Status]) > 8 {
			fmt.Printf("%s\t", TCPState[record.Status])
		} else {
			fmt.Printf("%s\t\t", TCPState[record.Status])
		}
		fmt.Printf("%d\t%d\t", record.RxQueue, record.TxQueue)
		fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, record.LocalAddr.String(), MaxRemoteAddrLength, record.RemoteAddr.String())
		if *flagProcess {
			fmt.Printf("[")
			for _, proc := range record.Procs {
				for _, fd := range proc.Fd {
					if fd.SysStat.Ino == record.Inode {
						fmt.Printf(`("%s",pid=%d,fd=%d)`, proc.Name, proc.Pid, fd.Name)
					}
				}
			}
			fmt.Printf("]")
		}
		fmt.Printf("\n")
	}
}

func SocketShow() {
	fmt.Printf("Netid\tState\t\tRecv-Q\tSend-Q\t")
	fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, "LocalAddress:Port", MaxRemoteAddrLength, "RemoteAddress:Port")
	if *flagProcess {
		fmt.Printf("Users")
	}
	fmt.Printf("\n")
	if Family|FbTCPv4 != 0 {
		ShowTCP(GlobalTCPv4Records)
	}
	if Family|FbTCPv6 != 0 {
		ShowTCP(GlobalTCPv6Records)
	}
}

func Show() {
	var (
		procRecords map[string]map[bool][]*TCPRecord
		lcOrRmt     map[bool][]*TCPRecord
		records     []*TCPRecord
		procName    string
		local       bool
		status      string
		ok          bool
		showFormat  string
	)
	localAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	localIP := make([]string, 0, 0)
	stringBuff := make([]string, 0, 0)
	for _, v := range localAddr {
		stringBuff = strings.Split(v.String(), "/")
		localIP = append(localIP, stringBuff[0])
	}
	statusMap := make(map[string]map[string]map[bool][]*TCPRecord)
	for _, record := range GlobalTCPv4Records {
		status = TCPState[int(record.Status)]
		if status != "LISTEN" && status != "ESTAB" {
			continue
		}
		if procRecords, ok = statusMap[status]; !ok {
			procRecords = make(map[string]map[bool][]*TCPRecord)
		}
		for _, proc := range record.Procs {
			for _, fd := range proc.Fd {
				if fd.SysStat.Ino == record.Inode {
					procName = proc.Name
					if lcOrRmt, ok = procRecords[procName]; !ok {
						lcOrRmt = make(map[bool][]*TCPRecord)
					}
					local = false
					for _, v := range localIP {
						if strings.Contains(record.RemoteAddr.String(), v) {
							local = true
						}
					}
					if records, ok = lcOrRmt[local]; !ok {
						records = make([]*TCPRecord, 0, 0)
					}
					break
				}
			}
		}
		records = append(records, record)
		lcOrRmt[local] = records
		procRecords[procName] = lcOrRmt
		statusMap[status] = procRecords
	}
	for status, procRecords = range statusMap {
		fmt.Println(status)
		for procName, lcOrRmt = range procRecords {
			fmt.Println("\t", procName)
			for local, records = range lcOrRmt {
				if status == "ESTAB" {
					if local {
						fmt.Println("\t\tLocal")
					} else {
						fmt.Println("\t\tRemote")
					}
				}
				sort.Slice(records, func(i, j int) bool { return records[i].LocalAddr.String() < records[j].LocalAddr.String() })
				for _, v := range records {
					if status == "ESTAB" {
						if len(v.LocalAddr.String()) >= 16 {
							showFormat = "\t\t\t%s\t %s\n"
						} else {
							showFormat = "\t\t\t%s\t\t %s\n"
						}
					} else {
						if len(v.LocalAddr.String()) >= 16 {
							showFormat = "\t\t%s\t %s\n"
						} else {
							showFormat = "\t\t%s\t\t %s\n"
						}
					}
					fmt.Printf(showFormat, v.LocalAddr, v.RemoteAddr)
				}
			}
		}
	}
}
