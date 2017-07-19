package topo

import (
	"math"

	"github.com/buck119br/psss/psss"
)

func NewServiceInfo() *ServiceInfo {
	s := new(ServiceInfo)
	s.ProcsStat = make(map[int]ProcStat)
	return s
}

func (t *Topology) GetProcInfo() error {
	defer func() {
		cleanReserve()
		SysInfoOld, SysInfoNew = SysInfoNew, SysInfoOld
	}()
	SysInfoNew.Reset()
	if err = SystemInfoNew.GetStat(); err != nil {
		return err
	}

	go psss.ScanProcFS()
	for originProcInfo = range psss.ProcInfoChan {
		if proc.IsEnd {
			return nil
		}
		if serviceInfo, ok = t[originProcInfo.Stat.Name]; !ok {
			serviceInfo = *(NewServiceInfo())
		}
		procStat, _ = serviceInfo.ProcsStat[originProcInfo.Stat.Pid]
		procStat.State = psss.ProcState[originProcInfo.Stat.State]
		procStat.StartTime = int64(SysInfoNew.Stat.Btime + originProcInfo.Stat.Starttime/psss.SC_CLK_TCK)
		procStat.LoadAvg = math.Trunc(float64(originProcInfo.Stat.Utime+originProcInfo.Stat.Stime)/float64(SysInfoNew.Stat.CPUTime.Total)*100000) / 100000
		procStat.LoadInstant = 0
		procStat.VmSize = originProcInfo.Stat.Vsize
		procStat.VmRSS = uint64(originProcInfo.Stat.Rss) * pageSize
		procStat.fresh = true
		// instant load
		if _, ok = procsInfoReserve[originProcInfo.Stat.Name]; !ok {
			procsInfoReserve[originProcInfo.Stat.Name] = make(map[int]*ProcInfoReserve)
		}
		if procInfoReserve, ok = procsInfoReserve[originProcInfo.Stat.Name][originProcInfo.Stat.Pid]; !ok {
			procInfoReserve = new(ProcInfoReserve)
		} else {
			newProcStat.LoadInstant = math.Trunc(float64(originProcInfo.Stat.Utime+originProcInfo.Stat.Stime-procInfoReserve.Utime-procInfoReserve.Stime)/
				float64((SystemInfoNew.Stat.CPUTime.Total-SystemInfoOld.Stat.CPUTime.Total)/NumCPU)*100000) / 100000
		}
		procInfoReserve.Utime = originProcInfo.Stat.Utime
		procInfoReserve.Stime = originProcInfo.Stat.Stime
		procInfoReserve.Fresh = true
		procsInfoReserve[originProcInfo.Stat.Name][originProcInfo.Stat.Pid] = procInfoReserve
		// assignment
		serviceInfo.ProcsStat[originProcInfo.Stat.Pid] = procStat
		t[originProcInfo.Stat.Name] = serviceInfo
	}
	return nil
}

func (t *Topology) Clear() {

}
