// +build linux

package psss

import (
	"bytes"
	"syscall"

	"golang.org/x/sys/unix"
)

const ProcRoot = "/proc"

var (
	// buffer
	sockDiagMsgBuffer   []byte
	unDiagRequestBuffer []byte
	inDiagRequestBuffer []byte
	fileContentBuffer   *bytes.Buffer

	sockAddrNl  unix.SockaddrNetlink
	nlAttr      unix.NlAttr
	unDiagReq   UnixDiagRequest
	unDiagMsg   UnixDiagMessage
	unDiagRQlen UnixDiagRQlen
	inDiagReq   InetDiagRequest
	inDiagMsg   InetDiagMessage

	procDirentReader *DirentReader
	fdDirentReader   *DirentReader
	fdStat           *syscall.Stat_t
)

func archInit() {
	sockDiagMsgBuffer = make([]byte, OSPageSize)
	unDiagRequestBuffer = make([]byte, SizeOfUnixDiagRequest)
	inDiagRequestBuffer = make([]byte, SizeOfInetDiagRequest)
	fileContentBuffer = bytes.NewBuffer(make([]byte, OSPageSize))

	procDirentReader = NewDirentReader()
	fdDirentReader = NewDirentReader()
	fdStat = new(syscall.Stat_t)

	var err error
	if err = KVer.Get(); err != nil {
		panic(err)
	}
}
