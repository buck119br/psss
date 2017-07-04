package psss

/*
#include <unistd.h>
*/
import "C"

import (
	"fmt"
)

const (
	ProtocalUnknown = 1 << iota
	ProtocalDCCP
	ProtocalNetlink
	ProtocalPacket
	ProtocalRAW
	ProtocalSCTP
	ProtocalTCP
	ProtocalUDP
	ProtocalUnix
	ProtocalMax
)

const (
	GlobalRecordsKey = "GlobalRecords"

	ProcRoot = "/proc"
)

var (
	AfFilter       uint64
	ProtocalFilter uint64
	SsFilter       uint32

	Summary          map[string]map[string]int
	GlobalRecords    map[string]map[uint32]*GenericRecord
	GlobalProcInfo   map[string]map[int]*ProcInfo
	GlobalSystemInfo *SystemInfo

	MaxLocalAddrLength  int
	MaxRemoteAddrLength int

	SC_CLK_TCK = uint64(C.sysconf(C._SC_CLK_TCK))

	ErrorDone = fmt.Errorf("Done")
)

func init() {
	Summary = make(map[string]map[string]int)
	for _, pf := range SummaryPF {
		Summary[pf] = make(map[string]int)
	}

	GlobalRecords = make(map[string]map[uint32]*GenericRecord)
	GlobalProcInfo = make(map[string]map[int]*ProcInfo)
}

func AddrLengthInit() {
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}
