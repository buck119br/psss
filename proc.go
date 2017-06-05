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

func GetProcFiles(pid int) (files []*FileInfo, err error) {
	var file *FileInfo
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
		if file, err = GetFileStat(fdPath + "/" + v); err != nil {
			continue
		}
		files = append(files, file)
	}
	return files, nil
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
}

func findRecordUser(records map[uint64]*GenericRecord) {
	for _, record := range records {
		for proc := range record.Procs {
			for _, fd := range proc.Fd {
				if (Sstate[record.Status] == "LISTEN" || Sstate[record.Status] == "ESTAB") && record.Inode == fd.SysStat.Ino {
					record.User = proc.Name
					goto found
				}
			}
		}
	found:
	}
}

func SetUpRelation() {
	var (
		record *GenericRecord
		ok     bool
	)
	for _, proc := range GlobalProcInfo {
		for _, fd := range proc.Fd {
			if record, ok = GlobalTCPv4Records[fd.SysStat.Ino]; ok {
				record.Procs[proc] = true
				GlobalTCPv4Records[fd.SysStat.Ino] = record
			}
			if record, ok = GlobalTCPv6Records[fd.SysStat.Ino]; ok {
				record.Procs[proc] = true
				GlobalTCPv6Records[fd.SysStat.Ino] = record
			}
			if record, ok = GlobalUDPv4Records[fd.SysStat.Ino]; ok {
				record.Procs[proc] = true
				GlobalUDPv4Records[fd.SysStat.Ino] = record
			}
			if record, ok = GlobalUDPv6Records[fd.SysStat.Ino]; ok {
				record.Procs[proc] = true
				GlobalUDPv6Records[fd.SysStat.Ino] = record
			}
		}
	}
	findRecordUser(GlobalTCPv4Records)
	findRecordUser(GlobalTCPv6Records)
	findRecordUser(GlobalUDPv4Records)
	findRecordUser(GlobalUDPv6Records)
}
