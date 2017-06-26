package net

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	IPPROTO_IP      = 0   /* Dummy protocol for TCP               */
	IPPROTO_ICMP    = 1   /* Internet Control Message Protocol    */
	IPPROTO_IGMP    = 2   /* Internet Group Management Protocol   */
	IPPROTO_IPIP    = 4   /* IPIP tunnels (older KA9Q tunnels use 94) */
	IPPROTO_TCP     = 6   /* Transmission Control Protocol        */
	IPPROTO_EGP     = 8   /* Exterior Gateway Protocol            */
	IPPROTO_PUP     = 12  /* PUP protocol                         */
	IPPROTO_UDP     = 17  /* User Datagram Protocol               */
	IPPROTO_IDP     = 22  /* XNS IDP protocol                     */
	IPPROTO_TP      = 29  /* SO Transport Protocol Class 4        */
	IPPROTO_DCCP    = 33  /* Datagram Congestion Control Protocol */
	IPPROTO_IPV6    = 41  /* IPv6-in-IPv4 tunnelling              */
	IPPROTO_RSVP    = 46  /* RSVP Protocol                        */
	IPPROTO_GRE     = 47  /* Cisco GRE tunnels (rfc 1701,1702)    */
	IPPROTO_ESP     = 50  /* Encapsulation Security Payload protocol */
	IPPROTO_AH      = 51  /* Authentication Header protocol       */
	IPPROTO_MTP     = 92  /* Multicast Transport Protocol         */
	IPPROTO_BEETPH  = 94  /* IP option pseudo header for BEET     */
	IPPROTO_ENCAP   = 98  /* Encapsulation Header                 */
	IPPROTO_PIM     = 103 /* Protocol Independent Multicast       */
	IPPROTO_COMP    = 108 /* Compression Header Protocol          */
	IPPROTO_SCTP    = 132 /* Stream Control Transport Protocol    */
	IPPROTO_UDPLITE = 136 /* UDP-Lite (RFC 3828)                  */
	IPPROTO_MPLS    = 137 /* MPLS in IP (RFC 4023)                */
	IPPROTO_RAW     = 255 /* Raw IP packets                       */
	IPPROTO_MAX
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

func RecvInetDiagMsgMulti(skfd int) (multi []SockStatInet, err error) {
	var (
		n      int
		cursor int
		nlAttr unix.NlAttr
	)
	p := make([]byte, os.Getpagesize())
	for {
		if n, _, _, _, err = unix.Recvmsg(skfd, p, nil, unix.MSG_PEEK); err != nil {
			return nil, err
		}
		if n < len(p) {
			break
		}
		p = make([]byte, 2*len(p))
	}
	if n, _, _, _, err = unix.Recvmsg(skfd, p, nil, 0); err != nil {
		return nil, err
	}
	p = p[:n]
	raw, err := syscall.ParseNetlinkMessage(p)
	if err != nil {
		return nil, err
	}
	for _, v := range raw {
		var ssi SockStatInet
		if v.Header.Type == unix.NLMSG_DONE {
			return multi, ErrorDone
		}
		ssi.Msg = *(*InetDiagMessage)(unsafe.Pointer(&v.Data[:SizeOfInetDiagMsg][0]))
		cursor = SizeOfInetDiagMsg
		for cursor+4 < len(v.Data) {
			for v.Data[cursor] == byte(0) {
				cursor++
			}
			nlAttr = *(*unix.NlAttr)(unsafe.Pointer(&v.Data[cursor : cursor+unix.SizeofNlAttr][0]))
			switch nlAttr.Type {
			case INET_DIAG_INFO:
				ssi.TCPInfo = *(*unix.TCPInfo)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case INET_DIAG_SKMEMINFO:
				if nlAttr.Len > 4 {
					ssi.SKMeminfo = make([]uint32, 0, 8)
					for i := cursor + unix.SizeofNlAttr; i < cursor+int(nlAttr.Len); i = i + 4 {
						ssi.SKMeminfo = append(ssi.SKMeminfo, *(*uint32)(unsafe.Pointer(&v.Data[i : i+4][0])))
					}
				}
			case UNIX_DIAG_SHUTDOWN:
				ssi.Shutdown = *(*uint8)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			default:
				return nil, fmt.Errorf("invalid NlAttr Type")
			}
			cursor += int(nlAttr.Len)
		}
		multi = append(multi, ssi)
	}
	return multi, nil
}

func RecvInetDiagMsgAll(skfd int) (list []SockStatInet, err error) {
	var multi []SockStatInet
	for {
		if multi, err = RecvInetDiagMsgMulti(skfd); err != nil {
			if err == ErrorDone {
				break
			}
			continue
		}
		list = append(list, multi...)
	}
	return list, nil
}
