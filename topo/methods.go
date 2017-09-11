package topo

import (
	"fmt"
	"math"
	"time"

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
	t.Services = make(map[string]*ServiceInfo)
	return t
}

func (a *Addr) String() string {
	return a.Host + ":" + a.Port
}

func (a *AddrState) Update() {
	if a.fresh {
		a.Count++
	} else {
		a.fresh = true
		a.Count = 1
	}
}

func (addrs AddrSet) clean() {
	for addr, addrState = range addrs {
		if addrState.fresh {
			addrState.fresh = false
			addrs[addr] = addrState
		} else {
			delete(addrs, addr)
		}
	}
}

func (s *ServiceInfo) cleanAddrSets() {
	s.addrs = nil
	if s.upstream != nil {
		s.upstream.clean()
	}
	if len(s.upstream) == 0 {
		s.upstream = nil
	}
	if s.downstream != nil {
		s.downstream.clean()
	}
	if len(s.downstream) == 0 {
		s.downstream = nil
	}
	var str string
	if s.Addrs != nil {
		for str, addrState = range s.Addrs {
			if addrState.fresh {
				addrState.fresh = false
				s.Addrs[str] = addrState
			} else {
				delete(s.Addrs, str)
			}
		}
	}
	if len(s.Addrs) == 0 {
		s.Addrs = nil
	}
	if s.UpStream != nil {
		for str, addrState = range s.UpStream {
			if addrState.fresh {
				addrState.fresh = false
				s.UpStream[str] = addrState
			} else {
				delete(s.UpStream, str)
			}
		}
	}
	if len(s.UpStream) == 0 {
		s.UpStream = nil
	}
	if s.DownStream != nil {
		for str, addrState = range s.DownStream {
			if addrState.fresh {
				addrState.fresh = false
				s.DownStream[str] = addrState
			} else {
				delete(s.DownStream, str)
			}
		}
	}
	if len(s.DownStream) == 0 {
		s.DownStream = nil
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
		if serviceInfo, ok = t.Services[originProcInfo.Stat.Name]; !ok {
			serviceInfo = NewServiceInfo()
		}
		serviceInfo.ProcsStat[originProcInfo.Stat.Pid] = procStat
		t.Services[originProcInfo.Stat.Name] = serviceInfo
	}
	return nil
}

func (t *Topology) cleanAll() {
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
		serviceInfo.cleanAddrSets()
	}
}

