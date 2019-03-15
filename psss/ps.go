package psss



type ProcInfo struct {
	Stat  ProcStat
	IsEnd bool
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	return p
}
