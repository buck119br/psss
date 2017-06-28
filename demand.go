package main

// import (
// 	"fmt"
// 	"net"
// 	"strings"

// 	mynet "github.com/buck119br/psss/net"
// )

// var (
// 	demandData = make(map[string]map[string]map[bool]map[string]bool)
// 	localIP    = make([]string, 0, 0)
// )

// func demandRecordHandler(r *GenericRecord) {
// 	var (
// 		status            = mynet.Sstate[r.Status]
// 		procMap           map[string]map[bool]map[string]bool
// 		local             bool
// 		locOrRmtMap       map[bool]map[string]bool
// 		remoteRecord      *GenericRecord
// 		remoteServiceName string
// 		remoteServiceMap  map[string]bool
// 		ok                bool
// 	)
// 	if status != "LISTEN" && status != "ESTAB" {
// 		return
// 	}
// 	if len(r.UserName) == 0 {
// 		return
// 	}
// 	if procMap, ok = demandData[status]; !ok {
// 		procMap = make(map[string]map[bool]map[string]bool)
// 	}
// 	if locOrRmtMap, ok = procMap[r.UserName]; !ok {
// 		locOrRmtMap = make(map[bool]map[string]bool)
// 	}
// 	switch status {
// 	case "LISTEN":
// 		local = true
// 		if remoteServiceMap, ok = locOrRmtMap[local]; !ok {
// 			remoteServiceMap = make(map[string]bool)
// 		}
// 		remoteServiceMap[r.LocalAddr.String()] = true
// 	case "ESTAB":
// 		local = false
// 		for _, ip := range localIP {
// 			if strings.Contains(r.RemoteAddr.Host, ip) {
// 				local = true
// 				break
// 			}
// 		}
// 		if remoteServiceMap, ok = locOrRmtMap[local]; !ok {
// 			remoteServiceMap = make(map[string]bool)
// 		}
// 		if local {
// 			for _, remoteRecord = range GlobalTCPv4Records {
// 				if (mynet.Sstate[remoteRecord.Status] == "LISTEN" || mynet.Sstate[remoteRecord.Status] == "ESTAB") && remoteRecord.LocalAddr.Port == r.RemoteAddr.Port {
// 					remoteServiceName = remoteRecord.UserName
// 					break
// 				}
// 			}
// 			for _, remoteRecord = range GlobalTCPv6Records {
// 				if (mynet.Sstate[remoteRecord.Status] == "LISTEN" || mynet.Sstate[remoteRecord.Status] == "ESTAB") && remoteRecord.LocalAddr.Port == r.RemoteAddr.Port {
// 					remoteServiceName = remoteRecord.UserName
// 					break
// 				}
// 			}
// 			if len(remoteServiceName) != 0 {
// 				remoteServiceMap[remoteServiceName] = true
// 			} else {
// 				remoteServiceMap[r.RemoteAddr.String()] = true
// 			}
// 		} else {
// 			remoteServiceMap[r.RemoteAddr.String()] = true
// 		}
// 	}
// 	locOrRmtMap[local] = remoteServiceMap
// 	procMap[r.UserName] = locOrRmtMap
// 	demandData[status] = procMap
// }

// func DemandShow() {
// 	var stringBuff []string
// 	localAddr, err := net.InterfaceAddrs()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, v := range localAddr {
// 		stringBuff = strings.Split(v.String(), "/")
// 		localIP = append(localIP, stringBuff[0])
// 	}
// 	if Family&FbTCPv4 != 0 {
// 		for _, record := range GlobalTCPv4Records {
// 			demandRecordHandler(record)
// 		}
// 	}
// 	if Family&FbTCPv6 != 0 {
// 		for _, record := range GlobalTCPv6Records {
// 			demandRecordHandler(record)
// 		}
// 	}
// 	for status, localServiceMap := range demandData {
// 		fmt.Println(status)
// 		for procName, locOrRmtMap := range localServiceMap {
// 			fmt.Println("\t" + procName)
// 			for local, remoteServiceMap := range locOrRmtMap {
// 				if status == "ESTAB" {
// 					if local {
// 						fmt.Println("\t\tLocal")
// 					} else {
// 						fmt.Println("\t\tRemote")
// 					}
// 				}
// 				for addr := range remoteServiceMap {
// 					fmt.Println("\t\t\t" + addr)
// 				}
// 			}
// 		}
// 	}
// }
