package psss

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	DT_UNKNOWN = 0
	DT_FIFO    = 1
	DT_CHR     = 2
	DT_DIR     = 4
	DT_BLK     = 6
	DT_REG     = 8
	DT_LNK     = 10
	DT_SOCK    = 12
	DT_WHT     = 14
)

type Dirent struct {
	Inode  uint64
	Offset uint64
	Reclen uint16
	Type   byte
	Name   string
	IsEnd  bool
}

type Fd struct {
	Dirent
	Fresh bool
}

type DirentReader struct {
	DataChan       chan Dirent
	Buffer         []byte
	Bufferx        []byte
	NameBuffer     []byte
	InternalDirent Dirent
	ExternalDirent Dirent
	BytesCounter   int
	IndexBuffer    int
	Cursor         int
}

func NewDirentReader() *DirentReader {
	d := new(DirentReader)
	d.DataChan = make(chan Dirent)
	d.Buffer = make([]byte, OSPageSize)
	d.Bufferx = make([]byte, OSPageSize)
	d.NameBuffer = make([]byte, 256)
	return d
}

func (d *DirentReader) Scan(fd *os.File) {
	defer func() {
		d.InternalDirent.IsEnd = true
		d.DataChan <- d.InternalDirent
	}()
	d.Bufferx = d.Bufferx[0:0]
	var err error
	for {
		if d.BytesCounter, err = unix.Getdents(int(fd.Fd()), d.Buffer); err != nil {
			return
		}
		if d.BytesCounter == 0 {
			break
		}
		d.Bufferx = append(d.Bufferx, d.Buffer[:d.BytesCounter]...)
	}
	d.Cursor = 0
	for d.Cursor < len(d.Bufferx)-1 {
		d.InternalDirent.Inode = *(*uint64)(unsafe.Pointer(&d.Bufferx[d.Cursor : d.Cursor+8][0]))
		d.InternalDirent.Offset = *(*uint64)(unsafe.Pointer(&d.Bufferx[d.Cursor+8 : d.Cursor+16][0]))
		d.InternalDirent.Reclen = *(*uint16)(unsafe.Pointer(&d.Bufferx[d.Cursor+16 : d.Cursor+18][0]))
		d.InternalDirent.Type = *(*byte)(unsafe.Pointer(&d.Bufferx[d.Cursor+18 : d.Cursor+19][0]))
		d.NameBuffer = d.NameBuffer[0:0]
		for d.IndexBuffer = d.Cursor + 19; d.IndexBuffer < d.Cursor+int(d.InternalDirent.Reclen); d.IndexBuffer++ {
			if d.Bufferx[d.IndexBuffer] == byte(0) {
				break
			}
			d.NameBuffer = append(d.NameBuffer, d.Bufferx[d.IndexBuffer])
		}
		d.Cursor += int(d.InternalDirent.Reclen)
		d.InternalDirent.Name = string(d.NameBuffer[:len(d.NameBuffer)])
		if d.InternalDirent.Name == "." || d.InternalDirent.Name == ".." {
			continue
		}
		d.InternalDirent.IsEnd = false
		d.DataChan <- d.InternalDirent
	}
}

// definition comes from http://man7.org/linux/man-pages/man5/proc.5.html
type MountInfo struct {
	ID             uint64 // a unique ID for the mount (may be reused after umount(2)).
	ParentID       uint64 // the ID of the parent mount (or of self for the root of this mount namespace's mount tree). If a new mount is stacked on top of a previous existing mount (so that it hides the existing mount) at pathname P, then the parent of the new mount is the previous mount at that location.  Thus, when looking at all the mounts stacked at a particular location, the top-most mount is the one that is not the parent of any other mount at the same location. (Note, however, that this top-most mount will be accessible only if the longest path subprefix of P that is a mount point is not itself hidden by a stacked mount.) If the parent mount point lies outside the process's root directory (see chroot(2)), the ID shown here won't have a corresponding record in mountinfo whose mount ID (field 1) matches this parent mount ID (because mount points that lie outside the process's root directory are not shown in mountinfo). As a special case of this point, the process's root mount point may have a parent mount (for the initramfs filesystem) that lies outside the process's root directory, and an entry for that mount point will not appear in mountinfo.
	DiskMajorNum   uint64 // the value of st_dev for files on this filesystem (see stat(2)).
	DiskMinorNum   uint64
	FileSystemRoot string // the pathname of the directory in the filesystem which forms the root of this mount.
	MountPoint     string // the pathname of the mount point relative to the process's root directory.
	MountOptions   string // per-mount options (see mount(2)).
	OptionalFields string // zero or more fields of the form "tag[:value]"; see below.
	FilesystemType string // the filesystem type in the form "type[.subtype]".
	MountSource    string // filesystem-specific information or "none".
	SuperOptions   string // per-superblock options (see mount(2)).
}

func (mi *MountInfo) Parse(raw string) (err error) {
	fields := strings.Fields(raw)
	for i, v := range fields {
		switch i {
		case 0:
			if mi.ID, err = strconv.ParseUint(v, 10, 64); err != nil {
				return fmt.Errorf("parse id error:[%v]", err)
			}
		case 1:
			if mi.ParentID, err = strconv.ParseUint(v, 10, 64); err != nil {
				return fmt.Errorf("parse parent id error:[%v]", err)
			}
		case 2:
			if mi.DiskMajorNum, err = strconv.ParseUint(strings.Split(v, ":")[0], 10, 64); err != nil {
				return fmt.Errorf("parse parent id error:[%v]", err)
			}
			if mi.DiskMinorNum, err = strconv.ParseUint(strings.Split(v, ":")[1], 10, 64); err != nil {
				return fmt.Errorf("parse parent id error:[%v]", err)
			}
		case 3:
			mi.FileSystemRoot = v
		case 4:
			mi.MountPoint = v
		case 5:
			mi.MountOptions = v
		case 6:
			mi.OptionalFields = v
		case 7:
		case 8:
			mi.FilesystemType = v
		case 9:
			mi.MountSource = v
		case 10:
			mi.SuperOptions = v
		default:
			return fmt.Errorf("invalid field:[%s]", v)
		}
	}
	return nil
}

type MountInfos []*MountInfo

func NewMountInfos() MountInfos {
	return make([]*MountInfo, 0)
}

func (mis *MountInfos) Get() error {
	fd, err := os.Open(ProcRoot + "/self/mountinfo")
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		mi := new(MountInfo)
		if err = mi.Parse(scanner.Text()); err != nil {
			fmt.Printf("mount info parse error:[%v]\n", err)
			continue
		}
		*mis = append(*mis, mi)
	}
	return nil
}
