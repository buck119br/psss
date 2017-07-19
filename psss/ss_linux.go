// +build linux

package psss

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	SOCK_DIAG_BY_FAMILY = 20

	SizeOfUnixDiagRequest = 40
	SizeOfUnixDiagMsg     = 16
	SizeOfInetDiagRequest = 72
	SizeOfInetDiagMsg     = 72
)

const (
	INET_DIAG_NONE = iota
	INET_DIAG_MEMINFO
	INET_DIAG_INFO
	INET_DIAG_VEGASINFO
	INET_DIAG_CONG
	INET_DIAG_TOS
	INET_DIAG_TCLASS
	INET_DIAG_SKMEMINFO
	INET_DIAG_SHUTDOWN
	INET_DIAG_DCTCPINFO
	INET_DIAG_PROTOCOL /* response attribute only */
	INET_DIAG_SKV6ONLY
	INET_DIAG_LOCALS
	INET_DIAG_PEERS
	INET_DIAG_PAD
	INET_DIAG_MARK
	INET_DIAG_BBRINFO
	INET_DIAG_MAX
)

const (
	TCPI_OPT_TIMESTAMPS = 1
	TCPI_OPT_SACK       = 2
	TCPI_OPT_WSCALE     = 4
	TCPI_OPT_ECN        = 8  /* ECN was negociated at TCP session init */
	TCPI_OPT_ECN_SEEN   = 16 /* we received at least one packet with ECT */
	TCPI_OPT_SYN_DATA   = 32 /* SYN-ACK acked data in SYN sent or rcvd */
)

const (
	UDIAG_SHOW_NAME    = 0x00000001 /* show name (not path) */
	UDIAG_SHOW_VFS     = 0x00000002 /* show VFS inode info */
	UDIAG_SHOW_PEER    = 0x00000004 /* show peer socket info */
	UDIAG_SHOW_ICONS   = 0x00000008 /* show pending connections */
	UDIAG_SHOW_RQLEN   = 0x00000010 /* show skb receive queue len */
	UDIAG_SHOW_MEMINFO = 0x00000020 /* show memory info of a socket */
)

const (
	UNIX_DIAG_NAME = iota
	UNIX_DIAG_VFS
	UNIX_DIAG_PEER
	UNIX_DIAG_ICONS
	UNIX_DIAG_RQLEN
	UNIX_DIAG_MEMINFO
	UNIX_DIAG_SHUTDOWN
	UNIX_DIAG_MAX
)

var (
	procFilePath = map[string]string{
		"sockstat4": "/proc/net/sockstat",
		"sockstat6": "/proc/net/sockstat6",
		"TCP4":      "/proc/net/tcp",
		"TCP6":      "/proc/net/tcp6",
		"UDP4":      "/proc/net/udp",
		"UDP6":      "/proc/net/udp6",
		"RAW4":      "/proc/net/raw",
		"RAW6":      "/proc/net/raw6",
		"Unix":      "/proc/net/unix",
	}

	UnixSstate = []uint8{
		SsUNCONN,
		SsSYNSENT,
		SsESTAB,
		SsCLOSING,
	}
)

type InetDiagSockID struct {
	IdiagSport  uint16
	IdiagDport  uint16
	IdiagSrc    [4]uint32
	IdiagDst    [4]uint32
	IdiagIF     uint32
	IdiagCookie [2]uint32
}

type InetDiagReq struct {
	SdiagFamily   uint8
	SdiagProtocol uint8
	IdiagExt      uint8
	Pad           uint8
	IdiagStates   uint32
	ID            InetDiagSockID
}

type InetDiagRequest struct {
	Header  unix.NlMsghdr
	Request InetDiagReq
}

type InetDiagMessage struct {
	IdiagFamily  uint8
	IdiagState   uint8
	IdiagTimer   uint8
	IdiagRetrans uint8
	ID           InetDiagSockID
	IdiagExpires uint32
	IdiagRqueue  uint32
	IdiagWqueue  uint32
	IdiagUid     uint32
	IdiagInode   uint32
}