func (t *Topology) doPortListen(port string) bool {
	for _, t.tempServiceInfo = range t.Services {
		if t.tempServiceInfo.DoListen {
			for t.tempAddr = range t.tempServiceInfo.addrs {
				if port == t.tempAddr.Port {
					return true
				}
			}
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
		localPortToName[si.LocalAddr.Port] = si.UserName
		if serviceInfo, ok = t.Services[si.UserName]; !ok {
			continue
		}
		if si.Status == psss.SsLISTEN {
			serviceInfo.DoListen = true
			if serviceInfo.addrs == nil {
				serviceInfo.addrs = make(map[Addr]AddrState)
			}
			addr.Host = si.LocalAddr.Host
			addr.Port = si.LocalAddr.Port
			addrState.Count = 1
			addrState.fresh = true
			serviceInfo.addrs[addr] = addrState
		} else {
			addr.Host = si.RemoteAddr.Host
			addr.Port = si.RemoteAddr.Port
			if t.doUserListen(si.UserName) {
				if t.doPortListen(si.LocalAddr.Port) {
					if serviceInfo.downstream == nil {
						serviceInfo.downstream = make(map[Addr]AddrState)
					}
					if addrState, ok = serviceInfo.downstream[addr]; !ok {
						addrState.Count = 1
						addrState.fresh = true
					} else {
						addrState.Update()
					}
					serviceInfo.downstream[addr] = addrState
				} else {
					if serviceInfo.upstream == nil {
						serviceInfo.upstream = make(map[Addr]AddrState)
					}
					if addrState, ok = serviceInfo.upstream[addr]; !ok {
						addrState.Count = 1
						addrState.fresh = true
					} else {
						addrState.Update()
					}
					serviceInfo.upstream[addr] = addrState
				}
				continue
			}
			if isHostLocal(si.RemoteAddr.Host) {
				if t.doPortListen(si.RemoteAddr.Port) {
					continue
				}
			}
			if serviceInfo.addrs == nil {
				serviceInfo.addrs = make(map[Addr]AddrState)
			}
			if addrState, ok = serviceInfo.addrs[addr]; !ok {
				addrState.Count = 1
				addrState.fresh = true
			} else {
				addrState.Update()
			}
			serviceInfo.addrs[addr] = addrState
		}
	}
	return nil
}

func (t *Topology) findUser() {
	var name string
	for _, serviceInfo = range t.Services {
		if serviceInfo.addrs == nil {
			goto upstream
		}
		if serviceInfo.Addrs == nil {
			serviceInfo.Addrs = make(map[string]AddrState)
		}
		for addr, addrState = range serviceInfo.addrs {
			serviceInfo.Addrs[addr.String()] = addrState
		}
	upstream:
		if serviceInfo.upstream == nil {
			goto downstream
		}
		if serviceInfo.UpStream == nil {
			serviceInfo.UpStream = make(map[string]AddrState)
		}
		for addr, addrState = range serviceInfo.upstream {
			if name, ok = localPortToName[addr.Port]; !ok {
				name = addr.String()
			}
			if t.tempAddrState, ok = serviceInfo.UpStream[name]; !ok {
				serviceInfo.UpStream[name] = AddrState{Count: 1, fresh: true}
				continue
			}
			if t.tempAddrState.fresh {
				t.tempAddrState.Count += addrState.Count
			} else {
				t.tempAddrState.Count = addrState.Count
				t.tempAddrState.fresh = true
			}
			serviceInfo.UpStream[name] = t.tempAddrState
		}
	downstream:
		if serviceInfo.downstream == nil {
			continue
		}
		if serviceInfo.DownStream == nil {
			serviceInfo.DownStream = make(map[string]AddrState)
		}
		for addr, addrState = range serviceInfo.downstream {
			if name, ok = localPortToName[addr.Port]; !ok {
				name = addr.String()
			}
			if t.tempAddrState, ok = serviceInfo.DownStream[name]; !ok {
				serviceInfo.DownStream[name] = AddrState{Count: 1, fresh: true}
				continue
			}
			if t.tempAddrState.fresh {
				t.tempAddrState.Count += addrState.Count
			} else {
				t.tempAddrState.Count = addrState.Count
				t.tempAddrState.fresh = true
			}
			serviceInfo.DownStream[name] = t.tempAddrState
		}
	}
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
	t.findUser()
	t.cleanAll()
	t.Time = time.Now().Unix()
	return nil
}

func (t *Topology) PrintAll() {
	for sname, si := range t.Services {
		fmt.Printf("Service: %s ", sname)
		if si.DoListen {
			fmt.Printf("(Listen)\n")
		} else {
			fmt.Printf("(NotListen)\n")
		}
		for pid, ps := range si.ProcsStat {
			fmt.Printf("\tPID:%d, StartTime:%d, State:%s, LoadAvg:%f, LoadInstant:%f, VmSize:%s, VmRSS:%s\n",
				pid, ps.StartTime, ps.State, ps.LoadAvg, ps.LoadInstant, psss.BwToStr(float64(ps.VmSize)), psss.BwToStr(float64(ps.VmRSS)),
			)
		}
		if si.DoListen {
			fmt.Println("\tListening Addr:")
			for addr, as := range si.Addrs {
				fmt.Printf("\t\t%s: %d\n", addr, as.Count)
			}
			if si.UpStream != nil {
				fmt.Println("\tUpstream Addr:")
				for addr, as := range si.UpStream {
					fmt.Printf("\t\t%s: %d\n", addr, as.Count)
				}
			}
			if si.DownStream != nil {
				fmt.Println("\tDownStream Addr:")
				for addr, as := range si.DownStream {
					fmt.Printf("\t\t%s: %d\n", addr, as.Count)
				}
			}
		} else {
			if len(si.Addrs) > 0 {
				fmt.Println("\tRemote Addr:")
				for addr, as := range si.Addrs {
					fmt.Printf("\t\t%s: %d\n", addr, as.Count)
				}
			}
		}
	}
}
