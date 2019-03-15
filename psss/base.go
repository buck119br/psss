package psss

/*
#include <unistd.h>
*/
import "C"

import (
	"fmt"
	"os"
	"regexp"
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
	OSPageSize = os.Getpagesize()

	SC_CLK_TCK = uint64(C.sysconf(C._SC_CLK_TCK))
)

var (
	ErrorDone = fmt.Errorf("Done")
)

var (
	SlimSpaceRegExp *regexp.Regexp
)

var (
	// channel
	SocketInfoChan chan SocketInfo
	ProcInfoChan   chan *ProcInfo

	GlobalProcFds map[string]map[int]map[uint32]Fd

	bytesCounter int
)

func init() {
	var err error

	if SlimSpaceRegExp, err = regexp.Compile(`[\s]+`); err != nil {
		panic(err)
	}

	SocketInfoChan = make(chan SocketInfo)
	ProcInfoChan = make(chan *ProcInfo)

	GlobalProcFds = make(map[string]map[int]map[uint32]Fd)

	archInit()
}

func AddrLengthInit() {
	MaxLocalAddrLength = 17
	MaxRemoteAddrLength = 18
}
