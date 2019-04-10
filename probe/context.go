package probe

import (
	"fmt"
	"io/ioutil"
	"runtime/debug"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/buck119br/psss/psss"
)

type extMountInfo struct {
	MountInfo    *psss.MountInfo
	Fstat        *syscall.Statfs_t
	DiskStat     *psss.DiskStat
	HWSectorSize uint64
}

type ProbeContext struct {
	SamplingCounter uint64

	Uptime     *psss.Uptime
	SystemStat *psss.SystemStat
	MemoryInfo *psss.MemoryInfo
	NetDevs    psss.NetDevs
	MountInfo  map[string]*extMountInfo
	FileInfo   []*psss.FileInfo
	ProcInfo   map[string]map[int]*psss.ProcInfo
}

func NewProbeContext() *ProbeContext {
	pc := new(ProbeContext)
	return pc
}

func (pc *ProbeContext) GetSystemUptime() error {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	pc.Uptime = new(psss.Uptime)
	return pc.Uptime.Get()
}

func (pc *ProbeContext) GetSystemStat() error {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	pc.SystemStat = new(psss.SystemStat)
	return pc.SystemStat.Get()
}

func (pc *ProbeContext) GetMemoryInfo() error {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	pc.MemoryInfo = new(psss.MemoryInfo)
	return pc.MemoryInfo.Get()
}

func (pc *ProbeContext) GetNetDevs() error {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	pc.NetDevs = psss.NewNetDevs()
	return pc.NetDevs.Get()
}

func (pc *ProbeContext) GetMountInfo() error {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	mis := psss.NewMountInfos()
	err := mis.Get()
	if err != nil {
		return err
	}
	dss := psss.NewDiskStats()
	if err = dss.Get(); err != nil {
		return err
	}

	pc.MountInfo = make(map[string]*extMountInfo)
	var ok bool
	var raw []byte
	for _, mi := range mis {
		if _, ok = GConfig.FileSystem.MountInfo.MountPointSet[mi.MountPoint]; !ok {
			continue
		}
		emi := new(extMountInfo)
		emi.MountInfo = mi
		emi.Fstat = new(syscall.Statfs_t)
		if err = syscall.Statfs(emi.MountInfo.MountPoint, emi.Fstat); err != nil {
			logger.Errorf("get file system stat error:[%v]", err)
			continue
		}
		for _, ds := range dss {
			if ds.MajorNumber != emi.MountInfo.DiskMajorNum || ds.MinorNumber != (emi.MountInfo.DiskMinorNum/16)*16 {
				continue
			}
			emi.DiskStat = ds
			if raw, err = ioutil.ReadFile(fmt.Sprintf("/sys/block/%s/queue/hw_sector_size", emi.DiskStat.Name)); err != nil {
				logger.Errorf("get sector size error:[%v]", err)
				continue
			}
			if emi.HWSectorSize, err = strconv.ParseUint(strings.Replace(string(raw), "\n", "", -1), 10, 64); err != nil {
				logger.Errorf("parse sector size error:[%v]", err)
				continue
			}
		}
		pc.MountInfo[emi.MountInfo.MountPoint] = emi
	}
	return nil
}

func (pc *ProbeContext) GetFileInfo() error {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	pc.FileInfo = make([]*psss.FileInfo, 0, len(GConfig.FileSystem.FileInfo.FilePath))
	var fi *psss.FileInfo
	var err error
	for _, v := range GConfig.FileSystem.FileInfo.FilePath {
		if fi, err = psss.GetFileInfo(v); err != nil {
			logger.WithField("path", v).Errorf("error:[%v]", err)
			continue
		}
		pc.FileInfo = append(pc.FileInfo, fi)
	}
	return nil
}

