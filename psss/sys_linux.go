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

var KVer = new(KernelVersion)

// definition comes from Linux kernel /fs/proc/meminfo.c
type MemoryInfo struct {
	MemTotal          uint64
	MemFree           uint64
	MemAvailable      uint64
	Buffers           uint64
	Cached            uint64
	SwapCached        uint64
	Active            uint64
	Inactive          uint64
	ActiveAnon        uint64
	InactiveAnon      uint64
	ActiveFile        uint64
	InactiveFile      uint64
	Unevictable       uint64
	Mlocked           uint64
	HighTotal         uint64
	HighFree          uint64
	LowTotal          uint64
	LowFree           uint64
	MmapCopy          uint64
	SwapTotal         uint64
	SwapFree          uint64
	Dirty             uint64
	Writeback         uint64
	AnonPages         uint64
	Mapped            uint64
	Shmem             uint64
	KReclaimable      uint64
	Slab              uint64
	SReclaimable      uint64
	SUnreclaim        uint64
	KernelStack       uint64
	PageTables        uint64
	Quicklists        uint64
	NFSUnstable       uint64
	Bounce            uint64
	WritebackTmp      uint64
	CommitLimit       uint64
	CommittedAS       uint64
	VmallocTotal      uint64
	VmallocUsed       uint64
	VmallocChunk      uint64
	Percpu            uint64
	HardwareCorrupted uint64
	AnonHugePages     uint64
	ShmemHugePages    uint64
	ShmemPmdMapped    uint64
	CmaTotal          uint64
	CmaFree           uint64
	HugePagesTotal    uint64
	HugePagesFree     uint64
	HugePagesRsvd     uint64
	HugePagesSurp     uint64
	Hugepagesize      uint64
	DirectMap4k       uint64
	DirectMap2M       uint64
	DirectMap4M       uint64
	DirectMap1G       uint64
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

		fields := strings.Split(SlimSpaceRegExp.ReplaceAllString(strings.Replace(scanner.Text(), "kB", "", -1), ""), ":")
		if len(fields) != 2 {
			return fmt.Errorf("line:[%v] two short", fields)
		}
		if v, err = strconv.ParseUint(fields[1], 10, 64); err != nil {
			return fmt.Errorf("parse field:[%s] error:[%v]", fields[1], err)
		}

		switch fields[0] {
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
		case "Active(anon)":
			mi.ActiveAnon = v
		case "Inactive(anon)":
			mi.InactiveAnon = v
		case "Active(file)":
			mi.ActiveFile = v
		case "Inactive(file)":
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
		case "KernelStack":
			mi.KernelStack = v
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
		case "VmallocTotal":
			mi.VmallocTotal = v
		case "VmallocUsed":
			mi.VmallocUsed = v
		case "VmallocChunk":
			mi.VmallocChunk = v
		case "Percpu":
			mi.Percpu = v
		case "HardwareCorrupted":
			mi.HardwareCorrupted = v
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
		default:
			return fmt.Errorf("invalid field:[%s]", fields[0])
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

// definition comes from Linux kernel /fs/proc/stat.c
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

func (ss *SystemStat) Get() (err error) {
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
			ss.CPUTotal = new(CPUJiffies)
			bytesCounter, err = fmt.Sscanf(line, "cpu %d %d %d %d %d %d %d %d %d %d",
				&ss.CPUTotal.User, &ss.CPUTotal.Nice, &ss.CPUTotal.System, &ss.CPUTotal.Idle, &ss.CPUTotal.Iowait,
				&ss.CPUTotal.Irq, &ss.CPUTotal.Softirq, &ss.CPUTotal.Steal, &ss.CPUTotal.Guest, &ss.CPUTotal.GuestNice,
			)
			if bytesCounter < 10 {
				return fmt.Errorf("not enough param read")
			}
			ss.CPUTotal.Total = ss.CPUTotal.User + ss.CPUTotal.Nice + ss.CPUTotal.System + ss.CPUTotal.Idle + ss.CPUTotal.Iowait +
				ss.CPUTotal.Irq + ss.CPUTotal.Softirq + ss.CPUTotal.Steal + ss.CPUTotal.Guest + ss.CPUTotal.GuestNice
		case strings.Contains(line, "page"):
			if bytesCounter, err = fmt.Sscanf(line, "page %d %d", &ss.PageIn, &ss.PageOut); err != nil {
				return err
			}
			if bytesCounter < 2 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "swap"):
			if bytesCounter, err = fmt.Sscanf(line, "swap %d %d", &ss.SwapIn, &ss.SwapOut); err != nil {
				return err
			}
			if bytesCounter < 2 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "intr"):
		case strings.Contains(line, "ctxt"):
			if bytesCounter, err = fmt.Sscanf(line, "ctxt %d", &ss.Ctxt); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "btime"):
			if bytesCounter, err = fmt.Sscanf(line, "btime %d", &ss.Btime); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "processes"):
			if bytesCounter, err = fmt.Sscanf(line, "processes %d", &ss.Processes); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "procs_running"):
			if bytesCounter, err = fmt.Sscanf(line, "procs_running %d", &ss.ProcsRunning); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		case strings.Contains(line, "procs_blocked"):
			if bytesCounter, err = fmt.Sscanf(line, "procs_blocked %d", &ss.ProcsRunning); err != nil {
				return err
			}
			if bytesCounter < 1 {
				return fmt.Errorf("not enough param read")
			}
		}
	}
	return nil
}

// definition comes from Linux kernel /fs/proc/uptime.c
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

// definition comes from Linux kernel /fs/proc/version.c
type KernelVersion struct {
	Origin      string
	UTSSysName  string
	UTSRelease  string
	CompileBy   string
	CompileHost string
	Compiler    string
	UTSVersion  string

	VersionNum int
}

func (kv *KernelVersion) Get() error {
	raw, err := ioutil.ReadFile(ProcRoot + "/version")
	if err != nil {
		return err
	}

	kv.Origin = strings.Replace(string(raw), "\n", "", -1)
	kv.UTSVersion = "#" + strings.Split(kv.Origin, "#")[1]

	fields := strings.Fields(strings.Split(kv.Origin, "#")[0])
	kv.UTSSysName = fields[0]
	kv.UTSRelease = fields[2]

	vfields := strings.Split(strings.Split(kv.UTSRelease, "-")[0], ".")
	major, err := strconv.Atoi(vfields[0])
	if err != nil {
		return err
	}
	minor, err := strconv.Atoi(vfields[1])
	if err != nil {
		return err
	}
	revision, err := strconv.Atoi(vfields[2])
	if err != nil {
		return err
	}
	kv.VersionNum = major*1000000 + minor*1000 + revision

	compile := strings.Split(fields[3], "@")
	kv.CompileBy = strings.TrimPrefix(compile[0], "(")
	kv.CompileHost = strings.TrimSuffix(compile[1], ")")

	kv.Compiler = strings.TrimSuffix(strings.TrimPrefix(strings.Join(fields[4:], " "), "("), ")")

	return nil
}
