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

type ProcInfo struct {
	Name       string
	Pid        int
	Fd         []*FileInfo
	TCPSockets []*TCPRecord
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
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	// Name
	if line, err = ReadLine(reader); err != nil {
		return err
	}
	fields = strings.Fields(string(line))
	p.Name = fields[1]

	return nil
}

func GetProcInfo() (procs []*ProcInfo, err error) {
	var tempPid int
	fd, err := os.Open(ProcRoot)
	if err != nil {
		return procs, err
	}
	defer fd.Close()
	names, err := fd.Readdirnames(0)
	if err != nil {
		return procs, err
	}
	procs = make([]*ProcInfo, 0, 0)
	for _, v := range names {
		if tempPid, err = strconv.Atoi(v); err != nil {
			continue
		}
		proc := NewProcInfo()
		proc.Pid = tempPid
		if proc.Fd, err = GetProcFiles(tempPid); err != nil {
			continue
		}
		procs = append(procs, proc)
	}
	return procs, nil
}
