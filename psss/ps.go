package psss

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

var ProcState = map[byte]string{
	'R': "Running",
	'S': "Sleeping", // sleeping in an interruptible wait
	'D': "Waiting",  // waiting in uninterruptible disk sleep
	'Z': "Zombie",
	'T': "Stopped",  // stopped (on a signal) or (before Linux 2.6.33) trace stopped
	't': "Tracing",  // stop (Linux 2.6.33 onward)
	'X': "Dead",     // (from Linux 2.6.0 onward)
	'x': "Dead",     // (Linux 2.6.33 to 3.13 only)
	'K': "Wakekill", // (Linux 2.6.33 to 3.13 only)
	'W': "Waking",   // (Linux 2.6.33 to 3.13 only)
	'P': "Parked",   // (Linux 3.9 to 3.13 only)
}

type ProcStat struct {
	Pid                 int
	Name                string // The filename of the executable, in parentheses. This is visible whether or not the executable is swapped out.
	State               byte   // single-char code for process state
	Ppid                int    // The PID of the parent of this process.
	Pgrp                int    // The process group ID of the process.
	Session             int    // The session ID of the process.
	TtyNr               int    // The controlling terminal of the process. (The minor device number is contained in the combination of bits 31 to 20 and 7 to 0; the major device number is in bits 15 to 8)
	Tpgid               int    // The ID of the foreground process group of the controlling terminal of the process.
	Flags               uint32 // The kernel flags word of the process. For bit meanings, see the PF_* defines in the Linux kernel source file include/linux/sched.h. Details depend on the kernel version.
	Minflt              uint64 // The number of minor faults the process has made which have not required loading a memory page from disk
	Cminflt             uint64 // The number of minor faults that the process's waited-for children have made
	Majflt              uint64 // The number of major faults the process has made which have required loading a memory page from disk
	Cmajflt             uint64 // The number of major faults that the process's waited-for children have made
	Utime               uint64 // Amount of time that this process has been scheduled in user mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)). This includes guest time, so that applications that are not aware of the guest time field do not lose that time from their calculations.
	Stime               uint64 // Amount of time that this process has been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	Cutime              uint64 // Amount of time that this process's waited-for children have been scheduled in user mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)). This includes guest time, cguest_time.
	Cstime              uint64 // Amount of time that this process's waited-for children have been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	Priority            int64  // For processes running a real-time scheduling policy, this is the negated scheduling priority, minus one; that is, a number in the range -2 to -100, corresponding to real-time priorities 1 to 99. For processes running under a non-real-time scheduling policy, this is the raw nice value as represented in the kernel. The kernel stores nice values as numbers in the range 0 (high) to 39 (low), corresponding to the user-visible nice range of -20 to 19.
	Nice                int64  // The nice value, a value in the range 19 (low priority) to -20 (high priority).
	NumThreads          int64  // Number of threads in this process (since Linux 2.6).
	Itrealvalue         int64  // Obsolete
	Starttime           uint64 // The time the process started after system boot. In kernels before Linux 2.6, this value was expressed in jiffies. Since Linux 2.6, the value is expressed in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	Vsize               uint64 // Virtual memory size in bytes.
	Rss                 int64  // Resident Set Size: number of pages the process has in real memory. This is just the pages which count toward text, data, or stack space. This does not include pages which have not been demand-loaded in, or which are swapped out.
	Rsslim              uint64 // Current soft limit in bytes on the rss of the process
	Startcode           uint64 // The address above which program text can run.
	Endcode             uint64 // The address below which program text can run.
	Startstack          uint64 // The address of the start (i.e., bottom) of the stack.
	Kstkesp             uint64 // The current value of ESP (stack pointer), as found in the kernel stack page for the process.
	Kstkeip             uint64 // The current EIP (instruction pointer).
	Signal              uint64 // Obsolete
	Blocked             uint64 // Obsolete
	Sigignore           uint64 // Obsolete
	Sigcatch            uint64 // Obsolete
	Wchan               uint64 // This is the "channel" in which the process is waiting. It is the address of a location in the kernel where the process is sleeping. The corresponding symbolic name can be found in /proc/[pid]/wchan.
	Nswap               uint64 // Obsolete
	Cnswap              uint64 // Obsolete
	ExitSignal          int    // Signal to be sent to parent when we die
	Processor           int    // CPU number last executed on
	RtPriority          uint64 // Real-time scheduling priority, a number in the range 1 to 99 for processes scheduled under a real-time policy, or 0, for non-real-time processes .
	Policy              uint32 // Scheduling policy. Decode using the SCHED_* constants in linux/sched.h.
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
	ExitCode            int    // The thread's exit status.
}

type ProcInfo struct {
	Stat *ProcStat
	Fd   map[uint32]string
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	p.Stat = new(ProcStat)
	p.Fd = make(map[uint32]string)
	return p
}

