package topo

import (
	"strings"
)

type ProcInfoReserve struct {
	Utime uint64
	Stime uint64
	Fresh bool
}

func clearReserve() {
	var pid int
	for name := range procsInfoReserve {
		for pid = range procsInfoReserve[name] {
			if procsInfoReserve[name][pid].Fresh {
				procsInfoReserve[name][pid].Fresh = false
			} else {
				delete(procsInfoReserve[name], pid)
				if len(procsInfoReserve[name]) == 0 {
					delete(procsInfoReserve, name)
				}
			}
		}
	}
}

func isHostLocal(host string) bool {
	for _, v := range localAddrs {
		if strings.Contains(host, v) {
			return true
		}
	}
	return false
}
