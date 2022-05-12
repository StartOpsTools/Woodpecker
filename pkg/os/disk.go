package os

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

func GetDiskInfo() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return
	}

	for _, partition := range partitions {
		disk.Usage(partition.Mountpoint)

		usage, err := disk.Usage(partition.Mountpoint)
		if err == nil {
			fmt.Printf("%s \t %d \t %f%% \n", usage.Path, usage.Total, usage.UsedPercent)
		}
	}
}