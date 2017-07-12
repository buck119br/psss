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

type DirentHandler struct {
	InputSignalChan  chan bool
	OutputSignalChan chan bool
	Buffer           []byte
	Bufferx          []byte
	NameBuffer       []byte
	Dirent           *Dirent
	BytesCounter     int
	IndexBuffer      int
	Cursor           int
	Signal           bool
}

func NewDirentHandler() *DirentHandler {
	d := new(DirentHandler)
	d.InputSignalChan = make(chan bool)
	d.OutputSignalChan = make(chan bool)
	d.Buffer = make([]byte, os.Getpagesize())
	d.Bufferx = make([]byte, os.Getpagesize())
	d.NameBuffer = make([]byte, 256)
	d.Dirent = new(Dirent)
	return d
}

func (d *DirentHandler) ReadDirents(fd *os.File) {
	defer func() {
		<-d.InputSignalChan
		d.OutputSignalChan <- false
	}()
	d.Bufferx = d.Bufferx[0:0]
	var err error
	for {
		if d.BytesCounter, err = unix.Getdents(int(fd.Fd()), d.Buffer); err != nil {
			return
		}
		runtime.KeepAlive(fd)
		if d.BytesCounter == 0 {
			break
		}
		d.Bufferx = append(d.Bufferx, d.Buffer[:d.BytesCounter]...)
	}
	d.Cursor = 0
	for d.Cursor < len(d.Bufferx)-1 {
		<-d.InputSignalChan
		d.Dirent.Inode = *(*uint64)(unsafe.Pointer(&d.Bufferx[d.Cursor : d.Cursor+8][0]))
		d.Dirent.Offset = *(*uint64)(unsafe.Pointer(&d.Bufferx[d.Cursor+8 : d.Cursor+16][0]))
		d.Dirent.Reclen = *(*uint16)(unsafe.Pointer(&d.Bufferx[d.Cursor+16 : d.Cursor+18][0]))
		d.Dirent.Type = *(*byte)(unsafe.Pointer(&d.Bufferx[d.Cursor+18 : d.Cursor+19][0]))
		d.NameBuffer = d.NameBuffer[0:0]
		for d.IndexBuffer = d.Cursor + 19; d.IndexBuffer < d.Cursor+int(d.Dirent.Reclen); d.IndexBuffer++ {
			if d.Bufferx[d.IndexBuffer] == byte(0) {
				break
			}
			d.NameBuffer = append(d.NameBuffer, d.Bufferx[d.IndexBuffer])
		}
		d.Dirent.Name = string(d.NameBuffer[:len(d.NameBuffer)])
		d.Cursor += int(d.Dirent.Reclen)
		d.OutputSignalChan <- true
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
