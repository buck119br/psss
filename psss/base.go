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
	SocketInfoChan chan SocketInfo
	ProcInfoChan   chan ProcInfo

	globalProcInfo map[string]map[int]ProcInfo

	GlobalProcFds map[string]map[int]map[uint32]Fd

	int64Buffer  int64
	intBuffer    int
	indexBuffer  int
	bytesCounter int
)

func init() {
	SocketInfoChan = make(chan SocketInfo)
	ProcInfoChan = make(chan ProcInfo)

	GlobalProcFds = make(map[string]map[int]map[uint32]Fd)

	archInit()
}

func AddrLengthInit() {
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}