type TCPInfo struct {
	State           uint8
	Ca_state        uint8
	Retransmits     uint8
	Probes          uint8
	Backoff         uint8
	Options         uint8
	Pad_cgo_0       [2]byte
	Rto             uint32
	Ato             uint32
	Snd_mss         uint32
	Rcv_mss         uint32
	Unacked         uint32
	Sacked          uint32
	Lost            uint32
	Retrans         uint32
	Fackets         uint32
	Last_data_sent  uint32
	Last_ack_sent   uint32
	Last_data_recv  uint32
	Last_ack_recv   uint32
	Pmtu            uint32
	Rcv_ssthresh    uint32
	Rtt             uint32
	Rttvar          uint32
	Snd_ssthresh    uint32
	Snd_cwnd        uint32
	Advmss          uint32
	Reordering      uint32
	Rcv_rtt         uint32
	Rcv_space       uint32
	Total_retrans   uint32
	Pacing_rate     uint64
	Max_pacing_rate uint64
	Bytes_acked     uint64 /* RFC4898 tcpEStatsAppHCThruOctetsAcked */
	Bytes_received  uint64 /* RFC4898 tcpEStatsAppHCThruOctetsReceived */
	Segs_out        uint32 /* RFC4898 tcpEStatsPerfSegsOut */
	Segs_in         uint32 /* RFC4898 tcpEStatsPerfSegsIn */
	Notsent_bytes   uint32
	Min_rtt         uint32
	Data_segs_in    uint32 /* RFC4898 tcpEStatsDataSegsIn */
	Data_segs_out   uint32 /* RFC4898 tcpEStatsDataSegsOut */
	Delivery_rate   uint64
	Busy_time       uint64 /* Time (usec) busy sending data */
	Rwnd_limited    uint64 /* Time (usec) limited by receive window */
	Sndbuf_limited  uint64 /* Time (usec) limited by send buffer */
}

type TCPVegasInfo struct {
	Enabled uint32
	Rttcnt  uint32
	Rtt     uint32
	Minrtt  uint32
}

type InetDiagMeminfo struct {
	IdiagRmem uint32
	IdiagWmem uint32
	IdiagFmem uint32
	IdiagTmem uint32
}

func SendInetDiagMsg(af uint8, protocal uint8, exts uint8, states uint32) (skfd int, err error) {
	if skfd, err = unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_SOCK_DIAG); err != nil {
		return -1, err
	}
	sockAddrNl.Family = unix.AF_NETLINK
	inDiagReq.Header.Type = SOCK_DIAG_BY_FAMILY
	inDiagReq.Header.Flags = unix.NLM_F_DUMP | unix.NLM_F_REQUEST
	inDiagReq.Request.SdiagFamily = af
	inDiagReq.Request.SdiagProtocol = protocal
	inDiagReq.Request.IdiagExt = exts
	inDiagReq.Request.IdiagStates = states
	inDiagReq.Header.Len = uint32(unsafe.Sizeof(inDiagReq))
	*(*InetDiagRequest)(unsafe.Pointer(&inDiagRequestBuffer[0])) = inDiagReq
	if err = unix.Sendmsg(skfd, inDiagRequestBuffer, nil, &sockAddrNl, 0); err != nil {
		return -1, err
	}
	return skfd, nil
}

