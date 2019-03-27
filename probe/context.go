package probe

import (
	"fmt"
	"io/ioutil"
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
	samplingCounter uint64

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
	pc.Uptime = new(psss.Uptime)
	return pc.Uptime.Get()
}

func (pc *ProbeContext) GetSystemStat() error {
	pc.SystemStat = new(psss.SystemStat)
	return pc.SystemStat.Get()
}

func (pc *ProbeContext) GetMemInfo() error {
	pc.MemoryInfo = new(psss.MemoryInfo)
	return pc.MemoryInfo.Get()
}

func (pc *ProbeContext) GetNetDevs() error {
	pc.NetDevs = psss.NewNetDevs()
	return pc.NetDevs.Get()
}

func (pc *ProbeContext) GetMountInfo() error {
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
		if _, ok = GConfig.Items.FileSystem.MountInfo.MountPointSet[mi.MountPoint]; !ok {
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
	pc.FileInfo = make([]*psss.FileInfo, 0, len(GConfig.Items.FileSystem.FileInfo.FilePath))
	var fi *psss.FileInfo
	var err error
	for _, v := range GConfig.Items.FileSystem.FileInfo.FilePath {
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
	defer tick.Stop()

	prev := NewProbeContext()
	err := prev.GetSystemStat()
	if err != nil {
		return err
	}
	if GConfig.Items.Process.Switch {
		prev.ProcInfo = psss.GetProcInfo(GConfig.Items.Process.ProcNameSet, false)
	}
	if GConfig.Items.IO.NIC.Switch {
		if err = prev.GetNetDevs(); err != nil {
			return err
		}
	}
	if GConfig.Items.FileSystem.MountInfo.Switch {
		if err = prev.GetMountInfo(); err != nil {
			return err
		}
	}

	select {
	case <-tick.C:
		if err = pc.GetSystemUptime(); err != nil {
			return err
		}
		if err = pc.GetSystemStat(); err != nil {
			return err
		}
		if err = pc.GetMemInfo(); err != nil {
			return err
		}
		if GConfig.Items.IO.NIC.Switch {
			if err = pc.GetNetDevs(); err != nil {
				return err
			}
		}
		if GConfig.Items.FileSystem.MountInfo.Switch {
			if err = pc.GetMountInfo(); err != nil {
				return err
			}
		}
		if GConfig.Items.Process.Switch {
			pc.ProcInfo = psss.GetProcInfo(GConfig.Items.Process.ProcNameSet, false)
		}

		// the following modules are costly
		if GConfig.Items.FileSystem.FileInfo.Switch {
			if err = pc.GetFileInfo(); err != nil {
				return err
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

	if GConfig.Items.IO.NIC.Switch {
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

	if GConfig.Items.FileSystem.MountInfo.Switch {
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
			emi.DiskStat.IOProgressing -= prevemi.DiskStat.IOProgressing
			emi.DiskStat.IOSpent -= prevemi.DiskStat.IOSpent
			emi.DiskStat.WeightedIOSpent -= prevemi.DiskStat.WeightedIOSpent
			emi.DiskStat.DiscardCompleted -= prevemi.DiskStat.DiscardCompleted
			emi.DiskStat.DiscardMerged -= prevemi.DiskStat.DiscardMerged
			emi.DiskStat.SectorDiscarded -= prevemi.DiskStat.SectorDiscarded
			emi.DiskStat.DiscardSpending -= prevemi.DiskStat.DiscardSpending
		}
	}

	if GConfig.Items.Process.Switch {
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

func (pc *ProbeContext) Fit(new *ProbeContext) {
	defer func() {
		pc.samplingCounter++
	}()

	if pc.samplingCounter == 0 {
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

	if GConfig.Items.IO.NIC.Switch {
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

	if GConfig.Items.FileSystem.MountInfo.Switch {
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

	if GConfig.Items.FileSystem.FileInfo.Switch {
		pc.FileInfo = new.FileInfo
	}

	if GConfig.Items.Process.Switch {
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
}
