package topo

import (
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/buck119br/psss/psss"
)

var (
	numCPU     uint64
	pageSize   uint64
	localAddrs []string

	SysInfoNew     *psss.SystemInfo
	SysInfoOld     *psss.SystemInfo
	GlobalTopology *Topology

	procsInfoReserve map[string]map[int]*ProcInfoReserve
	localAddrToName  map[string]string

	originProcInfo  psss.ProcInfo
	serviceInfo     *ServiceInfo
	procInfoReserve *ProcInfoReserve
	procStat        ProcStat

	addr      Addr
	addrState AddrState
	ok        bool
)

func init() {
	netAddrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	localAddrs = make([]string, 0)
	for _, v := range netAddrs {
		localAddrs = append(localAddrs, strings.Split(v.String(), "/")[0])
	}
	localAddrs = append(localAddrs, "127.0.1.1")

	numCPU = uint64(runtime.NumCPU())
	pageSize = uint64(os.Getpagesize())

	SysInfoNew = psss.NewSystemInfo()
	SysInfoOld = psss.NewSystemInfo()
	GlobalTopology = NewTopology()

	procsInfoReserve = make(map[string]map[int]*ProcInfoReserve)
	localAddrToName = make(map[string]string)

	procInfoReserve = new(ProcInfoReserve)

	psss.FlagProcess = true
}
