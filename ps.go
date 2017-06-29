package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ProcRoot = "/proc"
)

var GlobalProcInfo map[string]map[int]*ProcInfo

type ProcInfo struct {
	Name string
	Pid  int
	Fd   map[uint32]*FileInfo
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	p.Fd = make(map[uint32]*FileInfo, 0)
	return p
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

func (p *ProcInfo) GetFds() (err error) {
	fdPath := ProcRoot + fmt.Sprintf("/%d/fd", p.Pid)
	fd, err := os.Open(fdPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fd.Close()
	names, err := fd.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, v := range names {
		fi := NewFileInfo()
		if err = fi.GetStat(fdPath, v); err != nil {
			continue
		}
		p.Fd[uint32(fi.SysStat.Ino)] = fi
	}
	return nil
}

func GetProcInfo() {
	start := time.Now()
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
	var (
		tempInt int
		ok      bool
	)
	for _, v := range names {
		if tempInt, err = strconv.Atoi(v); err != nil {
			continue
		}
		proc := NewProcInfo()
		proc.Pid = tempInt
		if err = proc.GetFds(); err != nil {
			continue
		}
		if err = proc.GetStatus(); err != nil {
			continue
		}
		if _, ok = GlobalProcInfo[proc.Name]; !ok {
			GlobalProcInfo[proc.Name] = make(map[int]*ProcInfo)
		}
		GlobalProcInfo[proc.Name][proc.Pid] = proc
	}
	fmt.Println("GetProcInfo cost ", time.Since(start))
}

func SetUpRelation() {
	start := time.Now()
	var ok bool
	for key, records := range GlobalRecords {
		for ino := range records {
			for _, procMap := range GlobalProcInfo {
				for _, proc := range procMap {
					if _, ok = proc.Fd[ino]; ok {
						GlobalRecords[key][ino].UserName = proc.Name
						GlobalRecords[key][ino].Procs[proc] = true
					}
				}
			}
		}
	}
	fmt.Println("SetUpRelation cost ", time.Since(start))
}
