package os

import (
	"fmt"
	"github.com/StartOpsTools/Woodpecker/pkg/output"
	"github.com/shirou/gopsutil/cpu"
	"strings"
	"time"
)

func OutputCpuInfo() {
	output.Title.Println("CPU INFO")
	logicalCpuCount, _ := cpu.Counts(true)
	physicalCpuCount, _ := cpu.Counts(false)

	fmt.Println("logicalCpuCount: ", logicalCpuCount)
	fmt.Println("physicalCpuCount: ", physicalCpuCount)
	fmt.Print("\n")
	cpuInfos, err := cpu.Info()
	if err == nil {
		for _, cpuInfo := range cpuInfos {
			fmt.Printf("CPU Number: %d\n", cpuInfo.CPU)
			fmt.Printf("\tCores: %d\n", cpuInfo.Cores)
			fmt.Printf("\tModelName: %s\n", cpuInfo.ModelName)
			fmt.Printf("\tMhz: %f\n", cpuInfo.Mhz)
			fmt.Printf("\tCacheSize: %d\n", cpuInfo.CacheSize)
			fmt.Printf("\tFlags: %s\n", strings.Join(cpuInfo.Flags, "  "))
			fmt.Print("\n")
		}
	}
}


func OutputCpuUsage() {
	cpuPercent, err := cpu.Percent(1 * time.Second, false)
	if err == nil {
		fmt.Println("cpuPercent: ", cpuPercent)
	}
}