func RecvInetDiagMsgMulti(skfd int) (err error) {
	var cursor int

	for {
		if bytesCounter, _, _, _, err = unix.Recvmsg(skfd, sockDiagMsgBuffer, nil, unix.MSG_PEEK); err != nil {
			return err
		}
		if bytesCounter < len(sockDiagMsgBuffer) {
			break
		}
		sockDiagMsgBuffer = make([]byte, 2*len(sockDiagMsgBuffer))
	}
	if bytesCounter, _, _, _, err = unix.Recvmsg(skfd, sockDiagMsgBuffer, nil, 0); err != nil {
		return err
	}
	raw, err := syscall.ParseNetlinkMessage(sockDiagMsgBuffer[:bytesCounter])
	if err != nil {
		return err
	}
	var si = NewSocketInfo()
	for indexBuffer = range raw {
		si.Reset()
		if raw[indexBuffer].Header.Type == unix.NLMSG_DONE {
			return ErrorDone
		}
		inDiagMsg = *(*InetDiagMessage)(unsafe.Pointer(&raw[indexBuffer].Data[:SizeOfInetDiagMsg][0]))
		switch inDiagMsg.IdiagFamily {
		case unix.AF_INET:
			si.LocalAddr.Host, _ = IPv4HexToString(strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagSrc[0]), "0x"))
			si.RemoteAddr.Host, _ = IPv4HexToString(strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagDst[0]), "0x"))
		case unix.AF_INET6:
			si.LocalAddr.Host, _ = IPv6HexToString(
				strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagSrc[0]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagSrc[1]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagSrc[2]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagSrc[3]), "0x"),
			)
			si.RemoteAddr.Host, _ = IPv6HexToString(
				strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagDst[0]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagDst[1]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagDst[2]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", inDiagMsg.ID.IdiagDst[3]), "0x"),
			)
		}
		si.LocalAddr.Port = fmt.Sprintf("%d", (inDiagMsg.ID.IdiagSport&0xff00)>>8+(inDiagMsg.ID.IdiagSport&0xff)<<8)
		si.RemoteAddr.Port = fmt.Sprintf("%d", (inDiagMsg.ID.IdiagDport&0xff00)>>8+(inDiagMsg.ID.IdiagDport&0xff)<<8)
		si.Status = inDiagMsg.IdiagState
		si.RxQueue = inDiagMsg.IdiagRqueue
		si.TxQueue = inDiagMsg.IdiagWqueue
		si.Timer = int(inDiagMsg.IdiagTimer)
		si.Timeout = int(inDiagMsg.IdiagExpires)
		si.Retransmit = int(inDiagMsg.IdiagRetrans)
		si.UID = uint64(inDiagMsg.IdiagUid)
		si.Inode = inDiagMsg.IdiagInode
		si.RefCount = int(inDiagMsg.ID.IdiagIF)
		si.SK = uint64(inDiagMsg.ID.IdiagCookie[1])<<32 | uint64(inDiagMsg.ID.IdiagCookie[0])
		cursor = SizeOfInetDiagMsg
		for cursor+4 < len(raw[indexBuffer].Data) {
			for raw[indexBuffer].Data[cursor] == byte(0) {
				cursor++
			}
			nlAttr = *(*unix.NlAttr)(unsafe.Pointer(&raw[indexBuffer].Data[cursor : cursor+unix.SizeofNlAttr][0]))
			switch nlAttr.Type {
			case INET_DIAG_MEMINFO:
				// meminfo := *(*InetDiagMeminfo)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_INFO:
				si.TCPInfo = (*TCPInfo)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_VEGASINFO:
				si.VegasInfo = (*TCPVegasInfo)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_CONG:
				si.CONG = make([]byte, 0)
				si.CONG = append(si.CONG, raw[indexBuffer].Data[cursor+unix.SizeofNlAttr:cursor+int(nlAttr.Len)]...)
			case INET_DIAG_TOS:
				// tos := *(*uint8)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_TCLASS:
				// tclass := *(*uint8)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_SKMEMINFO:
				if nlAttr.Len > 4 {
					si.Meminfo = make([]uint32, 0, 8)
					for j := cursor + unix.SizeofNlAttr; j < cursor+int(nlAttr.Len); j = j + 4 {
						si.Meminfo = append(si.Meminfo, *(*uint32)(unsafe.Pointer(&raw[indexBuffer].Data[j : j+4][0])))
					}
				}
			case INET_DIAG_SHUTDOWN:
				// shutdown := *(*uint8)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			default:
			}
			cursor += int(nlAttr.Len)
		}
		if FlagProcess {
			si.SetUpRelation()
		}
		SocketInfoChan <- *si
	}
	return nil
}

