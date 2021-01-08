package sysinfo

import "github.com/shirou/gopsutil/disk"

func mountPoints() ([]string, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	paths := make([]string, len(partitions))

	for i, partition := range partitions {
		paths[i] = partition.Mountpoint
	}

	return paths, nil
}

func diskUsage(path string) ([]*disk.UsageStat, error) {
	paths, err := mountPoints()
	if err != nil {
		return nil, err
	}

	if path != "" {
		for _, p := range paths {
			if path == p {
				usage, err := disk.Usage(p)
				if err != nil {
					return nil, err
				}
				return []*disk.UsageStat{usage}, nil
			}
		}
	}

	usages := make([]*disk.UsageStat, len(paths))

	for i, p := range paths {
		usage, err := disk.Usage(p)
		if err != nil {
			return nil, err
		}
		usages[i] = usage
	}

	return usages, nil
}
