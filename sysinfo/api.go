package sysinfo

import (
	"github.com/curserio/control-agent/internal"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfoAPI struct {
}

func New() *SystemInfoAPI {
	return &SystemInfoAPI{}
}

func (a *SystemInfoAPI) DiskUsage(path string, reply *internal.DiskUsageReply) error {
	u, err := diskUsage(path)
	if err != nil {
		return err
	}

	stat := make([]*internal.DiskUsage, len(u))

	for i, d := range u {
		stat[i] = &internal.DiskUsage{
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

func (a *SystemInfoAPI) MemoryUsage(nilArg interface{}, reply *internal.MemoryUsageReply) error {
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

func (a *SystemInfoAPI) SwapUsage(nilArg interface{}, reply *internal.MemoryUsageReply) error {
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

func (a *SystemInfoAPI) Processes(req *internal.ProcessesReq, reply *internal.ProcessesReply) error {
	processes, err := processesInfo()
	if err != nil {
		return err
	}

	switch req.Sort {
	case internal.ByName:
		sortByName(processes)
	case internal.ByCPUUsage:
		sortByCPUUsage(processes)
	case internal.ByMemoryUsage:
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
