package psss

import (
	"bufio"
	"fmt"
	"math"
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
	CPUTimes        map[int]*CPUTime
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
	ss.CPUTimes = make(map[int]*CPUTime)
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

func (si *SystemInfo) GetStat() {
	var (
		line     string
		cpunr, n int
	)
	fd, err := os.Open(ProcRoot + "/stat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
		line = scanner.Text()
		switch {
		case strings.Contains(line, "cpu"):
			cputime := new(CPUTime)
			if strings.Contains(line, "cpu ") {
				cpunr = math.MaxInt16
				n, err = fmt.Sscanf(line, "cpu %d %d %d %d %d %d %d %d %d %d",
					&cputime.User, &cputime.Nice, &cputime.System, &cputime.Idle, &cputime.Iowait,
					&cputime.Irq, &cputime.Softirq, &cputime.Steal, &cputime.Guest, &cputime.GuestNice,
				)
				if n < 10 {
					fmt.Println("not enough param read")
					return
				}
			} else {
				n, err = fmt.Sscanf(line, "cpu%d %d %d %d %d %d %d %d %d %d %d",
					&cpunr,
					&cputime.User, &cputime.Nice, &cputime.System, &cputime.Idle, &cputime.Iowait,
					&cputime.Irq, &cputime.Softirq, &cputime.Steal, &cputime.Guest, &cputime.GuestNice,
				)
				if n < 11 {
					fmt.Println("not enough param read")
					return
				}
			}
			cputime.Total = cputime.User + cputime.Nice + cputime.System + cputime.Idle + cputime.Iowait +
				cputime.Irq + cputime.Softirq + cputime.Steal + cputime.Guest + cputime.GuestNice
			si.Stat.CPUTimes[cpunr] = cputime
		case strings.Contains(line, "page"):
			if n, err = fmt.Sscanf(line, "page %d %d", &si.Stat.PageIn, &si.Stat.PageOut); err != nil {
				fmt.Println(err)
				return
			}
			if n < 2 {
				fmt.Println("not enough param read")
				return
			}
		case strings.Contains(line, "swap"):
			if n, err = fmt.Sscanf(line, "swap %d %d", &si.Stat.SwapIn, &si.Stat.SwapOut); err != nil {
				fmt.Println(err)
				return
			}
			if n < 2 {
				fmt.Println("not enough param read")
				return
			}
		case strings.Contains(line, "intr"):
		case strings.Contains(line, "ctxt"):
			if n, err = fmt.Sscanf(line, "ctxt %d", &si.Stat.Ctxt); err != nil {
				fmt.Println(err)
				return
			}
			if n < 1 {
				fmt.Println("not enough param read")
				return
			}
		case strings.Contains(line, "btime"):
			if n, err = fmt.Sscanf(line, "btime %d", &si.Stat.Btime); err != nil {
				fmt.Println(err)
				return
			}
			if n < 1 {
				fmt.Println("not enough param read")
				return
			}
		case strings.Contains(line, "processes"):
			if n, err = fmt.Sscanf(line, "processes %d", &si.Stat.Processes); err != nil {
				fmt.Println(err)
				return
			}
			if n < 1 {
				fmt.Println("not enough param read")
				return
			}
		case strings.Contains(line, "procs_running"):
			if n, err = fmt.Sscanf(line, "procs_running %d", &si.Stat.ProcsRunning); err != nil {
				fmt.Println(err)
				return
			}
			if n < 1 {
				fmt.Println("not enough param read")
				return
			}
		case strings.Contains(line, "procs_blocked"):
			if n, err = fmt.Sscanf(line, "procs_blocked %d", &si.Stat.ProcsRunning); err != nil {
				fmt.Println(err)
				return
			}
			if n < 1 {
				fmt.Println("not enough param read")
				return
			}
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}