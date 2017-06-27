package net

import (
	"golang.org/x/sys/unix"
)

const (
	SsUNKNOWN uint8 = iota
	SsESTAB
	SsSYNSENT
	SsSYNRECV
	SsFINWAIT1
	SsFINWAIT2
	SsTIMEWAIT
	SsUNCONN
	SsCLOSEWAIT
	SsLASTACK
	SsLISTEN
	SsCLOSING
	SsMAX
)

const (
	SOCK_STREAM    = 1
	SOCK_DGRAM     = 2
	SOCK_RAW       = 3
	SOCK_RDM       = 4
	SOCK_SEQPACKET = 5
	SOCK_DCCP      = 6
	SOCK_PACKET    = 10
)

var (
	Sstate = []string{
		"UNKNOWN",
		"ESTAB",
		"SYN-SENT",
		"SYN-RECV",
		"FIN-WAIT-1",
		"FIN-WAIT-2",
		"TIME-WAIT",
		"UNCONN",
		"CLOSE-WAIT",
		"LAST-ACK",
		"LISTEN",
		"CLOSING",
		"MAX",
	}

	SstateActive = map[uint8]bool{
		SsUNKNOWN:   false,
		SsESTAB:     true,
		SsSYNSENT:   false,
		SsSYNRECV:   false,
		SsFINWAIT1:  false,
		SsFINWAIT2:  false,
		SsTIMEWAIT:  false,
		SsUNCONN:    false,
		SsCLOSEWAIT: false,
		SsLASTACK:   false,
		SsLISTEN:    true,
		SsCLOSING:   false,
		SsMAX:       false,
	}

	SstateListen = map[uint8]bool{
		SsUNKNOWN:   false,
		SsESTAB:     false,
		SsSYNSENT:   false,
		SsSYNRECV:   false,
		SsFINWAIT1:  false,
		SsFINWAIT2:  false,
		SsTIMEWAIT:  false,
		SsUNCONN:    true,
		SsCLOSEWAIT: false,
		SsLASTACK:   false,
		SsLISTEN:    true,
		SsCLOSING:   false,
		SsMAX:       false,
	}

	SocketType = map[uint8]string{
		SOCK_STREAM:    "str",
		SOCK_DGRAM:     "dgr",
		SOCK_RAW:       "raw",
		SOCK_RDM:       "rdm",
		SOCK_SEQPACKET: "seq",
		SOCK_DCCP:      "dccp",
		SOCK_PACKET:    "pack",
	}

	UnixSstate = []uint8{SsUNCONN, SsSYNSENT, SsESTAB, SsCLOSING}
)

type SockStatUnix struct {
	Msg      UnixDiagMessage
	Name     string
	VFS      UnixDiagVFS
	Peer     uint32
	Icons    []uint32
	RQlen    UnixDiagRQlen
	Meminfo  []uint32
	Shutdown uint8
}

type SockStatInet struct {
	Msg       InetDiagMessage
	TCPInfo   unix.TCPInfo
	VegasInfo TCPVegasInfo
	CONG      []byte
	SKMeminfo []uint32
	Shutdown  uint8
}
