package main

import (
	"fmt"
	"github.com/curserio/control-agent/agent"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9990")
	if err != nil {
		panic(err)
	}

	var reply = &agent.DiskUsageReply{}

	err = client.Call("SystemInfoAPI.DiskUsage", "", reply)
	if err != nil {
		panic(err)
	}

	for _, disk := range reply.Disks {
		fmt.Printf("%+v\n", disk)
	}
}
