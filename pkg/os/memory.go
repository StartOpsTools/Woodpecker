package os

import (
	"fmt"
	"github.com/StartOpsTools/Woodpecker/pkg/output"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/mem"

)

func MemoryInfo() {
	virtualMemoryStat, err := mem.VirtualMemory()

	if err == nil {
		output.Title.Print("MemoryInformation:\n")
		fmt.Printf("\tTotal: %d MB\n", virtualMemoryStat.Total/1024/1024)
		fmt.Printf("\tFree: %d MB\n", virtualMemoryStat.Free/1024/1024)
		fmt.Printf("\tUsed: %d MB\n", virtualMemoryStat.Used/1024/1024)
		fmt.Printf("\tUsedPercent: %f%%\n", virtualMemoryStat.UsedPercent)
		fmt.Printf("\tDirty: %d  MB\n", virtualMemoryStat.Dirty/1024/1024)
		fmt.Printf("\tBuffers: %d MB\n", virtualMemoryStat.Buffers/1024/1024)
		fmt.Printf("\tCached: %d MB\n", virtualMemoryStat.Cached/1024/1024)
		//fmt.Printf("\tVMallocTotal: %d MB\n", virtualMemoryStat.VMallocTotal/1024/1024)
		//fmt.Printf("\tVMallocUsed: %d MB\n", virtualMemoryStat.VMallocUsed/1024/1024)
	}

	swapMemoryStat, err := mem.SwapMemory()
	if err == nil {
		output.Title.Print("SwapInformation:\n")

		swapMemoryUsedPercent := swapMemoryStat.UsedPercent

		var writer *color.Color
		if swapMemoryUsedPercent >= 10 {
			writer = output.Warning
			writer.Print("\t您的系统有使用 Swap 空间，请检查是否合理.\n\n")
		} else {
			writer = output.Normal
		}

		writer.Printf("\tTotal: %d MB\n", swapMemoryStat.Total/1024/1024)
		writer.Printf("\tUsed: %d MB\n", swapMemoryStat.Used/1024/1024)
		writer.Printf("\tFree: %d MB\n", swapMemoryStat.Free/1024/1024)
		writer.Printf("\tUsedPercent: %f%%\n", swapMemoryStat.UsedPercent)

		fmt.Print("\n")
	}
}
