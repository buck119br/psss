package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	Sockstat4Path = "/proc/net/sockstat"
	Sockstat6Path = "/proc/net/sockstat6"
)

var (
	Summary map[string]map[string]int

	Protocal = []string{
		"RAW",
		"UDP",
		"TCP",
		"FRAG",
	}
)

func GetSocketCount(fields []string) (int, error) {
	for i := range fields {
		if fields[i] == "inuse" {
			return strconv.Atoi(fields[i+1])
		}
	}
	return 0, nil
}

// IPv6:versionFlag = true; IPv4:versionFlag = false
func GenericReadSockstat(versionFlag bool) (err error) {
	var (
		file      *os.File
		line      string
		fields    []string
		tempCount int
	)
	if versionFlag {
		file, err = os.Open(Sockstat6Path)
	} else {
		file, err = os.Open(Sockstat4Path)
	}
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		fields = strings.Fields(line)
		switch fields[0] {
		case "sockets:":
			continue
		case "TCP:":
			if Summary["TCP"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		case "TCP6:":
			if Summary["TCP"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		case "UDP:":
			if Summary["UDP"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		case "UDP6:":
			if Summary["UDP"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		case "UDPLITE:":
			if tempCount, err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
			Summary["UDP"][IPv4String] += tempCount
		case "UDPLITE6:":
			if tempCount, err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
			Summary["UDP"][IPv6String] += tempCount
		case "RAW:":
			if Summary["RAW"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		case "RAW6:":
			if Summary["RAW"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		case "FRAG:":
			if Summary["FRAG"][IPv4String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		case "FRAG6:":
			if Summary["FRAG"][IPv6String], err = GetSocketCount(fields[1:]); err != nil {
				return err
			}
		default:
			continue
		}
	}
	return nil
}
