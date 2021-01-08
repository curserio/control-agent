package main

import (
	"github.com/curserio/control-agent/httplistener"
	"github.com/curserio/control-agent/internal"
	"github.com/curserio/control-agent/sysinfo"
)

func main() {

	//processes, err := sysinfo.ProcessesInfo()
	//if err != nil {
	//	panic(err)
	//}
	//
	//sysinfo.SortByMemoryUsage(processes)
	//
	//for _, process := range processes {
	//	memUsage := "nil"
	//	if process.MemoryInfo != nil {
	//		memUsage = misc.ByteCountBinary(int64(process.MemoryInfo.VMS))
	//	}
	//
	//	log.Printf("%v: %v percent, %v", process.Name, process.CPUPercent, memUsage)
	//}

	systemInfoAPI := sysinfo.New()
	agentListener := httplistener.New()

	core := internal.Init(systemInfoAPI, agentListener)
	err := core.Start()
	if err != nil {
		panic(err)
	}
}
