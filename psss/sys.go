package psss

func NewSystemStat() *SystemStat {
	ss := new(SystemStat)
	return ss
}

type SystemInfo struct {
	Stat SystemStat
}

func NewSystemInfo() *SystemInfo {
	si := new(SystemInfo)
	return si
}
