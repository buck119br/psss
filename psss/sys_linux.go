// +build linux

package psss

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MemoryInfo struct {
	MemTotal       uint64
	MemFree        uint64
	MemAvailable   uint64
	Buffers        uint64
	Cached         uint64
	SwapCached     uint64
	Active         uint64
	Inactive       uint64
	ActiveAnon     uint64
	InactiveAnon   uint64
	ActiveFile     uint64
	InactiveFile   uint64
	Unevictable    uint64
	Mlocked        uint64
	HighTotal      uint64
	HighFree       uint64
	LowTotal       uint64
	LowFree        uint64
	MmapCopy       uint64
	SwapTotal      uint64
	SwapFree       uint64
	Dirty          uint64
	Writeback      uint64
	AnonPages      uint64
	Mapped         uint64
	Shmem          uint64
	KReclaimable   uint64
	Slab           uint64
	SReclaimable   uint64
	SUnreclaim     uint64
	PageTables     uint64
	Quicklists     uint64
	NFSUnstable    uint64
	Bounce         uint64
	WritebackTmp   uint64
	CommitLimit    uint64
	CommittedAS    uint64
	VmallocUsed    uint64
	VmallocChunk   uint64
	Percpu         uint64
	AnonHugePages  uint64
	ShmemHugePages uint64
	ShmemPmdMapped uint64
	CmaTotal       uint64
	CmaFree        uint64
	HugePagesTotal uint64
	HugePagesFree  uint64
	HugePagesRsvd  uint64
	HugePagesSurp  uint64
	Hugepagesize   uint64
	DirectMap4k    uint64
	DirectMap2M    uint64
	DirectMap4M    uint64
	DirectMap1G    uint64
}

func (mi *MemoryInfo) Get() error {
	fd, err := os.Open(ProcRoot + "/meminfo")
	if err != nil {
		return err
	}
	defer fd.Close()

	var v uint64

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		fields := strings.Split(strings.Replace(scanner.Text(), "kB", "", -1), ":")
		if len(fields) != 2 {
			continue
		}
		if v, err = strconv.ParseUint(strings.TrimSpace(fields[1]), 10, 64); err != nil {
			continue
		}

		switch strings.TrimSpace(fields[0]) {
		case "MemTotal":
			mi.MemTotal = v
		case "MemFree":
			mi.MemFree = v
		case "MemAvailable":
			mi.MemAvailable = v
		case "Buffers":
			mi.Buffers = v
		case "Cached":
			mi.Cached = v
		case "SwapCached":
			mi.SwapCached = v
		case "Active":
			mi.Active = v
		case "Inactive":
			mi.Inactive = v
		case "ActiveAnon":
			mi.ActiveAnon = v
		case "InactiveAnon":
			mi.InactiveAnon = v
		case "ActiveFile":
			mi.ActiveFile = v
		case "InactiveFile":
			mi.InactiveFile = v
		case "Unevictable":
			mi.Unevictable = v
		case "Mlocked":
			mi.Mlocked = v
		case "HighTotal":
			mi.HighTotal = v
		case "HighFree":
			mi.HighFree = v
		case "LowTotal":
			mi.LowTotal = v
		case "LowFree":
			mi.LowFree = v
		case "MmapCopy":
			mi.MmapCopy = v
		case "SwapTotal":
			mi.SwapTotal = v
		case "SwapFree":
			mi.SwapFree = v
		case "Dirty":
			mi.Dirty = v
		case "Writeback":
			mi.Writeback = v
		case "AnonPages":
			mi.AnonPages = v
		case "Mapped":
			mi.Mapped = v
		case "Shmem":
			mi.Shmem = v
		case "KReclaimable":
			mi.KReclaimable = v
		case "Slab":
			mi.Slab = v
		case "SReclaimable":
			mi.SReclaimable = v
		case "SUnreclaim":
			mi.SUnreclaim = v
		case "PageTables":
			mi.PageTables = v
		case "Quicklists":
			mi.Quicklists = v
		case "NFS_Unstable":
			mi.NFSUnstable = v
		case "Bounce":
			mi.Bounce = v
		case "WritebackTmp":
			mi.WritebackTmp = v
		case "CommitLimit":
			mi.CommitLimit = v
		case "Committed_AS":
			mi.CommittedAS = v
		case "VmallocUsed":
			mi.VmallocUsed = v
		case "VmallocChunk":
			mi.VmallocChunk = v
		case "Percpu":
			mi.Percpu = v
		case "AnonHugePages":
			mi.AnonHugePages = v
		case "ShmemHugePages":
			mi.ShmemHugePages = v
		case "ShmemPmdMapped":
			mi.ShmemPmdMapped = v
		case "CmaTotal":
			mi.CmaTotal = v
		case "CmaFree":
			mi.CmaFree = v
		case "HugePages_Total":
			mi.HugePagesTotal = v
		case "HugePages_Free":
			mi.HugePagesFree = v
		case "HugePages_Rsvd":
			mi.HugePagesRsvd = v
		case "HugePages_Surp":
			mi.HugePagesSurp = v
		case "Hugepagesize":
			mi.Hugepagesize = v
		case "DirectMap4k":
			mi.DirectMap4k = v
		case "DirectMap2M":
			mi.DirectMap2M = v
		case "DirectMap4M":
			mi.DirectMap4M = v
		case "DirectMap1G":
			mi.DirectMap1G = v
		}
	}
	return nil
}

