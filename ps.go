package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	ProcRoot = "/proc"
)

var GlobalProcInfo map[string]map[int]*ProcInfo

type ProcStat struct {
	Pid                 int
	Name                string
	State               string // single-char code for process state
	Ppid                int    // pid of parent process
	Pgrp                int    // process group id
	Session             int    // process group id
	TtyNr               int    // The controlling terminal of the process.  (The minor device number is contained in the combination of bits 31 to 20 and 7 to 0; the major device number is in bits 15 to 8.)
	Tpgid               int    // terminal process group id
	Flags               uint32 // kernel flags for the process
	Minflt              uint64 // number of minor page faults since process start
	Cminflt             uint64 // cumulative min_flt of process and child processes
	Majflt              uint64 // number of major page faults since process start
	Cmajflt             uint64 // cumulative maj_flt of process and child processes
	Utime               uint64 // user-mode CPU time accumulated by process
	Stime               uint64 // kernel-mode CPU time accumulated by process
	Cutime              uint64 // cumulative utime of process and reaped childre
	Cstime              uint64 // cumulative stime of process and reaped children
	Priority            int64  // kernel scheduling priority
	Nice                int64  // standard unix nice level of process
	NumThreads          int64  // number of threads, or 0 if no clue
	Itrealvalue         int64  // since kernel 2.6.17, this field is no longer maintained
	Starttime           uint64 // start time of process -- seconds since 1-1-70
	Vsize               uint64 // number of pages of virtual memory ...
	Rss                 int64  // resident set size from /proc/#/stat (pages)
	Rsslim              uint64 // resident set size limit
	Startcode           uint64 // address of beginning of code segment
	Endcode             uint64 // address of end of code segment
	Startstack          uint64 // address of the bottom of stack for the process
	Kstkesp             uint64 // kernel stack pointer
	Kstkeip             uint64 // kernel instruction pointer
	Signal              uint64 // mask of pending signals, per-task for readtask() but per-proc for readproc()
	Blocked             uint64 // mask of blocked signals
	Sigignore           uint64 // mask of ignored signals
	Sigcatch            uint64 // mask of caught  signals
	Wchan               uint64 // address of kernel wait channel proc is sleeping in
	Nswap               uint64 // Number of pages swapped (not maintained)
	Cnswap              uint64 // Cumulative nswap for child processes (not maintained).
	ExitSignal          int    // Signal to be sent to parent when we die
	Processor           int    // CPU number last executed on
	RtPriority          uint64 // real-time
	Policy              uint32 // Scheduling policy (see sched_setscheduler(2)). Decode using the SCHED_* constants in linux/sched.h.
	DelayacctBlkioTicks uint64 // Aggregated block I/O delays, measured in clock ticks (centiseconds).
	GuestTime           uint64 // Guest time of the process (time spent running a virtual CPU for a guest operating system), measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	CguestTime          int64  // Guest time of the process's children, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	StartData           uint64 // Address above which program initialized and uninitialized (BSS) data are placed.
	EndData             uint64 // Address below which program initialized and uninitialized (BSS) data are placed.
	StartBrk            uint64 // Address above which program heap can be expanded with brk(2).
	ArgStart            uint64 // Address above which program command-line arguments (argv) are placed.
	ArgEnd              uint64 // Address below program command-line arguments (argv) are placed.
	EnvStart            uint64 // Address above which program environment is placed.
	EnvEnd              uint64 // Address below which program environment is placed.
	ExitCode            int    // The thread's exit status in the form reported by waitpid(2).
}

type ProcInfo struct {
	Stat *ProcStat
	Fd   map[uint32]*FileInfo
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	p.Stat = new(ProcStat)
	p.Fd = make(map[uint32]*FileInfo, 0)
	return p
}

func (p *ProcInfo) GetStat() (err error) {
	fd, err := os.Open(ProcRoot + fmt.Sprintf("/%d/stat", p.Stat.Pid))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	statBuf, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return err
	}
	statBuf = statBuf[:len(statBuf)-1]
	fmt.Println(string(statBuf))
	n, err := fmt.Sscanf(string(statBuf),
		`%d (%s) `+
			`%s `+
			`%d %d %d %d %d `+
			`%d %d %d %d %d `+
			`%d %d %d %d `+
			`%d %d `+
			`%d %d %d `+
			`%d %d %d `+
			`%d %d %d `+
			`%d %d `+
			`%d %d %d %d `+
			`%d `+
			`%d %d `+
			`%d %d `+
			`%d %d `+
			`%d %d %d `+
			`%d %d `+
			`%d `+
			`%d %d `+
			`%d %d `+
			`%d`,
		&p.Stat.Pid, &p.Stat.Name,
		&p.Stat.State,
		&p.Stat.Ppid, &p.Stat.Pgrp, &p.Stat.Session, &p.Stat.TtyNr, &p.Stat.Tpgid,
		&p.Stat.Flags, &p.Stat.Minflt, &p.Stat.Cminflt, &p.Stat.Majflt, &p.Stat.Cmajflt,
		&p.Stat.Utime, &p.Stat.Stime, &p.Stat.Cutime, &p.Stat.Cstime,
		&p.Stat.Priority, &p.Stat.Nice,
		&p.Stat.NumThreads, &p.Stat.Itrealvalue, &p.Stat.Starttime,
		&p.Stat.Vsize, &p.Stat.Rss, &p.Stat.Rsslim,
		&p.Stat.Startcode, &p.Stat.Endcode, &p.Stat.Startstack,
		&p.Stat.Kstkesp, &p.Stat.Kstkeip,
		&p.Stat.Signal, &p.Stat.Blocked, &p.Stat.Sigignore, &p.Stat.Sigcatch, /* can't use */
		&p.Stat.Wchan,
		&p.Stat.Nswap, &p.Stat.Cnswap, /* nswap and cnswap dead for 2.4.xx and up */
		/* -- Linux 2.0.35 ends here -- */
		&p.Stat.ExitSignal, &p.Stat.Processor,
		/* -- Linux 2.2.8 to 2.5.17 end here -- */
		&p.Stat.RtPriority, &p.Stat.Policy, // /* both added to 2.5.18 */
		&p.Stat.DelayacctBlkioTicks, &p.Stat.GuestTime, &p.Stat.CguestTime,
		&p.Stat.StartData, &p.Stat.EndData,
		&p.Stat.StartBrk,
		&p.Stat.ArgStart, &p.Stat.ArgEnd,
		&p.Stat.EnvStart, &p.Stat.EnvEnd,
		&p.Stat.ExitCode,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if n < 52 {
		fmt.Println("not enough param read")
	}
	return nil
}

func (p *ProcInfo) GetFds() (err error) {
	fdPath := ProcRoot + fmt.Sprintf("/%d/fd", p.Stat.Pid)
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
		proc.Stat.Pid = tempInt
		if err = proc.GetFds(); err != nil {
			continue
		}
		if err = proc.GetStat(); err != nil {
			continue
		}
		if _, ok = GlobalProcInfo[proc.Stat.Name]; !ok {
			GlobalProcInfo[proc.Stat.Name] = make(map[int]*ProcInfo)
		}
		GlobalProcInfo[proc.Stat.Name][proc.Stat.Pid] = proc
	}
}

func SetUpRelation() {
	var ok bool
	for key, records := range GlobalRecords {
		for ino := range records {
			for _, procMap := range GlobalProcInfo {
				for _, proc := range procMap {
					if _, ok = proc.Fd[ino]; ok {
						GlobalRecords[key][ino].UserName = proc.Stat.Name
						GlobalRecords[key][ino].Procs[proc] = true
					}
				}
			}
		}
	}
}
