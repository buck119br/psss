package main

import (
	"fmt"
	"net"
	"strings"

	mynet "github.com/buck119br/psss/net"
)

var (
	MaxLocalAddrLength  = 17
	MaxRemoteAddrLength = 18

	NewlineFlag bool
)

func ShowSummary() {
	// Read
	if err := GenericReadSockstat(false); err != nil {
		return
	}
	if err := GenericReadSockstat(true); err != nil {
		return
	}
	// Display
	var format string
	fmt.Println("Transport\t Total\t IPv4\t IPv6\t")
	for _, pf := range SummaryPF {
		if len(pf) >= 8 {
			format = "%s\t %d\t %d\t %d\t\n"
		} else {
			format = "%s\t\t %d\t %d\t %d\t\n"
		}
		fmt.Printf(format, pf, Summary[pf][IPv4String]+Summary[pf][IPv6String], Summary[pf][IPv4String], Summary[pf][IPv6String])
	}
}

func GenericShow(family string, records map[uint32]*GenericRecord) {
	var ok bool
	for _, record := range records {
		if !*flagAll && !mynet.SstateActive[record.Status] {
			continue
		}
		if *flagListen && !mynet.SstateListen[record.Status] {
			continue
		}
		switch family {
		case TCPv4Str:
			fmt.Printf("tcp4\t")
		case TCPv6Str:
			fmt.Printf("tcp6\t")
		case UDPv4Str:
			fmt.Printf("udp4\t")
		case UDPv6Str:
			fmt.Printf("udp6\t")
		case RAWv4Str:
			fmt.Printf("raw4\t")
		case RAWv6Str:
			fmt.Printf("raw6\t")
		case UnixStr:
			if _, ok = mynet.SocketType[record.Type]; !ok {
				fmt.Printf("dgr\t")
			} else {
				fmt.Printf("%s\t", mynet.SocketType[record.Type])
			}
		}
		if len(mynet.Sstate[record.Status]) >= 8 {
			fmt.Printf("%s\t", mynet.Sstate[record.Status])
		} else {
			fmt.Printf("%s\t\t", mynet.Sstate[record.Status])
		}
		fmt.Printf("%d\t%d\t", record.RxQueue, record.TxQueue)
		fmt.Printf("%-*s\t%-*s\t", MaxLocalAddrLength, record.LocalAddr.String(), MaxRemoteAddrLength, record.RemoteAddr.String())
		// Process Info
		if *flagProcess && len(record.UserName) > 0 {
			record.ProcInfoPrint()
		}
		if NewlineFlag {
			fmt.Printf("\n")
		}
		if family != UnixStr {
			// Timer Info
			if *flagOption && record.Timer != 0 {
				record.TimerInfoPrint()
			}
			// Detailed Info
			if *flagExtended {
				record.ExtendInfoPrint()
			}
		}
		// Meminfo
		if *flagMemory && len(record.Meminfo) == 8 {
			record.MeminfoPrint()
		}
		// internal TCP info
		if *flagInfo && (family == TCPv4Str || family == TCPv6Str) {
			record.TCPInfoPrint()
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
	if Family&FbRAWv4 != 0 {
		GenericShow(RAWv4Str, GlobalRAWv4Records)
	}
	if Family&FbRAWv6 != 0 {
		GenericShow(RAWv6Str, GlobalRAWv6Records)
	}
	if Family&FbUnix != 0 {
		GenericShow(UnixStr, GlobalUnixRecords)
	}
}

var (
	demandData = make(map[string]map[string]map[bool]map[string]bool)
	localIP    = make([]string, 0, 0)
)

func demandRecordHandler(r *GenericRecord) {
	var (
		status            = mynet.Sstate[r.Status]
		procMap           map[string]map[bool]map[string]bool
		local             bool
		locOrRmtMap       map[bool]map[string]bool
		remoteRecord      *GenericRecord
		remoteServiceName string
		remoteServiceMap  map[string]bool
		ok                bool
	)
	if status != "LISTEN" && status != "ESTAB" {
		return
	}
	if len(r.UserName) == 0 {
		return
	}
	if procMap, ok = demandData[status]; !ok {
		procMap = make(map[string]map[bool]map[string]bool)
	}
	if locOrRmtMap, ok = procMap[r.UserName]; !ok {
		locOrRmtMap = make(map[bool]map[string]bool)
	}
	switch status {
	case "LISTEN":
		local = true
		if remoteServiceMap, ok = locOrRmtMap[local]; !ok {
			remoteServiceMap = make(map[string]bool)
		}
		remoteServiceMap[r.LocalAddr.String()] = true
	case "ESTAB":
		local = false
		for _, ip := range localIP {
			if strings.Contains(r.RemoteAddr.Host, ip) {
				local = true
				break
			}
		}
		if remoteServiceMap, ok = locOrRmtMap[local]; !ok {
			remoteServiceMap = make(map[string]bool)
		}
		if local {
			for _, remoteRecord = range GlobalTCPv4Records {
				if (mynet.Sstate[remoteRecord.Status] == "LISTEN" || mynet.Sstate[remoteRecord.Status] == "ESTAB") && remoteRecord.LocalAddr.Port == r.RemoteAddr.Port {
					remoteServiceName = remoteRecord.UserName
					break
				}
			}
			for _, remoteRecord = range GlobalTCPv6Records {
				if (mynet.Sstate[remoteRecord.Status] == "LISTEN" || mynet.Sstate[remoteRecord.Status] == "ESTAB") && remoteRecord.LocalAddr.Port == r.RemoteAddr.Port {
					remoteServiceName = remoteRecord.UserName
					break
				}
			}
			if len(remoteServiceName) != 0 {
				remoteServiceMap[remoteServiceName] = true
			} else {
				remoteServiceMap[r.RemoteAddr.String()] = true
			}
		} else {
			remoteServiceMap[r.RemoteAddr.String()] = true
		}
	}
	locOrRmtMap[local] = remoteServiceMap
	procMap[r.UserName] = locOrRmtMap
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
			demandRecordHandler(record)
		}
	}
	if Family&FbTCPv6 != 0 {
		for _, record := range GlobalTCPv6Records {
			demandRecordHandler(record)
		}
	}
	for status, localServiceMap := range demandData {
		fmt.Println(status)
		for procName, locOrRmtMap := range localServiceMap {
			fmt.Println("\t" + procName)
			for local, remoteServiceMap := range locOrRmtMap {
				if status == "ESTAB" {
					if local {
						fmt.Println("\t\tLocal")
					} else {
						fmt.Println("\t\tRemote")
					}
				}
				for addr := range remoteServiceMap {
					fmt.Println("\t\t\t" + addr)
				}
			}
		}
	}
}