type CPUJiffies struct {
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
	CPUTotal        *CPUJiffies
	PageIn, PageOut uint64 // The number of pages the system paged in and the number that were paged out (from disk).
	SwapIn, SwapOut uint64 // The number of swap pages that have been brought in and out.
	Intr            uint64 // This line shows counts of interrupts serviced since boot time, for each of the possible system interrupts. The first column is the total of all interrupts serviced including unnumbered architecture specific interrupts; each subsequent column is the total for that particular numbered interrupt. Unnumbered interrupts are not shown, only summed into the total.
	Ctxt            uint64 // The number of context switches that the system underwent.
	Btime           uint64 // boot time, in seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC).
	Processes       uint64 // Number of forks since boot.
	ProcsRunning    uint64 // Number of processes in runnable state. (Linux 2.5.45 onward.)
	ProcsBlocked    uint64 // Number of processes blocked waiting for I/O to complete. (Linux 2.5.45 onward.)
}

func (si *SystemInfo) Reset() {
	si.Stat.CPUTotal.User = 0
	si.Stat.CPUTotal.Nice = 0
	si.Stat.CPUTotal.System = 0
	si.Stat.CPUTotal.Idle = 0
	si.Stat.CPUTotal.Iowait = 0
	si.Stat.CPUTotal.Irq = 0
	si.Stat.CPUTotal.Softirq = 0
	si.Stat.CPUTotal.Steal = 0
	si.Stat.CPUTotal.Guest = 0
	si.Stat.CPUTotal.GuestNice = 0
	si.Stat.CPUTotal.Total = 0
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

func (si *SystemInfo) Get() (err error) {
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
		line := scanner.Text()
		switch {
		case strings.Contains(line, "cpu "):
			bytesCounter, err = fmt.Sscanf(line, "cpu %d %d %d %d %d %d %d %d %d %d",
				&si.Stat.CPUTotal.User, &si.Stat.CPUTotal.Nice, &si.Stat.CPUTotal.System, &si.Stat.CPUTotal.Idle, &si.Stat.CPUTotal.Iowait,
				&si.Stat.CPUTotal.Irq, &si.Stat.CPUTotal.Softirq, &si.Stat.CPUTotal.Steal, &si.Stat.CPUTotal.Guest, &si.Stat.CPUTotal.GuestNice,
			)
			if bytesCounter < 10 {
				return fmt.Errorf("not enough param read")
			}
			si.Stat.CPUTotal.Total =
				si.Stat.CPUTotal.User + si.Stat.CPUTotal.Nice + si.Stat.CPUTotal.System + si.Stat.CPUTotal.Idle + si.Stat.CPUTotal.Iowait +
					si.Stat.CPUTotal.Irq + si.Stat.CPUTotal.Softirq + si.Stat.CPUTotal.Steal + si.Stat.CPUTotal.Guest + si.Stat.CPUTotal.GuestNice
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
	}
	return nil
}

type Uptime struct {
	Uptime float64
	Idle   float64
}

func (ut *Uptime) Get() error {
	raw, err := ioutil.ReadFile(ProcRoot + "/uptime")
	if err != nil {
		return err
	}

	n, err := fmt.Sscanf(string(raw), "%f %f", &ut.Uptime, &ut.Idle)
	if err != nil {
		return fmt.Errorf("scan error:[%v] with [%d] succeeded", err, n)
	}

	return nil
}
