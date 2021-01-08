package internal

type Core struct {
	listener Listener

	systemInfoAPI SystemInfoAPI
}

type Listener interface {
	Handle() error
	RegisterReceiver(name string, receiver interface{}) error
}

type SystemInfoAPI interface {
	DiskUsage(path string, reply *DiskUsageReply) error
	MemoryUsage(nilArg interface{}, reply *MemoryUsageReply) error
	SwapUsage(nilArg interface{}, reply *MemoryUsageReply) error
	Processes(req *ProcessesReq, reply *ProcessesReply) error
}
