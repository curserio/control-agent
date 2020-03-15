package sysinfo

import (
	"github.com/shirou/gopsutil/process"
)

type Process struct {
	Name string
	CPUPercent float64
	MemoryInfo *process.MemoryInfoStat
}

func ProcessesInfo() ([]*Process, error) {
	processList, err := process.Processes()
	if err != nil {
		return nil, err
	}

	processes := make([]*Process, len(processList))

	for i, p := range processList {
		processes[i] = &Process{}

		processes[i].Name, err = p.Name()
		if err != nil {
			return nil, err
		}
		processes[i].CPUPercent, err = p.CPUPercent()
		if err != nil {
			return nil, err
		}
		processes[i].MemoryInfo, err = p.MemoryInfo()
		if err != nil {
			return nil, err
		}
	}

	return processes, nil
}
