package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
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
	sockAddrNl := unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
	}
	inDiagReq := InetDiagRequest{
		Header: unix.NlMsghdr{
			Type:  SOCK_DIAG_BY_FAMILY,
			Flags: unix.NLM_F_DUMP | unix.NLM_F_REQUEST,
		},
		Request: InetDiagReq{
			SdiagFamily:   af,
			SdiagProtocol: protocal,
			IdiagExt:      exts,
			IdiagStates:   states,
		},
	}
	inDiagReq.Header.Len = uint32(unsafe.Sizeof(inDiagReq))
	p := make([]byte, unsafe.Sizeof(inDiagReq))
	*(*InetDiagRequest)(unsafe.Pointer(&p[0])) = inDiagReq
	if err = unix.Sendmsg(skfd, p, nil, &sockAddrNl, 0); err != nil {
		return -1, err
	}
	return skfd, nil
}

func RecvInetDiagMsgMulti(skfd int, records map[uint32]*GenericRecord) (err error) {
	var (
		n      int
		cursor int
		nlAttr unix.NlAttr
	)
	p := make([]byte, os.Getpagesize())
	for {
		if n, _, _, _, err = unix.Recvmsg(skfd, p, nil, unix.MSG_PEEK); err != nil {
			return err
		}
		if n < len(p) {
			break
		}
		p = make([]byte, 2*len(p))
	}
	if n, _, _, _, err = unix.Recvmsg(skfd, p, nil, 0); err != nil {
		return err
	}
	p = p[:n]
	raw, err := syscall.ParseNetlinkMessage(p)
	if err != nil {
		return err
	}
	for _, v := range raw {
		record := NewGenericRecord()
		if v.Header.Type == unix.NLMSG_DONE {
			return ErrorDone
		}
		msg := *(*InetDiagMessage)(unsafe.Pointer(&v.Data[:SizeOfInetDiagMsg][0]))
		switch msg.IdiagFamily {
		case unix.AF_INET:
			record.LocalAddr.Host, _ = IPv4HexToString(strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagSrc[0]), "0x"))
			record.RemoteAddr.Host, _ = IPv4HexToString(strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagDst[0]), "0x"))
		case unix.AF_INET6:
			record.LocalAddr.Host, _ = IPv6HexToString(
				strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagSrc[0]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagSrc[1]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagSrc[2]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagSrc[3]), "0x"),
			)
			record.RemoteAddr.Host, _ = IPv6HexToString(
				strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagDst[0]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagDst[1]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagDst[2]), "0x") +
					strings.TrimPrefix(fmt.Sprintf("%08x", msg.ID.IdiagDst[3]), "0x"),
			)
		}
		record.LocalAddr.Port = fmt.Sprintf("%d", (msg.ID.IdiagSport&0xff00)>>8+(msg.ID.IdiagSport&0xff)<<8)
		record.RemoteAddr.Port = fmt.Sprintf("%d", (msg.ID.IdiagDport&0xff00)>>8+(msg.ID.IdiagDport&0xff)<<8)
		record.Status = msg.IdiagState
		record.RxQueue = msg.IdiagRqueue
		record.TxQueue = msg.IdiagWqueue
		record.Timer = int(msg.IdiagTimer)
		record.Timeout = int(msg.IdiagExpires)
		record.Retransmit = int(msg.IdiagRetrans)
		record.UID = uint64(msg.IdiagUid)
		record.Inode = msg.IdiagInode
		record.RefCount = int(msg.ID.IdiagIF)
		record.SK = uint64(msg.ID.IdiagCookie[1])<<32 | uint64(msg.ID.IdiagCookie[0])
		cursor = SizeOfInetDiagMsg
		for cursor+4 < len(v.Data) {
			for v.Data[cursor] == byte(0) {
				cursor++
			}
			nlAttr = *(*unix.NlAttr)(unsafe.Pointer(&v.Data[cursor : cursor+unix.SizeofNlAttr][0]))
			switch nlAttr.Type {
			case INET_DIAG_MEMINFO:
				// meminfo := *(*InetDiagMeminfo)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_INFO:
				record.TCPInfo = *(*TCPInfo)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_VEGASINFO:
				record.VegasInfo = *(*TCPVegasInfo)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_CONG:
				record.CONG = make([]byte, 0)
				record.CONG = append(record.CONG, v.Data[cursor+unix.SizeofNlAttr:cursor+int(nlAttr.Len)]...)
			case INET_DIAG_TOS:
				// tos := *(*uint8)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_TCLASS:
				// tclass := *(*uint8)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_SKMEMINFO:
				if nlAttr.Len > 4 {
					record.Meminfo = make([]uint32, 0, 8)
					for i := cursor + unix.SizeofNlAttr; i < cursor+int(nlAttr.Len); i = i + 4 {
						record.Meminfo = append(record.Meminfo, *(*uint32)(unsafe.Pointer(&v.Data[i : i+4][0])))
					}
				}
			case INET_DIAG_SHUTDOWN:
				// shutdown := *(*uint8)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			default:
				fmt.Println("invalid NlAttr Type")
			}
			cursor += int(nlAttr.Len)
		}
		records[record.Inode] = &record
	}
	return nil
}

func RecvInetDiagMsgAll(skfd int) (records map[uint32]*GenericRecord) {
	records = make(map[uint32]*GenericRecord)
	for {
		if err := RecvInetDiagMsgMulti(skfd, records); err != nil {
			if err == ErrorDone {
				break
			}
			continue
		}
	}
	return records
}
