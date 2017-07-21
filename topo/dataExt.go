package topo

const zebraSchemaId64 int64 = 0x4df6151fb497 // 85719311692951

type Addr struct {
	Host string `zid:"0"`
	Port string `zid:"1"`
}

type AddrState struct {
	Count int `zid:"0"`
	fresh bool
}

type AddrSet map[Addr]AddrState

type ProcStat struct {
	StartTime   int64   `zid:"0"`
	State       string  `zid:"2"`
	LoadAvg     float64 `zid:"1"`
	LoadInstant float64 `zid:"3"`
	VmSize      uint64  `zid:"4"`
	VmRSS       uint64  `zid:"5"`
	fresh       bool
}

type ServiceInfo struct {
	ProcsStat  map[int]ProcStat     `zid:"0"`
	DoListen   bool                 `zid:"1"`
	Addrs      map[Addr]AddrState   `zid:"2" msgp:"omitempty"` // this field represents: listening addrs when DoListen is set, and remote addrs when DoListen is reset
	UpStream   map[string]AddrState `zid:"3" msgp:"omitempty"` // this field will not be nil only DoListen is set
	DownStream map[string]AddrState `zid:"4" msgp:"omitempty"` // this field will not be nil only DoListen is set
	upstream   AddrSet
	downstream AddrSet
}

type Topology struct {
	Services        map[string]*ServiceInfo `zid:"0"`
	Time            int64                   `zid:"1"`
	tempServiceInfo *ServiceInfo
	tempAddr        Addr
	tempAddrState   AddrState
}
