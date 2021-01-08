package sysinfo

import (
	"github.com/curserio/control-agent/agent"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfoAPI struct {
}

func New() *SystemInfoAPI {
	return &SystemInfoAPI{}
}

func (a *SystemInfoAPI) DiskUsage(path string, reply *agent.DiskUsageReply) error {
	u, err := diskUsage(path)
	if err != nil {
		return err
	}

	stat := make([]*agent.DiskUsage, len(u))

	for i, d := range u {
		stat[i] = &agent.DiskUsage{
			Path:        d.Path,
			Total:       d.Total,
			Free:        d.Free,
			Used:        d.Used,
			UsedPercent: d.UsedPercent,
		}
	}

	reply.Disks = stat

	return nil
}

func (a *SystemInfoAPI) MemoryUsage(nilArg interface{}, reply *agent.MemoryUsageReply) error {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	reply.Total = vm.Total
	reply.Available = vm.Available
	reply.Used = vm.Used
	reply.UsedPercent = vm.UsedPercent
	reply.Free = vm.Free

	return nil
}

func (a *SystemInfoAPI) SwapUsage(nilArg interface{}, reply *agent.MemoryUsageReply) error {
	vm, err := mem.SwapMemory()
	if err != nil {
		return err
	}

	reply.Total = vm.Total
	reply.Used = vm.Used
	reply.UsedPercent = vm.UsedPercent
	reply.Free = vm.Free

	return nil
}

func (a *SystemInfoAPI) Processes(req *agent.ProcessesReq, reply *agent.ProcessesReply) error {
	processes, err := processesInfo()
	if err != nil {
		return err
	}

	switch req.Sort {
	case agent.ByName:
		sortByName(processes)
	case agent.ByCPUUsage:
		sortByCPUUsage(processes)
	case agent.ByMemoryUsage:
		sortByMemoryUsage(processes)
	}

	if req.Count == 0 || req.Count > uint(len(processes)) {
		req.Count = uint(len(processes))
	}

	if uint(len(processes)) > req.Count {
		processes = processes[:int(req.Count)]
	}

	reply.Processes = processes

	return nil
}
