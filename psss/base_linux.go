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

	procDirentHandler *DirentHandler
	fdDirentHandler   *DirentHandler
	fdStat_t          *syscall.Stat_t
)

func archInit() {
	sockDiagMsgBuffer = make([]byte, OSPageSize)
	unDiagRequestBuffer = make([]byte, SizeOfUnixDiagRequest)
	inDiagRequestBuffer = make([]byte, SizeOfInetDiagRequest)
	fileContentBuffer = bytes.NewBuffer(make([]byte, OSPageSize))

	procDirentHandler = NewDirentHandler()
	fdDirentHandler = NewDirentHandler()
	fdStat_t = new(syscall.Stat_t)

	var err error
	if err = KVer.Get(); err != nil {
		panic(err)
	}
}