func RecvInetDiagMsgAll(skfd int) {
	defer func() {
		SocketInfoChan <- SocketInfo{IsEnd: true}
	}()
	for {
		if err := RecvInetDiagMsgMulti(skfd); err != nil {
			if err == ErrorDone {
				return
			}
			continue
		}
	}
}

func GenericInetRead(protocal, af int) (sis map[uint32]SocketInfo, err error) {
	var (
		ipproto uint8
		exts    uint8
		skfd    int
	)
	switch protocal {
	case ProtocalTCP:
		ipproto = unix.IPPROTO_TCP
		if FlagInfo {
			exts |= 1 << (INET_DIAG_INFO - 1)
			exts |= 1 << (INET_DIAG_VEGASINFO - 1)
			exts |= 1 << (INET_DIAG_CONG - 1)
		}
	case ProtocalUDP:
		ipproto = unix.IPPROTO_UDP
	case ProtocalRAW:
		ipproto = unix.IPPROTO_RAW
	default:
		return nil, fmt.Errorf("invalid protocal:[%d]", protocal)
	}
	if FlagMemory {
		exts |= 1 << (INET_DIAG_SKMEMINFO - 1)
	}
	if skfd, err = SendInetDiagMsg(uint8(af), ipproto, exts, SsFilter); err != nil {
		goto readProc
	}
	defer unix.Close(skfd)

	sis = make(map[uint32]SocketInfo)
	go RecvInetDiagMsgAll(skfd)
	for si := range SocketInfoChan {
		if si.IsEnd {
			return sis, nil
		}
		sis[si.Inode] = si
	}

readProc:
	var (
		procPath    string
		file        *os.File
		line        string
		fields      []string
		fieldsIndex int
		stringBuff  []string
		int64Buffer int64
	)
	sis = make(map[uint32]SocketInfo)

	switch protocal {
	case ProtocalTCP:
		procPath = "TCP"
	case ProtocalUDP:
		procPath = "UDP"
	case ProtocalRAW:
		procPath = "RAW"
	}
	switch af {
	case unix.AF_INET:
		procPath += "4"
	case unix.AF_INET6:
		procPath += "6"
	}
	if file, err = os.Open(procFilePath[procPath]); err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return sis, err
		}
		line = scanner.Text()
		fields = strings.Fields(line)
		if fields[0] == "sl" {
			continue
		}
		si := NewSocketInfo()
		// Local address
		fieldsIndex = 1
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		switch af {
		case unix.AF_INET:
			si.LocalAddr.Host, err = IPv4HexToString(stringBuff[0])
		case unix.AF_INET6:
			si.LocalAddr.Host, err = IPv6HexToString(stringBuff[0])
		}
		if err != nil {
			continue
		}
		if int64Buffer, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			continue
		}
		si.LocalAddr.Port = fmt.Sprintf("%d", int64Buffer)
		if MaxLocalAddrLength < len(si.LocalAddr.String()) {
			MaxLocalAddrLength = len(si.LocalAddr.String())
		}
		fieldsIndex++
		// Remote address
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		switch af {
		case unix.AF_INET:
			si.RemoteAddr.Host, err = IPv4HexToString(stringBuff[0])
		case unix.AF_INET6:
			si.RemoteAddr.Host, err = IPv6HexToString(stringBuff[0])
		}
		if err != nil {
			continue
		}
		if int64Buffer, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			continue
		}
		si.RemoteAddr.Port = fmt.Sprintf("%d", int64Buffer)
		if MaxRemoteAddrLength < len(si.RemoteAddr.String()) {
			MaxRemoteAddrLength = len(si.RemoteAddr.String())
		}
		fieldsIndex++
		// Status
		if int64Buffer, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			continue
		}
		si.Status = uint8(int64Buffer)
		if SsFilter&(1<<si.Status) == 0 {
			continue
		}
		fieldsIndex++
		// TxQueue:RxQueue
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if int64Buffer, err = strconv.ParseInt(stringBuff[0], 16, 64); err != nil {
			continue
		}
		si.TxQueue = uint32(int64Buffer)
		if int64Buffer, err = strconv.ParseInt(stringBuff[1], 16, 64); err != nil {
			continue
		}
		si.RxQueue = uint32(int64Buffer)
		fieldsIndex++
		// Timer:TmWhen
		stringBuff = strings.Split(fields[fieldsIndex], ":")
		if int64Buffer, err = strconv.ParseInt(stringBuff[0], 16, 32); err != nil {
			continue
		}
		si.Timer = int(int64Buffer)
		if si.Timer > 4 {
			si.Timer = 5
		}
		if int64Buffer, err = strconv.ParseInt(stringBuff[1], 16, 32); err != nil {
			continue
		}
		si.Timeout = int(int64Buffer)
		si.Timeout = (si.Timeout*1000 + int(SC_CLK_TCK) - 1) / int(SC_CLK_TCK)
		fieldsIndex++
		// Retransmit
		if int64Buffer, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			continue
		}
		si.Retransmit = int(int64Buffer)
		fieldsIndex++
		if si.UID, err = strconv.ParseUint(fields[fieldsIndex], 10, 64); err != nil {
			continue
		}
		fieldsIndex++
		if si.Probes, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			continue
		}
		fieldsIndex++
		if int64Buffer, err = strconv.ParseInt(fields[fieldsIndex], 10, 64); err != nil {
			continue
		}
		si.Inode = uint32(int64Buffer)
		fieldsIndex++
		if si.RefCount, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
			continue
		}
		fieldsIndex++
		if si.SK, err = strconv.ParseUint(fields[fieldsIndex], 16, 64); err != nil {
			continue
		}
		switch protocal {
		case ProtocalTCP:
			if len(fields) > 12 {
				fieldsIndex++
				if si.RTO, err = strconv.ParseFloat(fields[fieldsIndex], 64); err != nil {
					continue
				}
				fieldsIndex++
				if si.ATO, err = strconv.ParseFloat(fields[fieldsIndex], 64); err != nil {
					continue
				}
				si.ATO = si.ATO / float64(SC_CLK_TCK)
				fieldsIndex++
				if si.QACK, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
					continue
				}
				si.QACK = si.QACK / 2
				fieldsIndex++
				if si.CongestionWindow, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
					continue
				}
				fieldsIndex++
				if si.SlowStartThreshold, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
					continue
				}
			} else {
				si.RTO = 0
				si.ATO = 0
				si.QACK = 0
				si.CongestionWindow = 2
				si.SlowStartThreshold = -1
			}
			if si.SlowStartThreshold == -1 {
				si.SlowStartThreshold = 0
			}
			if si.RTO == float64(3*SC_CLK_TCK) {
				si.RTO = 0
			}
			if si.Timer != 1 {
				si.Retransmit = si.Probes
			}
		case ProtocalUDP, ProtocalRAW:
			fieldsIndex++
			if si.Drops, err = strconv.Atoi(fields[fieldsIndex]); err != nil {
				continue
			}
		}
		if len(fields) > 17 {
			si.Opt = fields[17:]
		}
		if FlagProcess {
			si.SetUpRelation()
		}
		sis[si.Inode] = *si
	}
	return
}

