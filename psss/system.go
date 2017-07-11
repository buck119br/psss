package psss

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CPUTime struct {
	User      uint64 // Time spent in user mode.
	Nice      uint64 // Time spent in user mode with low priority (nice).
	System    uint64 // Time spent in system mode.
	Idle      uint64 // Time spent in the idle task. This value should be USER_HZ times the second entry in the /proc/uptime pseudo-file.
	Iowait    uint64 // Time waiting for I/O to complete.
	Irq       uint64 // Time servicing interrupts.
	Softirq   uint64 // Time servicing softirqs.
	Steal     uint64 // Stolen time, which is the time spent in other operating systems when running in a virtualized environment
	Guest     uint64 // Time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.
	GuestNice uint64 // Time spent running a niced guest (virtual CPU for guest operating ystems under the control of the Linux kernel).
	Total     uint64 // not specified in /proc/stat
}

type SystemStat struct {
	CPUTime         CPUTime
	PageIn, PageOut uint64 // The number of pages the system paged in and the number that were paged out (from disk).
	SwapIn, SwapOut uint64 // The number of swap pages that have been brought in and out.
	Intr            uint64 // This line shows counts of interrupts serviced since boot time, for each of the possible system interrupts. The first column is the total of all interrupts serviced including unnumbered architecture specific interrupts; each subsequent column is the total for that particular numbered interrupt. Unnumbered interrupts are not shown, only summed into the total.
	Ctxt            uint64 // The number of context switches that the system underwent.
	Btime           uint64 // boot time, in seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC).
	Processes       uint64 // Number of forks since boot.
	ProcsRunning    uint64 // Number of processes in runnable state. (Linux 2.5.45 onward.)
	ProcsBlocked    uint64 // Number of processes blocked waiting for I/O to complete. (Linux 2.5.45 onward.)
}

func NewSystemStat() *SystemStat {
	ss := new(SystemStat)
	return ss
}

type SystemInfo struct {
	Stat *SystemStat
}

func NewSystemInfo() *SystemInfo {
	si := new(SystemInfo)
	si.Stat = NewSystemStat()
	return si
}

func (si *SystemInfo) Reset() {
	si.Stat.CPUTime.User = 0
	si.Stat.CPUTime.Nice = 0
	si.Stat.CPUTime.System = 0
	si.Stat.CPUTime.Idle = 0
	si.Stat.CPUTime.Iowait = 0
	si.Stat.CPUTime.Irq = 0
	si.Stat.CPUTime.Softirq = 0
	si.Stat.CPUTime.Steal = 0
	si.Stat.CPUTime.Guest = 0
	si.Stat.CPUTime.GuestNice = 0
	si.Stat.CPUTime.Total = 0
	si.Stat.PageIn = 0
	si.Stat.PageOut = 0
	si.Stat.SwapIn = 0
	si.Stat.SwapOut = 0
	si.Stat.Intr = 0
	si.Stat.Ctxt = 0
	si.Stat.Btime = 0
	si.Stat.Processes = 0
	si.Stat.ProcsRunning = 0
	si.Stat.ProcsBlocked = 0
}

func (si *SystemInfo) GetStat() (err error) {
	var line string
	fd, err := os.Open(ProcRoot + "/stat")
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		line = scanner.Text()
		switch {
		case strings.Contains(line, "cpu "):
			bytesCounter, err = fmt.Sscanf(line, "cpu %d %d %d %d %d %d %d %d %d %d",
				&si.Stat.CPUTime.User, &si.Stat.CPUTime.Nice, &si.Stat.CPUTime.System, &si.Stat.CPUTime.Idle, &si.Stat.CPUTime.Iowait,
				&si.Stat.CPUTime.Irq, &si.Stat.CPUTime.Softirq, &si.Stat.CPUTime.Steal, &si.Stat.CPUTime.Guest, &si.Stat.CPUTime.GuestNice,
			)
			if bytesCounter < 10 {
				return fmt.Errorf("not enough param read")
			}
			si.Stat.CPUTime.Total = si.Stat.CPUTime.User + si.Stat.CPUTime.Nice + si.Stat.CPUTime.System + si.Stat.CPUTime.Idle + si.Stat.CPUTime.Iowait +
				si.Stat.CPUTime.Irq + si.Stat.CPUTime.Softirq + si.Stat.CPUTime.Steal + si.Stat.CPUTime.Guest + si.Stat.CPUTime.GuestNice
		case strings.Contains(line, "page"):
			if bytesCounter, err = fmt.Sscanf(line, "page %d %d", &si.Stat.PageIn, &si.Stat.PageOut); err != nil {
				return err
			}
			if bytesCounter < 2 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "swap"):
			if bytesCounter, err = fmt.Sscanf(line, "swap %d %d", &si.Stat.SwapIn, &si.Stat.SwapOut); err != nil {
				return err
			}
			if bytesCounter < 2 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "intr"):
		case strings.Contains(line, "ctxt"):
			if bytesCounter, err = fmt.Sscanf(line, "ctxt %d", &si.Stat.Ctxt); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "btime"):
			if bytesCounter, err = fmt.Sscanf(line, "btime %d", &si.Stat.Btime); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "processes"):
			if bytesCounter, err = fmt.Sscanf(line, "processes %d", &si.Stat.Processes); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "procs_running"):
			if bytesCounter, err = fmt.Sscanf(line, "procs_running %d", &si.Stat.ProcsRunning); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "procs_blocked"):
			if bytesCounter, err = fmt.Sscanf(line, "procs_blocked %d", &si.Stat.ProcsRunning); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
