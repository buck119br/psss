package net

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
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
	ErrorDone = fmt.Errorf("Done")
)

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
	sockAddrNl := unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
	}
	unDiagReq := UnixDiagRequest{
		Header: unix.NlMsghdr{
			Type:  SOCK_DIAG_BY_FAMILY,
			Flags: unix.NLM_F_DUMP | unix.NLM_F_REQUEST,
		},
		Request: UnixDiagReq{
			SdiagFamily: unix.AF_UNIX,
			UdiagStates: states,
			UdiagShow:   show,
		},
	}
	unDiagReq.Header.Len = uint32(unsafe.Sizeof(unDiagReq))
	p := make([]byte, unsafe.Sizeof(unDiagReq))
	*(*UnixDiagRequest)(unsafe.Pointer(&p[0])) = unDiagReq
	if err = unix.Sendmsg(skfd, p, nil, &sockAddrNl, 0); err != nil {
		return -1, err
	}
	return skfd, nil
}

func RecvUnixDiagMsgMulti(skfd int) (multi []SockStatUnix, err error) {
	var (
		n      int
		cursor int
		nlAttr unix.NlAttr
		ssu    SockStatUnix
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
		if v.Header.Type == unix.NLMSG_DONE {
			return multi, ErrorDone
		}
		ssu.Msg = *(*UnixDiagMessage)(unsafe.Pointer(&v.Data[:SizeOfUnixDiagMsg][0]))
		cursor = SizeOfUnixDiagMsg
		for cursor+4 < len(v.Data) {
			for v.Data[cursor] == byte(0) {
				cursor++
			}
			nlAttr = *(*unix.NlAttr)(unsafe.Pointer(&v.Data[cursor : cursor+unix.SizeofNlAttr][0]))
			switch nlAttr.Type {
			case UNIX_DIAG_NAME:
				ssu.Name = string(v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)])
			case UNIX_DIAG_VFS:
				ssu.VFS = *(*UnixDiagVFS)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case UNIX_DIAG_PEER:
				ssu.Peer = *(*uint32)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case UNIX_DIAG_ICONS:
				if nlAttr.Len > 4 {
					ssu.Icons = make([]uint32, 0)
					for i := cursor + unix.SizeofNlAttr; i < cursor+int(nlAttr.Len); i = i + 4 {
						ssu.Icons = append(ssu.Icons, *(*uint32)(unsafe.Pointer(&v.Data[i : i+4][0])))
					}
				}
			case UNIX_DIAG_RQLEN:
				ssu.RQlen = *(*UnixDiagRQlen)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case UNIX_DIAG_MEMINFO:
				if nlAttr.Len > 4 {
					ssu.Meminfo = make([]uint32, 0, 8)
					for i := cursor + unix.SizeofNlAttr; i < cursor+int(nlAttr.Len); i = i + 4 {
						ssu.Meminfo = append(ssu.Meminfo, *(*uint32)(unsafe.Pointer(&v.Data[i : i+4][0])))
					}
				}
			case UNIX_DIAG_SHUTDOWN:
				ssu.Shutdown = *(*uint8)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			default:
				return nil, fmt.Errorf("invalid NlAttr Type")
			}
			cursor += int(nlAttr.Len)
		}
		multi = append(multi, ssu)
	}
	return multi, nil
}

func RecvUnixDiagMsgAll(skfd int) (list []SockStatUnix, err error) {
	var multi []SockStatUnix
	for {
		if multi, err = RecvUnixDiagMsgMulti(skfd); err != nil {
			if err == ErrorDone {
				break
			}
			continue
		}
		list = append(list, multi...)
	}
	return list, nil
}
