package psss

import (
	"fmt"
	"math"
	"unsafe"
)

type Dirent struct {
	Inode  uint64
	Offset uint64
	Reclen uint16
	Type   byte
	Name   string
}

func ParseDirent(buffer []byte) (dirents map[Dirent]bool, err error) {
	var cursor int
	dirents = make(map[Dirent]bool)
	for cursor < len(buffer)-1 {
		dirent.Inode = *(*uint64)(unsafe.Pointer(&buffer[cursor : cursor+8][0]))
		if dirent.Inode == 0 {
			return
		}
		dirent.Offset = *(*uint64)(unsafe.Pointer(&buffer[cursor+8 : cursor+16][0]))
		dirent.Reclen = *(*uint16)(unsafe.Pointer(&buffer[cursor+16 : cursor+18][0]))
		dirent.Type = *(*byte)(unsafe.Pointer(&buffer[cursor+18 : cursor+19][0]))
		dirent.Name = string(buffer[cursor+19 : cursor+int(dirent.Reclen)])
		cursor += int(dirent.Reclen)
		dirents[dirent] = true
	}
	return dirents, nil
}

func RefillBuffer(buffer []byte) {
	for i := range buffer {
		buffer[i] = 0
	}
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