type UnixDiagReq struct {
	SdiagFamily   uint8
	SdiagProtocol uint8
	Pad           uint16
	UdiagStates   uint32
	UdiagIno      uint32
	UdiagShow     uint32
	UdiagCookie   [2]uint32
}

type UnixDiagRequest struct {
	Header  unix.NlMsghdr
	Request UnixDiagReq
}

type UnixDiagMessage struct {
	UdiagFalimy uint8
	UdiagType   uint8
	UdiagState  uint8
	Pad         uint8
	UdiagIno    uint32
	UdiagCookie [2]uint32
}

type UnixDiagVFS struct {
	Dev uint32
	Ino uint32
}

type UnixDiagRQlen struct {
	RQ uint32
	WQ uint32
}

// Make sure the caller of the function will close skfd
func SendUnixDiagMsg(states uint32, show uint32) (skfd int, err error) {
	if skfd, err = unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_SOCK_DIAG); err != nil {
		return -1, err
	}
	sockAddrNl.Family = unix.AF_NETLINK
	unDiagReq.Header.Type = SOCK_DIAG_BY_FAMILY
	unDiagReq.Header.Flags = unix.NLM_F_DUMP | unix.NLM_F_REQUEST
	unDiagReq.Request.SdiagFamily = unix.AF_UNIX
	unDiagReq.Request.UdiagStates = states
	unDiagReq.Request.UdiagShow = show
	unDiagReq.Header.Len = uint32(unsafe.Sizeof(unDiagReq))
	*(*UnixDiagRequest)(unsafe.Pointer(&unDiagRequestBuffer[0])) = unDiagReq
	if err = unix.Sendmsg(skfd, unDiagRequestBuffer, nil, &sockAddrNl, 0); err != nil {
		return -1, err
	}
	return skfd, nil
}

