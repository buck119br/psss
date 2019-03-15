package topo

import (
	"bytes"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/buck119br/psss/psss"
	"github.com/glycerine/zebrapack/msgp"
)

var (
	numCPU     uint64
	pageSize   uint64
	localAddrs []string

	SysStatNew     *psss.SystemStat
	SysStatOld     *psss.SystemStat
	GlobalTopology *Topology

	MsgpBuffer *bytes.Buffer
	MsgpWriter *msgp.Writer

	procsInfoReserve map[string]map[int]*ProcInfoReserve
	localPortToName  map[string]string

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

	SysStatNew = new(psss.SystemStat)
	SysStatOld = new(psss.SystemStat)
	GlobalTopology = NewTopology()

	MsgpBuffer = bytes.NewBuffer(make([]byte, 0, 512*1024))
	MsgpWriter = msgp.NewWriter(MsgpBuffer)

	procsInfoReserve = make(map[string]map[int]*ProcInfoReserve)
	localPortToName = make(map[string]string)

	procInfoReserve = new(ProcInfoReserve)

	psss.FlagProcess = true
}
