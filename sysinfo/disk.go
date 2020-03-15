package sysinfo

import (
	"github.com/shirou/gopsutil/disk"
)

func diskMountpoints() ([]string, error) {
	d, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(d))

	for i, partition := range d {
		names[i] = partition.Mountpoint
	}

	return names, nil
}

func DisksUsage() ([]*disk.UsageStat, error) {
	names, err := diskMountpoints()
	if err != nil {
		return nil, err
	}

	usages := make([]*disk.UsageStat, len(names))

	for i, name := range names {
		usage, err := disk.Usage(name)
		if err != nil {
			return nil, err
		}
		usages[i] = usage
	}

	return usages, nil
}
