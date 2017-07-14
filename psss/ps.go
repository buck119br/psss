package psss

type Fd struct {
	Inode uint32
	Name  string
}

type ProcInfo struct {
	Stat ProcStat
	Fds  []Fd
}

func NewProcInfo() *ProcInfo {
	p := new(ProcInfo)
	p.Fds = make([]Fd, 0)
	return p
}

func (p *ProcInfo) Reset() {
	p.Fds = make([]Fd, 0)
}
