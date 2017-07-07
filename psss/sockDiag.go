package psss

const (
	SOCK_DIAG_BY_FAMILY = 20

	SizeOfUnixDiagRequest = 40
	SizeOfUnixDiagMsg     = 16
	SizeOfInetDiagRequest = 72
	SizeOfInetDiagMsg     = 72
)

const (
	SK_MEMINFO_RMEM_ALLOC = iota
	SK_MEMINFO_RCVBUF
	SK_MEMINFO_WMEM_ALLOC
	SK_MEMINFO_SNDBUF
	SK_MEMINFO_FWD_ALLOC
	SK_MEMINFO_WMEM_QUEUED
	SK_MEMINFO_OPTMEM
	SK_MEMINFO_BACKLOG
	SK_MEMINFO_DROPS
	SK_MEMINFO_VARS
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

	SocketType = map[uint8]string{
		SOCK_STREAM:    "str",
		SOCK_DGRAM:     "dgr",
		SOCK_RAW:       "raw",
		SOCK_RDM:       "rdm",
		SOCK_SEQPACKET: "seq",
		SOCK_DCCP:      "dccp",
		SOCK_PACKET:    "pack",
	}

	UnixSstate = []uint8{
		SsUNCONN,
		SsSYNSENT,
		SsESTAB,
		SsCLOSING,
	}
)
