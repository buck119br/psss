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
	SocketInfoInputChan  chan *SocketInfo
	SocketInfoOutputChan chan *SocketInfo
	ProcInfoInputChan    chan *ProcInfo
	ProcInfoOutputChan   chan *ProcInfo

	globalProcInfo map[string]map[int]*ProcInfo

	int64Buffer  int64
	intBuffer    int
	indexBuffer  int
	bytesCounter int
)

func init() {
	SocketInfoInputChan = make(chan *SocketInfo)
	SocketInfoOutputChan = make(chan *SocketInfo)
	ProcInfoInputChan = make(chan *ProcInfo)
	ProcInfoOutputChan = make(chan *ProcInfo)

	archInit()
}

func AddrLengthInit() {
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}