func RecvUnixDiagMsgMulti(skfd int) (err error) {
	var cursor int

	for {
		if bytesCounter, _, _, _, err = unix.Recvmsg(skfd, sockDiagMsgBuffer, nil, unix.MSG_PEEK); err != nil {
			return err
		}
		if bytesCounter < len(sockDiagMsgBuffer) {
			break
		}
		sockDiagMsgBuffer = make([]byte, 2*len(sockDiagMsgBuffer))
	}
	if bytesCounter, _, _, _, err = unix.Recvmsg(skfd, sockDiagMsgBuffer, nil, 0); err != nil {
		return err
	}
	raw, err := syscall.ParseNetlinkMessage(sockDiagMsgBuffer[:bytesCounter])
	if err != nil {
		return err
	}
	si := NewSocketInfo()
	for indexBuffer = range raw {
		si.Reset()
		if raw[indexBuffer].Header.Type == unix.NLMSG_DONE {
			return ErrorDone
		}
		unDiagMsg = *(*UnixDiagMessage)(unsafe.Pointer(&raw[indexBuffer].Data[:SizeOfUnixDiagMsg][0]))
		si.Inode = unDiagMsg.UdiagIno
		si.LocalAddr.Port = fmt.Sprintf("%d", unDiagMsg.UdiagIno)
		si.Status = unDiagMsg.UdiagState
		si.Type = unDiagMsg.UdiagType
		si.SK = uint64(unDiagMsg.UdiagCookie[1])<<32 | uint64(unDiagMsg.UdiagCookie[0])
		cursor = SizeOfUnixDiagMsg
		for cursor+4 < len(raw[indexBuffer].Data) {
			for raw[indexBuffer].Data[cursor] == byte(0) {
				cursor++
			}
			nlAttr = *(*unix.NlAttr)(unsafe.Pointer(&raw[indexBuffer].Data[cursor : cursor+unix.SizeofNlAttr][0]))
			switch nlAttr.Type {
			case UNIX_DIAG_NAME:
				si.LocalAddr.Host = string(raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)])
				if len(si.LocalAddr.Host) == 0 {
					si.LocalAddr.Host = "*"
				}
				if MaxLocalAddrLength < len(si.LocalAddr.String()) {
					MaxLocalAddrLength = len(si.LocalAddr.String())
				}
			case UNIX_DIAG_VFS:
				// vfs := *(*UnixDiagVFS)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case UNIX_DIAG_PEER:
				si.RemoteAddr.Host = "*"
				si.RemoteAddr.Port = fmt.Sprintf("%d", *(*uint32)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0])))
				if MaxRemoteAddrLength < len(si.RemoteAddr.String()) {
					MaxRemoteAddrLength = len(si.RemoteAddr.String())
				}
			case UNIX_DIAG_ICONS:
				// if nlAttr.Len > 4 {
				// 	icons := make([]uint32, 0)
				// 	for j := cursor + unix.SizeofNlAttr; j < cursor+int(nlAttr.Len); j = j + 4 {
				// 		icons = append(icons, *(*uint32)(unsafe.Pointer(&raw[indexBuffer].Data[j : j+4][0])))
				// 	}
				// }
			case UNIX_DIAG_RQLEN:
				unDiagRQlen = *(*UnixDiagRQlen)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
				si.RxQueue = unDiagRQlen.RQ
				si.TxQueue = unDiagRQlen.WQ
			case UNIX_DIAG_MEMINFO:
				if nlAttr.Len > 4 {
					si.Meminfo = make([]uint32, 0, 8)
					for j := cursor + unix.SizeofNlAttr; j < cursor+int(nlAttr.Len); j = j + 4 {
						si.Meminfo = append(si.Meminfo, *(*uint32)(unsafe.Pointer(&raw[indexBuffer].Data[j : j+4][0])))
					}
				}
			case UNIX_DIAG_SHUTDOWN:
				// shutdown := *(*uint8)(unsafe.Pointer(&raw[indexBuffer].Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			default:
				fmt.Println("invalid NlAttr Type")
			}
			cursor += int(nlAttr.Len)
		}
		if FlagProcess {
			si.SetUpRelation()
		}
		SocketInfoChan <- *si
	}
	return nil
}

