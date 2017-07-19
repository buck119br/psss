package psss

type Fd struct {
	Name  string
	Fresh bool
}

type ProcInfo struct {
	Stat  ProcStat
	IsEnd bool
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	return p
}
