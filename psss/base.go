package psss

/*
#include <unistd.h>
*/
import "C"

import (
	"fmt"
	"os"
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

const ProcRoot = "/proc"

var (
	AfFilter       uint64
	ProtocalFilter uint64
	SsFilter       uint32

	FlagInfo   bool
	FlagMemory bool
)

var (
	Summary               map[string]map[string]int
	GlobalRecords         map[uint32]*GenericRecord
	GlobalProcInfo        map[string]map[int]*ProcInfo
	GlobalSystemInfo      *SystemInfo
	GlobalBuffer          []byte
	UnixDiagRequestBuffer []byte
	InetDiagRequestBuffer []byte
	RecordInputChan       chan *GenericRecord
	RecordOutputChan      chan *GenericRecord
)

var (
	MaxLocalAddrLength  int
	MaxRemoteAddrLength int
)

var (
	SC_CLK_TCK = uint64(C.sysconf(C._SC_CLK_TCK))
)

var (
	ErrorDone = fmt.Errorf("Done")
)

func init() {
	GlobalProcInfo = make(map[string]map[int]*ProcInfo)
	GlobalBuffer = make([]byte, os.Getpagesize())
	UnixDiagRequestBuffer = make([]byte, SizeOfUnixDiagRequest)
	InetDiagRequestBuffer = make([]byte, SizeOfInetDiagRequest)
	RecordInputChan = make(chan *GenericRecord)
	RecordOutputChan = make(chan *GenericRecord)
}

func AddrLengthInit() {
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}
