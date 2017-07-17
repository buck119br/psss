package psss

type Fd struct {
	Name  string
	Fresh bool
}

type ProcInfo struct {
	Stat ProcStat
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	return p
}
