package main

import (
	"fmt"
	"net"
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

func GenericShow(family string, records map[uint64]*GenericRecord) {
	for _, record := range records {
		if !*flagAll && !SstateActive[record.Status] {
			continue
		}
		switch family {
		case TCPv4Str, TCPv6Str:
			fmt.Printf("tcp\t")
		case UDPv4Str, UDPv6Str:
			fmt.Printf("udp\t")
		}
		if len(Sstate[record.Status]) > 8 {
			fmt.Printf("%s\t", Sstate[record.Status])
		} else {
			fmt.Printf("%s\t\t", Sstate[record.Status])
		}
		fmt.Printf("%d\t%d\t", record.RxQueue, record.TxQueue)
		fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, record.LocalAddr.String(), MaxRemoteAddrLength, record.RemoteAddr.String())
		if *flagProcess {
			fmt.Printf("[")
			for _, proc := range record.Procs {
				for _, fd := range proc.Fd {
					if fd.SysStat.Ino == record.Inode {
						fmt.Printf(`("%s",pid=%d,fd=%s)`, proc.Name, proc.Pid, fd.Name)
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
	if Family&FbTCPv4 != 0 {
		GenericShow(TCPv4Str, GlobalTCPv4Records)
	}
	if Family&FbTCPv6 != 0 {
		GenericShow(TCPv6Str, GlobalTCPv6Records)
	}
	if Family&FbUDPv4 != 0 {
		GenericShow(UDPv4Str, GlobalUDPv4Records)
	}
	if Family&FbUDPv6 != 0 {
		GenericShow(UDPv6Str, GlobalUDPv6Records)
	}
}

var (
	demandData = make(map[string]map[string]map[bool]map[string]string)
	localIP    = make([]string, 0, 0)
)

func demandRecordHandler(family string, r *GenericRecord) {
	var (
		status           = Sstate[r.Status]
		procName         string
		procMap          map[string]map[bool]map[string]string
		local            bool
		locOrRmtMap      map[bool]map[string]string
		remoteServiceMap map[string]string
		ok               bool
		isFind           bool
	)
	if status != "LISTEN" && status != "ESTAB" {
		return
	}
	if procMap, ok = demandData[status]; !ok {
		procMap = make(map[string]map[bool]map[string]string)
	}
	for _, proc := range r.Procs {
		for _, fd := range proc.Fd {
			if r.Inode == fd.SysStat.Ino {
				procName = proc.Name
				if locOrRmtMap, ok = procMap[procName]; !ok {
					locOrRmtMap = make(map[bool]map[string]string)
				}
				local = false
				for _, ip := range localIP {
					if strings.Contains(r.RemoteAddr.Host, ip) {
						local = true
						break
					}
				}
				if remoteServiceMap, ok = locOrRmtMap[local]; !ok {
					remoteServiceMap = make(map[string]string)
				}
				remoteServiceMap[r.RemoteAddr.String()] = family
				isFind = true
				goto end
			}
		}
	}
end:
	if !isFind {
		return
	}
	locOrRmtMap[local] = remoteServiceMap
	procMap[procName] = locOrRmtMap
	demandData[status] = procMap
}
func DemandShow() {
	var stringBuff []string
	localAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range localAddr {
		stringBuff = strings.Split(v.String(), "/")
		localIP = append(localIP, stringBuff[0])
	}
	if Family&FbTCPv4 != 0 {
		for _, record := range GlobalTCPv4Records {
			demandRecordHandler("tcp", record)
		}
	}
	if Family&FbTCPv6 != 0 {
		for _, record := range GlobalTCPv6Records {
			demandRecordHandler("tcp", record)
		}
	}
	if Family&FbUDPv4 != 0 {
		for _, record := range GlobalUDPv4Records {
			demandRecordHandler("udp", record)
		}
	}
	if Family&FbUDPv6 != 0 {
		for _, record := range GlobalUDPv6Records {
			demandRecordHandler("udp", record)
		}
	}
	for status, localServiceMap := range demandData {
		fmt.Println(status)
		for procName, locOrRmtMap := range localServiceMap {
			fmt.Println("\t" + procName)
			for local, remoteServiceMap := range locOrRmtMap {
				if local {
					fmt.Println("\t\tLocal")
				} else {
					fmt.Println("\t\tRemote")
				}
				for remoteAddr, family := range remoteServiceMap {
					fmt.Println("\t\t\t" + family + "\t" + remoteAddr)
				}
			}
		}
	}
}