func RecvUnixDiagMsgAll(skfd int) {
	defer func() {
		SocketInfoChan <- SocketInfo{IsEnd: true}
	}()
	for {
		if err := RecvUnixDiagMsgMulti(skfd); err != nil {
			if err == ErrorDone {
				return
			}
			continue
		}
	}
}

func GenericUnixRead() (sis map[uint32]SocketInfo, err error) {
	skfd, err := SendUnixDiagMsg(SsFilter,
		UDIAG_SHOW_NAME|UDIAG_SHOW_VFS|UDIAG_SHOW_PEER|UDIAG_SHOW_ICONS|UDIAG_SHOW_RQLEN|UDIAG_SHOW_MEMINFO)
	if err != nil {
		goto readProc
	}
	defer unix.Close(skfd)
	sis = make(map[uint32]SocketInfo)
	go RecvUnixDiagMsgAll(skfd)
	for si := range SocketInfoChan {
		if si.IsEnd {
			return sis, nil
		}
		sis[si.Inode] = si
	}

readProc:
	// In this way, so much information cannot get.
	var (
		line        string
		fields      []string
		fieldsIndex int
		flag        int64
	)
	sis = make(map[uint32]SocketInfo)
	file, err := os.Open(procFilePath["Unix"])
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return sis, err
		}
		line = scanner.Text()
		fields = strings.Fields(line)
		if len(fields) < 7 {
			continue
		}
		if fields[0] == "Num" {
			continue
		}
		si := NewSocketInfo()
		// Num: the kernel table slot number.
		fieldsIndex = 0
		if si.SK, err = strconv.ParseUint(strings.Replace(fields[fieldsIndex], ":", "", -1), 16, 64); err != nil {
			continue
		}
		si.RemoteAddr.Host = "*"
		si.RemoteAddr.Port = "Unknown"
		if MaxRemoteAddrLength < len(si.RemoteAddr.String()) {
			MaxRemoteAddrLength = len(si.RemoteAddr.String())
		}
		fieldsIndex++
		// RefCount: the number of users of the socket.
		si.RxQueue = 0
		fieldsIndex++
		// Protocol: currently always 0.
		fieldsIndex++
		// Flags: the internal kernel flags holding the status of the socket.
		if flag, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			continue
		}
		fieldsIndex++
		// Type: the socket type.
		// For SOCK_STREAM sockets, this is 0001; for SOCK_DGRAM sockets, it is 0002; and for SOCK_SEQPACKET sockets, it is 0005.
		if int64Buffer, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			continue
		}
		si.Type = uint8(int64Buffer)
		fieldsIndex++
		// St: the internal state of the socket.
		if int64Buffer, err = strconv.ParseInt(fields[fieldsIndex], 16, 32); err != nil {
			continue
		}
		if flag&(1<<16) != 0 {
			si.Status = SsLISTEN
		} else {
			si.Status = UnixSstate[int(int64Buffer)-1]
		}
		if SsFilter&(1<<si.Status) == 0 {
			continue
		}
		fieldsIndex++
		// Inode
		if int64Buffer, err = strconv.ParseInt(fields[fieldsIndex], 10, 64); err != nil {
			continue
		}
		si.Inode = uint32(int64Buffer)
		si.LocalAddr.Port = fmt.Sprintf("%d", si.Inode)
		// Path: the bound path (if any) of the socket.
		// Sockets in the abstract namespace are included in the list, and are shown with a Path that commences with the character '@'.
		if len(fields) > 7 {
			fieldsIndex++
			si.LocalAddr.Host = fields[fieldsIndex]
		} else {
			si.LocalAddr.Host = "*"
		}
		if MaxLocalAddrLength < len(si.LocalAddr.String()) {
			MaxLocalAddrLength = len(si.LocalAddr.String())
		}
		if FlagProcess {
			si.SetUpRelation()
		}
		sis[si.Inode] = *si
	}
	return sis, nil
}

