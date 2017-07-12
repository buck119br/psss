package psss

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Dirent struct {
	Inode  uint64
	Offset uint64
	Reclen uint16
	Type   byte
	Name   string
}

func ReadDirents(fd *os.File) (dirents map[Dirent]bool, err error) {
	dentBufferx = dentBufferx[0:0]
	for {
		if bytesCounter, err = unix.Getdents(fd, dentBuffer); err != nil {
			return
		}
		runtime.KeepAlive(fd)
		dentBufferx = append(dentBufferx, dentBuffer[:bytesCounter]...)
		if bytesCounter < len(dentBuffer) {
			break
		}
	}
	var cursor int
	dirents = make(map[Dirent]bool)
	for cursor < len(dentBufferx)-1 {
		dirent.Inode = *(*uint64)(unsafe.Pointer(&dentBufferx[cursor : cursor+8][0]))
		dirent.Offset = *(*uint64)(unsafe.Pointer(&dentBufferx[cursor+8 : cursor+16][0]))
		dirent.Reclen = *(*uint16)(unsafe.Pointer(&dentBufferx[cursor+16 : cursor+18][0]))
		dirent.Type = *(*byte)(unsafe.Pointer(&dentBufferx[cursor+18 : cursor+19][0]))
		nameBuffer = nameBuffer[0:0]
		for indexBuffer = cursor + 19; indexBuffer < cursor+int(dirent.Reclen); indexBuffer++ {
			if dentBufferx[indexBuffer] == byte(0) {
				break
			}
			nameBuffer = append(nameBuffer, dentBufferx[indexBuffer])
		}
		dirent.Name = string(nameBuffer[:len(nameBuffer)])
		cursor += int(dirent.Reclen)
		if dirent.Name == "." || dirent.Name == ".." {
			continue
		}
		dirents[dirent] = true
	}
	return dirents, nil
}

func BwToStr(bw float64) string {
	switch {
	case bw > math.Pow(1000, 7):
		return fmt.Sprintf("%.2fZ", bw/math.Pow(1000, 7))
	case bw > math.Pow(1000, 6):
		return fmt.Sprintf("%.2fE", bw/math.Pow(1000, 6))
	case bw > math.Pow(1000, 5):
		return fmt.Sprintf("%.2fP", bw/math.Pow(1000, 5))
	case bw > math.Pow(1000, 4):
		return fmt.Sprintf("%.2fT", bw/math.Pow(1000, 4))
	case bw > math.Pow(1000, 3):
		return fmt.Sprintf("%.2fG", bw/math.Pow(1000, 3))
	case bw > math.Pow(1000, 2):
		return fmt.Sprintf("%.2fM", bw/math.Pow(1000, 2))
	case bw > math.Pow(1000, 1):
		return fmt.Sprintf("%.2fK", bw/math.Pow(1000, 1))
	}
	return fmt.Sprintf("%g", bw)
}