func (pc *ProbeContext) Sample() error {
	tick := time.NewTicker(time.Second)
	defer func() {
		tick.Stop()
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	prev := NewProbeContext()
	err := prev.GetSystemStat()
	if err != nil {
		logger.Errorf("get system stat error:[%v]", err)
	}
	if GConfig.Process.Switch {
		prev.ProcInfo = psss.GetProcInfo(GConfig.Process.ProcNameSet, false)
	}
	if GConfig.IO.NIC.Switch {
		if err = prev.GetNetDevs(); err != nil {
			logger.Errorf("get net devs error:[%v]", err)
		}
	}
	if GConfig.FileSystem.MountInfo.Switch {
		if err = prev.GetMountInfo(); err != nil {
			logger.Errorf("get mount info error:[%v]", err)
		}
	}

	select {
	case <-tick.C:
		if err = pc.GetSystemUptime(); err != nil {
			logger.Errorf("get system uptime error:[%v]", err)
		}
		if err = pc.GetSystemStat(); err != nil {
			logger.Errorf("get system stat error:[%v]", err)
		}
		if err = pc.GetMemoryInfo(); err != nil {
			logger.Errorf("get memory info error:[%v]", err)
		}
		if GConfig.IO.NIC.Switch {
			if err = pc.GetNetDevs(); err != nil {
				logger.Errorf("get net devs error:[%v]", err)
			}
		}
		if GConfig.FileSystem.MountInfo.Switch {
			if err = pc.GetMountInfo(); err != nil {
				logger.Errorf("get mount info error:[%v]", err)
			}
		}
		if GConfig.Process.Switch {
			pc.ProcInfo = psss.GetProcInfo(GConfig.Process.ProcNameSet, false)
		}

		// the following modules are costly
		if GConfig.FileSystem.FileInfo.Switch {
			if err = pc.GetFileInfo(); err != nil {
				logger.Errorf("get file info error:[%v]", err)
			}
		}
	}

	pc.SystemStat.CPUTotal.User -= prev.SystemStat.CPUTotal.User
	pc.SystemStat.CPUTotal.Nice -= prev.SystemStat.CPUTotal.Nice
	pc.SystemStat.CPUTotal.System -= prev.SystemStat.CPUTotal.System
	pc.SystemStat.CPUTotal.Idle -= prev.SystemStat.CPUTotal.Idle
	pc.SystemStat.CPUTotal.Iowait -= prev.SystemStat.CPUTotal.Iowait
	pc.SystemStat.CPUTotal.Irq -= prev.SystemStat.CPUTotal.Irq
	pc.SystemStat.CPUTotal.Softirq -= prev.SystemStat.CPUTotal.Softirq
	pc.SystemStat.CPUTotal.Steal -= prev.SystemStat.CPUTotal.Steal
	pc.SystemStat.CPUTotal.Guest -= prev.SystemStat.CPUTotal.Guest
	pc.SystemStat.CPUTotal.GuestNice -= prev.SystemStat.CPUTotal.GuestNice
	pc.SystemStat.CPUTotal.Total -= prev.SystemStat.CPUTotal.Total

	if GConfig.IO.NIC.Switch {
		for _, nic := range pc.NetDevs {
			prevnic, ok := prev.NetDevs[nic.Interface]
			if !ok {
				continue
			}
			nic.ReceiveBytes -= prevnic.ReceiveBytes
			nic.ReceivePackets -= prevnic.ReceivePackets
			nic.ReceiveErrs -= prevnic.ReceiveErrs
			nic.ReceiveDrop -= prevnic.ReceiveDrop
			nic.ReceiveFifo -= prevnic.ReceiveFifo
			nic.ReceiveFrame -= prevnic.ReceiveFrame
			nic.ReceiveCompressed -= prevnic.ReceiveCompressed
			nic.ReceiveMulticast -= prevnic.ReceiveMulticast
			nic.TransmitBytes -= prevnic.TransmitBytes
			nic.TransmitPackets -= prevnic.TransmitPackets
			nic.TransmitErrs -= prevnic.TransmitErrs
			nic.TransmitDrop -= prevnic.TransmitDrop
			nic.TransmitFifo -= prevnic.TransmitFifo
			nic.TransmitColls -= prevnic.TransmitColls
			nic.TransmitCarrier -= prevnic.TransmitCarrier
			nic.TransmitCompressed -= prevnic.TransmitCompressed
		}
	}

	if GConfig.FileSystem.MountInfo.Switch {
		for _, emi := range pc.MountInfo {
			if emi.DiskStat == nil {
				continue
			}
			prevemi, ok := prev.MountInfo[emi.MountInfo.MountPoint]
			if !ok {
				continue
			}
			emi.DiskStat.ReadCompleted -= prevemi.DiskStat.ReadCompleted
			emi.DiskStat.ReadMerged -= prevemi.DiskStat.ReadMerged
			emi.DiskStat.SectorsRead -= prevemi.DiskStat.SectorsRead
			emi.DiskStat.ReadingSpent -= prevemi.DiskStat.ReadingSpent
			emi.DiskStat.WriteCompleted -= prevemi.DiskStat.WriteCompleted
			emi.DiskStat.WriteMerged -= prevemi.DiskStat.WriteMerged
			emi.DiskStat.SectorsWritten -= prevemi.DiskStat.SectorsWritten
			emi.DiskStat.WritingSpent -= prevemi.DiskStat.WritingSpent
			emi.DiskStat.IOSpent -= prevemi.DiskStat.IOSpent
			emi.DiskStat.WeightedIOSpent -= prevemi.DiskStat.WeightedIOSpent
			emi.DiskStat.DiscardCompleted -= prevemi.DiskStat.DiscardCompleted
			emi.DiskStat.DiscardMerged -= prevemi.DiskStat.DiscardMerged
			emi.DiskStat.SectorDiscarded -= prevemi.DiskStat.SectorDiscarded
			emi.DiskStat.DiscardSpending -= prevemi.DiskStat.DiscardSpending
		}
	}

	if GConfig.Process.Switch {
		for name, procs := range pc.ProcInfo {
			prevprocs, ok := prev.ProcInfo[name]
			if !ok {
				continue
			}
			for pid, pi := range procs {
				prevpi, ok := prevprocs[pid]
				if !ok {
					continue
				}
				pi.Stat.Utime -= prevpi.Stat.Utime
				pi.Stat.Stime -= prevpi.Stat.Stime
				pi.Stat.Cutime -= prevpi.Stat.Cutime
				pi.Stat.Cstime -= prevpi.Stat.Cstime
			}
		}
	}

	return nil
}

func (pc *ProbeContext) FitSystemStat(new *ProbeContext) {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	pc.SystemStat.CPUTotal.User += new.SystemStat.CPUTotal.User
	pc.SystemStat.CPUTotal.Nice += new.SystemStat.CPUTotal.Nice
	pc.SystemStat.CPUTotal.System += new.SystemStat.CPUTotal.System
	pc.SystemStat.CPUTotal.Idle += new.SystemStat.CPUTotal.Idle
	pc.SystemStat.CPUTotal.Iowait += new.SystemStat.CPUTotal.Iowait
	pc.SystemStat.CPUTotal.Irq += new.SystemStat.CPUTotal.Irq
	pc.SystemStat.CPUTotal.Softirq += new.SystemStat.CPUTotal.Softirq
	pc.SystemStat.CPUTotal.Steal += new.SystemStat.CPUTotal.Steal
	pc.SystemStat.CPUTotal.Guest += new.SystemStat.CPUTotal.Guest
	pc.SystemStat.CPUTotal.GuestNice += new.SystemStat.CPUTotal.GuestNice
	pc.SystemStat.CPUTotal.Total += new.SystemStat.CPUTotal.Total
	pc.SystemStat.Btime = new.SystemStat.Btime
}

func (pc *ProbeContext) FitMemoryInfo(new *ProbeContext) {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	pc.MemoryInfo.MemTotal += new.MemoryInfo.MemTotal
	pc.MemoryInfo.MemFree += new.MemoryInfo.MemFree
	pc.MemoryInfo.MemAvailable += new.MemoryInfo.MemAvailable
	pc.MemoryInfo.Buffers += new.MemoryInfo.Buffers
	pc.MemoryInfo.Cached += new.MemoryInfo.Cached
	pc.MemoryInfo.SwapCached += new.MemoryInfo.SwapCached
	pc.MemoryInfo.Active += new.MemoryInfo.Active
	pc.MemoryInfo.Inactive += new.MemoryInfo.Inactive
	pc.MemoryInfo.ActiveAnon += new.MemoryInfo.ActiveAnon
	pc.MemoryInfo.InactiveAnon += new.MemoryInfo.InactiveAnon
	pc.MemoryInfo.ActiveFile += new.MemoryInfo.ActiveFile
	pc.MemoryInfo.InactiveFile += new.MemoryInfo.InactiveFile
	pc.MemoryInfo.Unevictable += new.MemoryInfo.Unevictable
	pc.MemoryInfo.Mlocked += new.MemoryInfo.Mlocked
	pc.MemoryInfo.HighTotal += new.MemoryInfo.HighTotal
	pc.MemoryInfo.HighFree += new.MemoryInfo.HighFree
	pc.MemoryInfo.LowTotal += new.MemoryInfo.LowTotal
	pc.MemoryInfo.LowFree += new.MemoryInfo.LowFree
	pc.MemoryInfo.MmapCopy += new.MemoryInfo.MmapCopy
	pc.MemoryInfo.SwapTotal += new.MemoryInfo.SwapTotal
	pc.MemoryInfo.SwapFree += new.MemoryInfo.SwapFree
	pc.MemoryInfo.Dirty += new.MemoryInfo.Dirty
	pc.MemoryInfo.Writeback += new.MemoryInfo.Writeback
	pc.MemoryInfo.AnonPages += new.MemoryInfo.AnonPages
	pc.MemoryInfo.Mapped += new.MemoryInfo.Mapped
	pc.MemoryInfo.Shmem += new.MemoryInfo.Shmem
	pc.MemoryInfo.KReclaimable += new.MemoryInfo.KReclaimable
	pc.MemoryInfo.Slab += new.MemoryInfo.Slab
	pc.MemoryInfo.SReclaimable += new.MemoryInfo.SReclaimable
	pc.MemoryInfo.SUnreclaim += new.MemoryInfo.SUnreclaim
	pc.MemoryInfo.KernelStack += new.MemoryInfo.KernelStack
	pc.MemoryInfo.PageTables += new.MemoryInfo.PageTables
	pc.MemoryInfo.Quicklists += new.MemoryInfo.Quicklists
	pc.MemoryInfo.NFSUnstable += new.MemoryInfo.NFSUnstable
	pc.MemoryInfo.Bounce += new.MemoryInfo.Bounce
	pc.MemoryInfo.WritebackTmp += new.MemoryInfo.WritebackTmp
	pc.MemoryInfo.CommitLimit += new.MemoryInfo.CommitLimit
	pc.MemoryInfo.CommittedAS += new.MemoryInfo.CommittedAS
	pc.MemoryInfo.VmallocTotal += new.MemoryInfo.VmallocTotal
	pc.MemoryInfo.VmallocUsed += new.MemoryInfo.VmallocUsed
	pc.MemoryInfo.VmallocChunk += new.MemoryInfo.VmallocChunk
	pc.MemoryInfo.Percpu += new.MemoryInfo.Percpu
	pc.MemoryInfo.HardwareCorrupted += new.MemoryInfo.HardwareCorrupted
	pc.MemoryInfo.AnonHugePages += new.MemoryInfo.AnonHugePages
	pc.MemoryInfo.ShmemHugePages += new.MemoryInfo.ShmemHugePages
	pc.MemoryInfo.ShmemPmdMapped += new.MemoryInfo.ShmemPmdMapped
	pc.MemoryInfo.CmaTotal += new.MemoryInfo.CmaTotal
	pc.MemoryInfo.CmaFree += new.MemoryInfo.CmaFree
	pc.MemoryInfo.HugePagesTotal += new.MemoryInfo.HugePagesTotal
	pc.MemoryInfo.HugePagesFree += new.MemoryInfo.HugePagesFree
	pc.MemoryInfo.HugePagesRsvd += new.MemoryInfo.HugePagesRsvd
	pc.MemoryInfo.HugePagesSurp += new.MemoryInfo.HugePagesSurp
	pc.MemoryInfo.Hugepagesize += new.MemoryInfo.Hugepagesize
	pc.MemoryInfo.DirectMap4k += new.MemoryInfo.DirectMap4k
	pc.MemoryInfo.DirectMap2M += new.MemoryInfo.DirectMap2M
	pc.MemoryInfo.DirectMap4M += new.MemoryInfo.DirectMap4M
	pc.MemoryInfo.DirectMap1G += new.MemoryInfo.DirectMap1G
}

func (pc *ProbeContext) FitNetDevs(new *ProbeContext) {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	for _, newnic := range new.NetDevs {
		nic, ok := pc.NetDevs[newnic.Interface]
		if !ok {
			continue
		}
		nic.ReceiveBytes += newnic.ReceiveBytes
		nic.ReceivePackets += newnic.ReceivePackets
		nic.ReceiveErrs += newnic.ReceiveErrs
		nic.ReceiveDrop += newnic.ReceiveDrop
		nic.ReceiveFifo += newnic.ReceiveFifo
		nic.ReceiveFrame += newnic.ReceiveFrame
		nic.ReceiveCompressed += newnic.ReceiveCompressed
		nic.ReceiveMulticast += newnic.ReceiveMulticast
		nic.TransmitBytes += newnic.TransmitBytes
		nic.TransmitPackets += newnic.TransmitPackets
		nic.TransmitErrs += newnic.TransmitErrs
		nic.TransmitDrop += newnic.TransmitDrop
		nic.TransmitFifo += newnic.TransmitFifo
		nic.TransmitColls += newnic.TransmitColls
		nic.TransmitCarrier += newnic.TransmitCarrier
		nic.TransmitCompressed += newnic.TransmitCompressed
	}
}

func (pc *ProbeContext) FitDiskStat(new *ProbeContext) {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	for _, newemi := range new.MountInfo {
		if newemi.DiskStat == nil {
			continue
		}
		emi, ok := pc.MountInfo[newemi.MountInfo.MountPoint]
		if !ok {
			continue
		}
		emi.DiskStat.ReadCompleted += newemi.DiskStat.ReadCompleted
		emi.DiskStat.ReadMerged += newemi.DiskStat.ReadMerged
		emi.DiskStat.SectorsRead += newemi.DiskStat.SectorsRead
		emi.DiskStat.ReadingSpent += newemi.DiskStat.ReadingSpent
		emi.DiskStat.WriteCompleted += newemi.DiskStat.WriteCompleted
		emi.DiskStat.WriteMerged += newemi.DiskStat.WriteMerged
		emi.DiskStat.SectorsWritten += newemi.DiskStat.SectorsWritten
		emi.DiskStat.WritingSpent += newemi.DiskStat.WritingSpent
		emi.DiskStat.IOProgressing += newemi.DiskStat.IOProgressing
		emi.DiskStat.IOSpent += newemi.DiskStat.IOSpent
		emi.DiskStat.WeightedIOSpent += newemi.DiskStat.WeightedIOSpent
		emi.DiskStat.DiscardCompleted += newemi.DiskStat.DiscardCompleted
		emi.DiskStat.DiscardMerged += newemi.DiskStat.DiscardMerged
		emi.DiskStat.SectorDiscarded += newemi.DiskStat.SectorDiscarded
		emi.DiskStat.DiscardSpending += newemi.DiskStat.DiscardSpending
	}
}

func (pc *ProbeContext) FitProcInfo(new *ProbeContext) {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()

	for name, newprocs := range new.ProcInfo {
		procs, ok := pc.ProcInfo[name]
		if !ok {
			continue
		}
		for pid, newpi := range newprocs {
			pi, ok := procs[pid]
			if !ok {
				continue
			}
			pi.Stat.Utime += newpi.Stat.Utime
			pi.Stat.Stime += newpi.Stat.Stime
			pi.Stat.Cutime += newpi.Stat.Cutime
			pi.Stat.Cstime += newpi.Stat.Cstime
		}
	}
}

func (pc *ProbeContext) Fit(new *ProbeContext) {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic with error:[%v]", rcvErr)
			debug.PrintStack()
		}
		pc.SamplingCounter++
	}()

	if pc.SamplingCounter == 0 {
		pc.Uptime = new.Uptime
		pc.SystemStat = new.SystemStat
		pc.MemoryInfo = new.MemoryInfo
		pc.NetDevs = new.NetDevs
		pc.MountInfo = new.MountInfo
		pc.FileInfo = new.FileInfo
		pc.ProcInfo = new.ProcInfo
		return
	}

	pc.Uptime = new.Uptime
	pc.FitSystemStat(new)
	pc.FitMemoryInfo(new)

	if GConfig.IO.NIC.Switch {
		pc.FitNetDevs(new)
	}

	if GConfig.FileSystem.MountInfo.Switch {
		pc.FitDiskStat(new)
	}

	if GConfig.FileSystem.FileInfo.Switch {
		pc.FileInfo = new.FileInfo
	}

	if GConfig.Process.Switch {
		pc.FitProcInfo(new)
	}
}

