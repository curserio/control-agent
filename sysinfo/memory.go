package sysinfo

import (
	"github.com/shirou/gopsutil/mem"
)

func MemoryUsage() (*mem.VirtualMemoryStat, *mem.SwapMemoryStat, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, nil, err
	}

	sm, err := mem.SwapMemory()
	if err != nil {
		return nil, nil, err
	}

	return vm, sm, nil
}
