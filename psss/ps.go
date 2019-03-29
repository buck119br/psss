package psss

type ProcInfo struct {
	Cmdline []string
	Stat    ProcStat
	IsEnd   bool
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	return p
}