func (p *ProcInfo) Reset() {
	p.Fd = make(map[uint32]string)
}

func (p *ProcInfo) GetStat() (err error) {
	fd, err := os.Open(ProcRoot + fmt.Sprintf("/%d/stat", p.Stat.Pid))
	if err != nil {
		return err
	}
	defer fd.Close()
	FileContentBuffer.Reset()
	if _, err = FileContentBuffer.ReadFrom(fd); err != nil {
		return err
	}
	n, err := fmt.Sscanf(string(FileContentBuffer.Bytes()[:FileContentBuffer.Len()-1]),
		`%d %s %c %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d`,
		&p.Stat.Pid, &p.Stat.Name, &p.Stat.State,
		&p.Stat.Ppid, &p.Stat.Pgrp, &p.Stat.Session, &p.Stat.TtyNr, &p.Stat.Tpgid,
		&p.Stat.Flags, &p.Stat.Minflt, &p.Stat.Cminflt, &p.Stat.Majflt, &p.Stat.Cmajflt,
		&p.Stat.Utime, &p.Stat.Stime, &p.Stat.Cutime, &p.Stat.Cstime,
		&p.Stat.Priority, &p.Stat.Nice,
		&p.Stat.NumThreads, &p.Stat.Itrealvalue, &p.Stat.Starttime,
		&p.Stat.Vsize, &p.Stat.Rss, &p.Stat.Rsslim,
		&p.Stat.Startcode, &p.Stat.Endcode, &p.Stat.Startstack,
		&p.Stat.Kstkesp, &p.Stat.Kstkeip,
		&p.Stat.Signal, &p.Stat.Blocked, &p.Stat.Sigignore, &p.Stat.Sigcatch,
		&p.Stat.Wchan,
		&p.Stat.Nswap, &p.Stat.Cnswap,
		// since linux 2.1.22
		&p.Stat.ExitSignal,
		// since linux 2.2.8
		&p.Stat.Processor,
		// since linux 2.5.19
		&p.Stat.RtPriority, &p.Stat.Policy,
		// since linux 2.6.18
		&p.Stat.DelayacctBlkioTicks,
		// since linux 2.6.24
		&p.Stat.GuestTime, &p.Stat.CguestTime,
		// since linux 3.3
		&p.Stat.StartData, &p.Stat.EndData, &p.Stat.StartBrk,
		// since linux 3.5
		&p.Stat.ArgStart, &p.Stat.ArgEnd, &p.Stat.EnvStart, &p.Stat.EnvEnd, &p.Stat.ExitCode,
	)
	if err != nil {
		return err
	}
	if n < 52 {
		return fmt.Errorf("not enough param read")
	}
	p.Stat.Name = strings.TrimSuffix(strings.TrimPrefix(p.Stat.Name, "("), ")")
	return nil
}

func (p *ProcInfo) GetFds() (err error) {
	fdPath = ProcRoot + fmt.Sprintf("/%d/fd", p.Stat.Pid)
	fd, err := os.Open(fdPath)
	if err != nil {
		return err
	}
	defer fd.Close()
	go fdDirentHandler.ReadDirents(fd)
	fdDirentHandler.InputSignalChan <- true
	for fdDirentHandler.Signal = range fdDirentHandler.OutputSignalChan {
		if !fdDirentHandler.Signal {
			return
		}
		if err = syscall.Stat(fdPath+"/"+fdDirentHandler.Dirent.Name, fdStat_t); err != nil {
			goto next
		}
		p.Fd[uint32(fdStat_t.Ino)] = fdDirentHandler.Dirent.Name
	next:
		fdDirentHandler.InputSignalChan <- true
	}
	return nil
}

func GetProcInfo() {
	defer func() {
		<-ProcInfoInputChan
		ProcInfoOutputChan <- nil
	}()
	fd, err := os.Open(ProcRoot)
	if err != nil {
		return
	}
	defer fd.Close()

	var proc *ProcInfo
	go procDirentHandler.ReadDirents(fd)
	procDirentHandler.InputSignalChan <- true
	for procDirentHandler.Signal = range procDirentHandler.OutputSignalChan {
		if !procDirentHandler.Signal {
			return
		}
		if intBuffer, err = strconv.Atoi(procDirentHandler.Dirent.Name); err != nil {
			goto next
		}
		proc = <-ProcInfoInputChan
		proc.Stat.Pid = intBuffer
		if err = proc.GetStat(); err != nil {
			proc.Stat.Name = "NULL"
		}
		if err = proc.GetFds(); err != nil {
			proc.Stat.Name = "NULL"
		}
		ProcInfoOutputChan <- proc
	next:
		procDirentHandler.InputSignalChan <- true
	}
}
