package internal

type DiskUsageReply struct {
	Disks []*DiskUsage
}

type DiskUsage struct {
	Path        string
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

type MemoryUsageReply struct {
	Total       uint64
	Available   uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

type ProcessesReq struct {
	Count uint
	Sort  ProcessesSort
}

type ProcessesSort int

const (
	WithoutSort ProcessesSort = iota
	ByName
	ByCPUUsage
	ByMemoryUsage
)

type ProcessesReply struct {
	Processes []*Process
}

type Process struct {
	Name          string
	CPUPercent    float64
	MemoryPercent float32
	MemoryUsed    uint64
}
