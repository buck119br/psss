package psss

/*
#include <unistd.h>
*/
import "C"

import (
	"bytes"
	"fmt"
	"os"
	"syscall"

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

	MaxLocalAddrLength  int
	MaxRemoteAddrLength int
)

var (
	pageSize = os.Getpagesize()

	SC_CLK_TCK = uint64(C.sysconf(C._SC_CLK_TCK))
)

var (
	ErrorDone = fmt.Errorf("Done")
)

var (
	// channel
	RecordInputChan    chan *GenericRecord
	RecordOutputChan   chan *GenericRecord
	ProcInfoInputChan  chan *ProcInfo
	ProcInfoOutputChan chan *ProcInfo

	globalProcInfo map[string]map[int]*ProcInfo
	// buffer
	sockDiagMsgBuffer   []byte
	unDiagRequestBuffer []byte
	inDiagRequestBuffer []byte
	fileContentBuffer   *bytes.Buffer

	int64Buffer  int64
	intBuffer    int
	indexBuffer  int
	bytesCounter int

	sockAddrNl  unix.SockaddrNetlink
	nlAttr      unix.NlAttr
	unDiagReq   UnixDiagRequest
	unDiagMsg   UnixDiagMessage
	unDiagRQlen UnixDiagRQlen
	inDiagReq   InetDiagRequest
	inDiagMsg   InetDiagMessage

	procDirentHandler *DirentHandler
	fdDirentHandler   *DirentHandler
	fdStat_t          *syscall.Stat_t
)

func init() {
	RecordInputChan = make(chan *GenericRecord)
	RecordOutputChan = make(chan *GenericRecord)
	ProcInfoInputChan = make(chan *ProcInfo)
	ProcInfoOutputChan = make(chan *ProcInfo)

	sockDiagMsgBuffer = make([]byte, pageSize)
	unDiagRequestBuffer = make([]byte, SizeOfUnixDiagRequest)
	inDiagRequestBuffer = make([]byte, SizeOfInetDiagRequest)
	fileContentBuffer = bytes.NewBuffer(make([]byte, pageSize))

	procDirentHandler = NewDirentHandler()
	fdDirentHandler = NewDirentHandler()
	fdStat_t = new(syscall.Stat_t)
}

func AddrLengthInit() {
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}
