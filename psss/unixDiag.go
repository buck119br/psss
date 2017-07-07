package psss

import (
	"fmt"
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
	*(*UnixDiagRequest)(unsafe.Pointer(&UnixDiagRequestBuffer[0])) = unDiagReq
	if err = unix.Sendmsg(skfd, UnixDiagRequestBuffer, nil, &sockAddrNl, 0); err != nil {
		return -1, err
	}
	return skfd, nil
}

func RecvUnixDiagMsgMulti(skfd int, records map[uint32]*GenericRecord) (err error) {
	var (
		n      int
		cursor int
		msg    UnixDiagMessage
		nlAttr unix.NlAttr
		rqlen  UnixDiagRQlen
	)
	for {
		if n, _, _, _, err = unix.Recvmsg(skfd, GlobalBuffer, nil, unix.MSG_PEEK); err != nil {
			return err
		}
		if n < len(GlobalBuffer) {
			break
		}
		GlobalBuffer = make([]byte, 2*len(GlobalBuffer))
	}
	if n, _, _, _, err = unix.Recvmsg(skfd, GlobalBuffer, nil, 0); err != nil {
		return err
	}
	raw, err := syscall.ParseNetlinkMessage(GlobalBuffer[:n])
	if err != nil {
		return err
	}
	for _, v := range raw {
		record := NewGenericRecord()
		if v.Header.Type == unix.NLMSG_DONE {
			return ErrorDone
		}
		msg = *(*UnixDiagMessage)(unsafe.Pointer(&v.Data[:SizeOfUnixDiagMsg][0]))
		record.Inode = msg.UdiagIno
		record.LocalAddr.Port = fmt.Sprintf("%d", msg.UdiagIno)
		record.Status = msg.UdiagState
		record.Type = msg.UdiagType
		record.SK = uint64(msg.UdiagCookie[1])<<32 | uint64(msg.UdiagCookie[0])
		cursor = SizeOfUnixDiagMsg
		for cursor+4 < len(v.Data) {
			for v.Data[cursor] == byte(0) {
				cursor++
			}
			nlAttr = *(*unix.NlAttr)(unsafe.Pointer(&v.Data[cursor : cursor+unix.SizeofNlAttr][0]))
			switch nlAttr.Type {
			case UNIX_DIAG_NAME:
				record.LocalAddr.Host = string(v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)])
				if len(record.LocalAddr.Host) == 0 {
					record.LocalAddr.Host = "*"
				}
				if MaxLocalAddrLength < len(record.LocalAddr.String()) {
					MaxLocalAddrLength = len(record.LocalAddr.String())
				}
			case UNIX_DIAG_VFS:
				// vfs := *(*UnixDiagVFS)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
			case UNIX_DIAG_PEER:
				record.RemoteAddr.Host = "*"
				record.RemoteAddr.Port = fmt.Sprintf("%d", *(*uint32)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0])))
				if MaxRemoteAddrLength < len(record.RemoteAddr.String()) {
					MaxRemoteAddrLength = len(record.RemoteAddr.String())
				}
			case UNIX_DIAG_ICONS:
				// if nlAttr.Len > 4 {
				// 	icons := make([]uint32, 0)
				// 	for i := cursor + unix.SizeofNlAttr; i < cursor+int(nlAttr.Len); i = i + 4 {
				// 		icons = append(icons, *(*uint32)(unsafe.Pointer(&v.Data[i : i+4][0])))
				// 	}
				// }
			case UNIX_DIAG_RQLEN:
				rqlen = *(*UnixDiagRQlen)(unsafe.Pointer(&v.Data[cursor+unix.SizeofNlAttr : cursor+int(nlAttr.Len)][0]))
				record.RxQueue = rqlen.RQ
				record.TxQueue = rqlen.WQ
			case UNIX_DIAG_MEMINFO:
				if nlAttr.Len > 4 {
					record.Meminfo = make([]uint32, 0, 8)
					for i := cursor + unix.SizeofNlAttr; i < cursor+int(nlAttr.Len); i = i + 4 {
						record.Meminfo = append(record.Meminfo, *(*uint32)(unsafe.Pointer(&v.Data[i : i+4][0])))
					}
				}
			case UNIX_DIAG_SHUTDOWN:
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

func RecvUnixDiagMsgAll(skfd int) (records map[uint32]*GenericRecord) {
	records = make(map[uint32]*GenericRecord)
	for {
		if err := RecvUnixDiagMsgMulti(skfd, records); err != nil {
			if err == ErrorDone {
				break
			}
			continue
		}
	}
	return
}
