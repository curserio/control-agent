package sysinfo

import (
	"github.com/curserio/control-agent/internal"
	"github.com/shirou/gopsutil/process"
	"log"
	"sort"
)

func processesInfo() ([]*internal.Process, error) {
	processList, err := process.Processes()
	if err != nil {
		return nil, err
	}

	processes := make([]*internal.Process, len(processList))

	for i, p := range processList {
		processes[i] = &internal.Process{}

		processes[i].Name, err = p.Name()
		if err != nil {
			log.Printf("get Name for pid: %v error: %v", p.Pid, err)
		}

		processes[i].CPUPercent, err = p.CPUPercent()
		if err != nil {
			log.Printf("get CPUPercent for pid: %v error: %v", p.Pid, err)
		}
		processes[i].CPUPercent = processes[i].CPUPercent / 10

		processes[i].MemoryPercent, err = p.MemoryPercent()
		if err != nil {
			log.Printf("get MemoryPercent for pid: %v error: %v", p.Pid, err)
		}
		processes[i].MemoryPercent = processes[i].MemoryPercent / 10

		memInfo, err := p.MemoryInfo()
		if err != nil {
			log.Printf("get MemoryInfo for pid: %v error: %v", p.Pid, err)
		}

		processes[i].MemoryUsed = memInfo.RSS
	}

	return processes, nil
}

func sortByName(processes []*internal.Process) {
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].Name < processes[j].Name
	})
}

func sortByCPUUsage(processes []*internal.Process) {
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].CPUPercent < processes[j].CPUPercent
	})
}

func sortByMemoryUsage(processes []*internal.Process) {
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].MemoryPercent < processes[j].MemoryPercent
	})
}
