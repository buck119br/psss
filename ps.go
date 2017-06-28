package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ProcRoot = "/proc"
)

var GlobalProcInfo []*ProcInfo

type ProcInfo struct {
	Name string
	Pid  int
	Fd   []*FileInfo
}

func NewProcInfo() *ProcInfo {
	pi := new(ProcInfo)
	pi.Fd = make([]*FileInfo, 0, 0)
	return pi
}

func (p *ProcInfo) GetStatus() (err error) {
	var (
		line   []byte
		fields []string
	)
	fd, err := os.Open(ProcRoot + fmt.Sprintf("/%d/status", p.Pid))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	// Name
	if line, err = ReadLine(reader); err != nil {
		fmt.Println(err)
		return err
	}
	fields = strings.Fields(string(line))
	p.Name = fields[1]

	return nil
}

func GetProcInfo() {
	var tempPid int
	fd, err := os.Open(ProcRoot)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()
	names, err := fd.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	GlobalProcInfo = make([]*ProcInfo, 0, 0)
	for _, v := range names {
		if tempPid, err = strconv.Atoi(v); err != nil {
			continue
		}
		proc := NewProcInfo()
		proc.Pid = tempPid
		if proc.Fd, err = GetProcFiles(tempPid); err != nil {
			continue
		}
		if err = proc.GetStatus(); err != nil {
			continue
		}
		GlobalProcInfo = append(GlobalProcInfo, proc)
	}
	SetUpRelation()
}

func GetProcFiles(pid int) (files []*FileInfo, err error) {
	fdPath := ProcRoot + fmt.Sprintf("/%d/fd", pid)
	fd, err := os.Open(fdPath)
	if err != nil {
		fmt.Println(err)
		return files, err
	}
	defer fd.Close()
	names, err := fd.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return files, err
	}
	files = make([]*FileInfo, 0, 0)
	for _, v := range names {
		var file *FileInfo
		if file, err = GetFileStat(fdPath, v); err != nil {
			continue
		}
		files = append(files, file)
	}
	return files, nil
}

func SetUpRelation() {
	for _, proc := range GlobalProcInfo {
		for _, fd := range proc.Fd {
			for key, records := range GlobalRecords {
				if record, ok := records[uint32(fd.SysStat.Ino)]; ok {
					GlobalRecords[key][record.Inode].Procs[proc] = true
				}
			}
		}
	}
	findRecordUser()
}

func findRecordUser() {
	for _, records := range GlobalRecords {
		for _, record := range records {
			for proc := range record.Procs {
				for _, fd := range proc.Fd {
					if record.Inode == uint32(fd.SysStat.Ino) {
						record.UserName = proc.Name
						goto found
					}
				}
			}
		found:
		}
	}
}
