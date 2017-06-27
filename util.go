package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"syscall"
)

func ReadLine(reader *bufio.Reader) (line []byte, err error) {
	var buffer []byte
	isPrefix := true
	line = make([]byte, 0, 0)
	for isPrefix {
		if buffer, isPrefix, err = reader.ReadLine(); err != nil {
			fmt.Println(err)
			return nil, err
		}
		line = append(line, buffer...)
	}
	return line, nil
}

type FileInfo struct {
	Path    string
	Name    string
	SysStat *syscall.Stat_t
}

func NewFileInfo() *FileInfo {
	fi := new(FileInfo)
	fi.SysStat = new(syscall.Stat_t)
	return fi
}

func GetFileStat(path string, name string) (fi *FileInfo, err error) {
	var (
		stat os.FileInfo
		ok   bool
	)
	if stat, err = os.Stat(path + "/" + name); err != nil {
		fmt.Println(err)
		return nil, err
	}
	fi = NewFileInfo()
	fi.Path = path
	fi.Name = name
	if fi.SysStat, ok = stat.Sys().(*syscall.Stat_t); !ok {
		fmt.Printf("FileInfo.Sys:[%v] assertion failure\n", stat)
		return nil, fmt.Errorf("FileInfo.Sys:[%v] assertion failure", stat)
	}
	return fi, nil
}

func BwToStr(bw float64) string {
	switch {
	case bw > math.Pow(1000, 7):
		return fmt.Sprintf("%.1gZ", bw/math.Pow(1000, 7))
	case bw > math.Pow(1000, 6):
		return fmt.Sprintf("%.1gE", bw/math.Pow(1000, 6))
	case bw > math.Pow(1000, 5):
		return fmt.Sprintf("%.1gP", bw/math.Pow(1000, 5))
	case bw > math.Pow(1000, 4):
		return fmt.Sprintf("%.1gT", bw/math.Pow(1000, 4))
	case bw > math.Pow(1000, 3):
		return fmt.Sprintf("%.1gG", bw/math.Pow(1000, 3))
	case bw > math.Pow(1000, 2):
		return fmt.Sprintf("%.1gM", bw/math.Pow(1000, 2))
	case bw > math.Pow(1000, 1):
		return fmt.Sprintf("%.1gK", bw/math.Pow(1000, 1))
	}
	return fmt.Sprintf("%g", bw)
}
