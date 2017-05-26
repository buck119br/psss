package main

import (
	"fmt"
	"os"
	"syscall"
)

type FileInfo struct {
	Name    string
	SysStat *syscall.Stat_t
}

func NewFileInfo() *FileInfo {
	fi := new(FileInfo)
	fi.SysStat = new(syscall.Stat_t)
	return fi
}

func GetFileStat(path string) (fi *FileInfo, err error) {
	var (
		stat os.FileInfo
		ok   bool
	)
	if stat, err = os.Stat(path); err != nil {
		return nil, err
	}
	fi = NewFileInfo()
	fi.Name = path
	if fi.SysStat, ok = stat.Sys().(*syscall.Stat_t); !ok {
		return nil, fmt.Errorf("FileInfo.Sys:[%v] assertion failure", stat)
	}
	return fi, nil
}

func GetProcFiles(pid int) (files []*FileInfo, err error) {
	var file *FileInfo
	fdPath := ProcRoot + fmt.Sprintf("/%d/fd", pid)
	fd, err := os.Open(fdPath)
	if err != nil {
		return files, err
	}
	defer fd.Close()
	names, err := fd.Readdirnames(0)
	if err != nil {
		return files, err
	}
	files = make([]*FileInfo, 0, 0)
	for _, v := range names {
		if file, err = GetFileStat(fdPath + "/" + v); err != nil {
			continue
		}
		files = append(files, file)
	}
	return files, nil
}
