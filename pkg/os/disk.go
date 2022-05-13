package os

import (
	"fmt"
	"github.com/StartOpsTools/Woodpecker/pkg/output"
	"github.com/shirou/gopsutil/disk"
)

func OutputDiskUsageInfo() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return
	}

	_, _ = output.Title.Println("DISK INFO:")
	fmt.Printf("%-30s \t %-20s \t %-20s \t %-15s \t %s \n", "Path", "Total(MB)", "Free(MB)", "UsedPercent(%)", "InodesUsedPercent(%)")
	for _, partition := range partitions {

		usage, err := disk.Usage(partition.Mountpoint)
		if err == nil {
			if usage.UsedPercent >= 80 {
				_, _ = output.Warning.Printf("%-30s \t %-20d \t %-20d \t %-15f \t %f \n", usage.Path, usage.Total/1000/1024, usage.Free/1000/1024, usage.InodesUsedPercent, usage.UsedPercent)
			} else {
				fmt.Printf("%-30s \t %-20d \t %-20d \t %-15f \t %f \n", usage.Path, usage.Total/1000/1024, usage.Free/1000/1024, usage.InodesUsedPercent, usage.UsedPercent)
			}
		}
	}
}