func (pc *ProbeContext) Average() {
	if pc.SamplingCounter < 1 {
		return
	}

	pc.SystemStat.CPUTotal.User /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Nice /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.System /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Idle /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Iowait /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Irq /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Softirq /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Steal /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Guest /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.GuestNice /= pc.SamplingCounter
	pc.SystemStat.CPUTotal.Total /= pc.SamplingCounter

	pc.MemoryInfo.MemTotal /= pc.SamplingCounter
	pc.MemoryInfo.MemFree /= pc.SamplingCounter
	pc.MemoryInfo.MemAvailable /= pc.SamplingCounter
	pc.MemoryInfo.Buffers /= pc.SamplingCounter
	pc.MemoryInfo.Cached /= pc.SamplingCounter
	pc.MemoryInfo.SwapCached /= pc.SamplingCounter
	pc.MemoryInfo.Active /= pc.SamplingCounter
	pc.MemoryInfo.Inactive /= pc.SamplingCounter
	pc.MemoryInfo.ActiveAnon /= pc.SamplingCounter
	pc.MemoryInfo.InactiveAnon /= pc.SamplingCounter
	pc.MemoryInfo.ActiveFile /= pc.SamplingCounter
	pc.MemoryInfo.InactiveFile /= pc.SamplingCounter
	pc.MemoryInfo.Unevictable /= pc.SamplingCounter
	pc.MemoryInfo.Mlocked /= pc.SamplingCounter
	pc.MemoryInfo.HighTotal /= pc.SamplingCounter
	pc.MemoryInfo.HighFree /= pc.SamplingCounter
	pc.MemoryInfo.LowTotal /= pc.SamplingCounter
	pc.MemoryInfo.LowFree /= pc.SamplingCounter
	pc.MemoryInfo.MmapCopy /= pc.SamplingCounter
	pc.MemoryInfo.SwapTotal /= pc.SamplingCounter
	pc.MemoryInfo.SwapFree /= pc.SamplingCounter
	pc.MemoryInfo.Dirty /= pc.SamplingCounter
	pc.MemoryInfo.Writeback /= pc.SamplingCounter
	pc.MemoryInfo.AnonPages /= pc.SamplingCounter
	pc.MemoryInfo.Mapped /= pc.SamplingCounter
	pc.MemoryInfo.Shmem /= pc.SamplingCounter
	pc.MemoryInfo.KReclaimable /= pc.SamplingCounter
	pc.MemoryInfo.Slab /= pc.SamplingCounter
	pc.MemoryInfo.SReclaimable /= pc.SamplingCounter
	pc.MemoryInfo.SUnreclaim /= pc.SamplingCounter
	pc.MemoryInfo.KernelStack /= pc.SamplingCounter
	pc.MemoryInfo.PageTables /= pc.SamplingCounter
	pc.MemoryInfo.Quicklists /= pc.SamplingCounter
	pc.MemoryInfo.NFSUnstable /= pc.SamplingCounter
	pc.MemoryInfo.Bounce /= pc.SamplingCounter
	pc.MemoryInfo.WritebackTmp /= pc.SamplingCounter
	pc.MemoryInfo.CommitLimit /= pc.SamplingCounter
	pc.MemoryInfo.CommittedAS /= pc.SamplingCounter
	pc.MemoryInfo.VmallocTotal /= pc.SamplingCounter
	pc.MemoryInfo.VmallocUsed /= pc.SamplingCounter
	pc.MemoryInfo.VmallocChunk /= pc.SamplingCounter
	pc.MemoryInfo.Percpu /= pc.SamplingCounter
	pc.MemoryInfo.HardwareCorrupted /= pc.SamplingCounter
	pc.MemoryInfo.AnonHugePages /= pc.SamplingCounter
	pc.MemoryInfo.ShmemHugePages /= pc.SamplingCounter
	pc.MemoryInfo.ShmemPmdMapped /= pc.SamplingCounter
	pc.MemoryInfo.CmaTotal /= pc.SamplingCounter
	pc.MemoryInfo.CmaFree /= pc.SamplingCounter
	pc.MemoryInfo.HugePagesTotal /= pc.SamplingCounter
	pc.MemoryInfo.HugePagesFree /= pc.SamplingCounter
	pc.MemoryInfo.HugePagesRsvd /= pc.SamplingCounter
	pc.MemoryInfo.HugePagesSurp /= pc.SamplingCounter
	pc.MemoryInfo.Hugepagesize /= pc.SamplingCounter
	pc.MemoryInfo.DirectMap4k /= pc.SamplingCounter
	pc.MemoryInfo.DirectMap2M /= pc.SamplingCounter
	pc.MemoryInfo.DirectMap4M /= pc.SamplingCounter
	pc.MemoryInfo.DirectMap1G /= pc.SamplingCounter

	if GConfig.IO.NIC.Switch {
		for _, nic := range pc.NetDevs {
			nic.ReceiveBytes /= pc.SamplingCounter
			nic.ReceivePackets /= pc.SamplingCounter
			nic.ReceiveErrs /= pc.SamplingCounter
			nic.ReceiveDrop /= pc.SamplingCounter
			nic.ReceiveFifo /= pc.SamplingCounter
			nic.ReceiveFrame /= pc.SamplingCounter
			nic.ReceiveCompressed /= pc.SamplingCounter
			nic.ReceiveMulticast /= pc.SamplingCounter
			nic.TransmitBytes /= pc.SamplingCounter
			nic.TransmitPackets /= pc.SamplingCounter
			nic.TransmitErrs /= pc.SamplingCounter
			nic.TransmitDrop /= pc.SamplingCounter
			nic.TransmitFifo /= pc.SamplingCounter
			nic.TransmitColls /= pc.SamplingCounter
			nic.TransmitCarrier /= pc.SamplingCounter
			nic.TransmitCompressed /= pc.SamplingCounter
		}
	}

	if GConfig.FileSystem.MountInfo.Switch {
		for _, emi := range pc.MountInfo {
			if emi.DiskStat == nil {
				continue
			}
			emi.DiskStat.ReadCompleted /= pc.SamplingCounter
			emi.DiskStat.ReadMerged /= pc.SamplingCounter
			emi.DiskStat.SectorsRead /= pc.SamplingCounter
			emi.DiskStat.ReadingSpent /= pc.SamplingCounter
			emi.DiskStat.WriteCompleted /= pc.SamplingCounter
			emi.DiskStat.WriteMerged /= pc.SamplingCounter
			emi.DiskStat.SectorsWritten /= pc.SamplingCounter
			emi.DiskStat.WritingSpent /= pc.SamplingCounter
			emi.DiskStat.IOProgressing /= pc.SamplingCounter
			emi.DiskStat.IOSpent /= pc.SamplingCounter
			emi.DiskStat.WeightedIOSpent /= pc.SamplingCounter
			emi.DiskStat.DiscardCompleted /= pc.SamplingCounter
			emi.DiskStat.DiscardMerged /= pc.SamplingCounter
			emi.DiskStat.SectorDiscarded /= pc.SamplingCounter
			emi.DiskStat.DiscardSpending /= pc.SamplingCounter
		}
	}

	if GConfig.Process.Switch {
		for _, procs := range pc.ProcInfo {
			for _, pi := range procs {
				pi.Stat.Utime /= pc.SamplingCounter
				pi.Stat.Stime /= pc.SamplingCounter
				pi.Stat.Cutime /= pc.SamplingCounter
				pi.Stat.Cstime /= pc.SamplingCounter
			}
		}
	}
}