func GetSocketCount(fields []string) (int, error) {
	for indexBuffer = range fields {
		if fields[indexBuffer] == "inuse" {
			return strconv.Atoi(fields[indexBuffer+1])
		}
	}
	return 0, nil
}

// IPv6:versionFlag = true; IPv4:versionFlag = false
func GenericReadSockstat() (summary map[string]map[string]int, err error) {
	summary = make(map[string]map[string]int)
	for _, pf := range SummaryPF {
		summary[pf] = make(map[string]int)
	}

	var file *os.File
	for _, v := range []string{"sockstat4", "sockstat6"} {
		if file, err = os.Open(procFilePath[v]); err != nil {
			return nil, err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if err = scanner.Err(); err != nil {
				return nil, err
			}
			line := scanner.Text()
			fields := strings.Fields(line)
			switch fields[0] {
			case "sockets:":
				continue
			case "TCP:":
				summary["TCP"][IPv4String], err = GetSocketCount(fields[1:])
			case "TCP6:":
				summary["TCP"][IPv6String], err = GetSocketCount(fields[1:])
			case "UDP:":
				summary["UDP"][IPv4String], err = GetSocketCount(fields[1:])
			case "UDP6:":
				summary["UDP"][IPv6String], err = GetSocketCount(fields[1:])
			case "UDPLITE:":
				summary["UDPLITE"][IPv4String], err = GetSocketCount(fields[1:])
			case "UDPLITE6:":
				summary["UDPLITE"][IPv6String], err = GetSocketCount(fields[1:])
			case "RAW:":
				summary["RAW"][IPv4String], err = GetSocketCount(fields[1:])
			case "RAW6:":
				summary["RAW"][IPv6String], err = GetSocketCount(fields[1:])
			case "FRAG:":
				summary["FRAG"][IPv4String], err = GetSocketCount(fields[1:])
			case "FRAG6:":
				summary["FRAG"][IPv6String], err = GetSocketCount(fields[1:])
			default:
				continue
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return summary, nil
}
