package topo

import (
	"math"

	"github.com/buck119br/psss/psss"
	"golang.org/x/sys/unix"
)

func NewServiceInfo() *ServiceInfo {
	s := new(ServiceInfo)
	s.ProcsStat = make(map[int]ProcStat)
	return s
}

func NewTopology() *Topology {
	t := new(Topology)
	t.Services = make(map[string]ServiceInfo)
	return t
}

func (a *AddrState) Update() {
	if a.fresh {
		a.Count++
	} else {
		a.fresh = true
		a.Count = 1
	}
}

func (t *Topology) GetProcInfo() (err error) {
	defer func() {
		clearReserve()
		SysInfoOld, SysInfoNew = SysInfoNew, SysInfoOld
	}()
	SysInfoNew.Reset()
	if err = SysInfoNew.GetStat(); err != nil {
		return err
	}

	go psss.ScanProcFS()
	for originProcInfo = range psss.ProcInfoChan {
		if originProcInfo.IsEnd {
			return nil
		}
		if serviceInfo, ok = t.Services[originProcInfo.Stat.Name]; !ok {
			serviceInfo = NewServiceInfo()
		}
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
			procStat.LoadInstant = math.Trunc(float64(originProcInfo.Stat.Utime+originProcInfo.Stat.Stime-procInfoReserve.Utime-procInfoReserve.Stime)/
				float64((SysInfoNew.Stat.CPUTime.Total-SysInfoOld.Stat.CPUTime.Total)/numCPU)*100000) / 100000
		}
		procInfoReserve.Utime = originProcInfo.Stat.Utime
		procInfoReserve.Stime = originProcInfo.Stat.Stime
		procInfoReserve.Fresh = true
		procsInfoReserve[originProcInfo.Stat.Name][originProcInfo.Stat.Pid] = procInfoReserve
		// assignment
		serviceInfo.ProcsStat[originProcInfo.Stat.Pid] = procStat
		t.Services[originProcInfo.Stat.Name] = serviceInfo
	}
	return nil
}

func (t *Topology) Clear() {
	var (
		name string
		pid  int
	)
	for name, serviceInfo = range t.Services {
		for pid, procStat = range serviceInfo.ProcsStat {
			if procStat.fresh {
				procStat.fresh = false
				serviceInfo.ProcsStat[pid] = procStat
			} else {
				delete(serviceInfo.ProcsStat, pid)
				if len(serviceInfo.ProcsStat) == 0 {
					delete(t.Services, name)
				}
			}
		}
	}
}

func (t *Topology) doPortListen(user, port string) bool {
	for t.tempAddr = range t.Services[user].Addrs {
		if port == t.tempAddr.Port {
			return true
		}
	}
	return false
}

func (t *Topology) doUserListen(user string) bool {
	return t.Services[user].DoListen
}

func (t *Topology) getSockInfo(af uint8, ssFilter uint32) (err error) {
	skfd, err := psss.SendInetDiagMsg(af, unix.IPPROTO_TCP, 0, ssFilter)
	if err != nil {
		return err
	}
	defer unix.Close(skfd)
	go psss.RecvInetDiagMsgAll(skfd)
	for si := range psss.SocketInfoChan {
		if si.IsEnd {
			return nil
		}
		// handle socket info
		localAddrToName[si.LocalAddr.String()] = si.Username
		serviceInfo, _ = t.Services[si.Username]
		if si.Status == psss.SsLISTEN {
			serviceInfo.DoListen = true
			if serviceInfo.Addrs == nil {
				serviceInfo.Addrs = make(map[Addr]AddrState)
			}
			addr.Host = si.LocalAddr.Host
			addr.Port = si.LocalAddr.Port
			addrState.Count = 1
			addrState.fresh = true
			serviceInfo.Addrs[addr] = addrState
		} else {
			addr.Host = si.Remote.Host
			addr.Port = si.Remote.Port
			if t.doUserListen(si.Username) {
				if t.doPortListen(si.Username, si.LocalAddr.Port) {
					if serviceInfo.DownStream == nil {
						serviceInfo.DownStream = make(map[Addr]AddrState)
					}
					addrState, _ = serviceInfo.DownStream[addr]
					addrState.Update()
					serviceInfo.DownStream[addr] = addrState
				} else {
					if serviceInfo.UpStream == nil {
						serviceInfo.UpStream = make(map[Addr]AddrState)
					}
					addrState, _ = serviceInfo.UpStream[addr]
					addrState.Update()
					serviceInfo.UpStream[addr] = addrState
				}
				t.Services[si.Username] = serviceInfo
				continue
			}
			if isHostLocal(si.Remote.Host) {
				if t.doPortListen(si.RemoteAddr.Port) {
					return
				}
			}
			serviceInfo.DoListen = false
			if serviceInfo.Addrs == nil {
				serviceInfo.Addrs = make(map[Addr]AddrState)
			}
			addrState, _ = serviceInfo.Addrs[addr]
			addrState.Update()
			serviceInfo.Addrs[addr] = addrState
		}
	}
	return nil
}

func (t *Topology) GetSockInfo() (err error) {
	if err = t.getSockInfo(unix.AF_INET, 1<<psss.SsLISTEN); err != nil {
		return err
	}
	if err = t.getSockInfo(unix.AF_INET6, 1<<psss.SsLISTEN); err != nil {
		return err
	}
	if err = t.getSockInfo(unix.AF_INET, 1<<psss.SsESTAB); err != nil {
		return err
	}
	if err = t.getSockInfo(unix.AF_INET6, 1<<psss.SsESTAB); err != nil {
		return err
	}
	psss.CleanGlobalProcFds()
	return nil
}
