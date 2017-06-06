package main

import (
	"bufio"
	"fmt"
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
