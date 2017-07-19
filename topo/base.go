package topo

import (
	"os"
	"runtime"

	"github.com/buck119br/GoUtils/bufferPool"
	"github.com/buck119br/psss/psss"
)

var (
	SysInfoNew     *psss.SystemInfo
	SysInfoOld     *psss.SystemInfo
	GlobalTopology *Topology

	globalBufferPool *bufferPool.BufferPool

	numCPU   uint64
	pageSize uint64

	procsInfoReserve map[string]map[int]*ProcInfoReserve
	serviceInfo      ServiceInfo
	originProcInfo   psss.ProcInfo
	procStat         ProcStat
	procInfoReserve  *ProcInfoReserve
	ok               bool
)

func init() {
	SysInfoNew = psss.NewSystemInfo()
	SysInfoOld = psss.NewSystemInfo()
	GlobalTopology = NewTopology()

	bufferPool.BufferSize = 200
	bufferPool.BufferPoolInitSize = 10
	bufferPool.BufferPoolEnlargeFactor = 10
	bufferPool.BufferPoolMaxCapacity = 100
	bufferPool.Init()
	globalBufferPool = bufferPool.NewBufferPool()

	numCPU = uint64(runtime.NumCPU())
	pageSize = uint64(os.Getpagesize())

	procsInfoReserve = make(map[string]map[int]*ProcInfoReserve)

	procInfoReserve = new(ProcInfoReserve)
}
