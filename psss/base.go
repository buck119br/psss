package psss

/*
#include <unistd.h>
*/
import "C"

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/sys/unix"
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

	FlagProcess bool
	FlagInfo    bool
	FlagMemory  bool
)

var (
	Summary          map[string]map[string]int
	GlobalRecords    map[uint32]*GenericRecord
	GlobalProcInfo   map[string]map[int]*ProcInfo
	GlobalSystemInfo *SystemInfo
	// buffer
	GlobalBuffer        []byte
	FileContentBuffer   *bytes.Buffer
	unDiagRequestBuffer []byte
	inDiagRequestBuffer []byte
	int64Buffer         int64
	intBuffer           int
	indexBuffer         int
	bytesCounter        int
	sockAddrNl          unix.SockaddrNetlink
	nlAttr              unix.NlAttr
	unDiagMsg           UnixDiagMessage
	unDiagRQlen         UnixDiagRQlen
	inDiagReq           InetDiagRequest
	inDiagMsg           InetDiagMessage
	fdPath              string
	fdLink              string
	fdInode             uint32
	// channel
	RecordInputChan    chan *GenericRecord
	RecordOutputChan   chan *GenericRecord
	ProcInfoInputChan  chan *ProcInfo
	ProcInfoOutputChan chan *ProcInfo
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
	GlobalBuffer = make([]byte, os.Getpagesize())
	FileContentBuffer = bytes.NewBuffer(make([]byte, os.Getpagesize()))
	unDiagRequestBuffer = make([]byte, SizeOfUnixDiagRequest)
	inDiagRequestBuffer = make([]byte, SizeOfInetDiagRequest)
	RecordInputChan = make(chan *GenericRecord)
	RecordOutputChan = make(chan *GenericRecord)
	ProcInfoInputChan = make(chan *ProcInfo)
	ProcInfoOutputChan = make(chan *ProcInfo)
}

func AddrLengthInit() {
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}
