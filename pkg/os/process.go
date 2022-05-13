package os

import (
	"fmt"
	"github.com/StartOpsTools/Woodpecker/pkg/output"
	"github.com/shirou/gopsutil/process"
	"time"
)

func Process() {
	procs, err := process.Processes()
	if err == nil {

		_, _ = output.Title.Println("运行程序信息:")
		output.Title.Printf("%-10s %-10s %-20s %-6s %-6s %-10s\n", "Id", "内存百分比", "启动时间", "RSS(MB)", "CPU百分比", "名字")
		for _, p := range procs {
			name, _ := p.Name()
			createTime, _ := p.CreateTime()
			cpuPercent, _ := p.CPUPercent()
			pid := p.Pid
			createTimeString := time.Unix(createTime/1000, createTime%1000*int64(time.Millisecond)).Format("2006-Jan-2 15:04:05")
			memoryInfoStat, _ := p.MemoryInfo()

			totalMemoryPercent, _ := p.MemoryPercent()

			if totalMemoryPercent > 30 {
				output.Warning.Printf("%-10d %-10f %-30s %-6d %-6f %-10s\n", pid, totalMemoryPercent, createTimeString, memoryInfoStat.RSS/1024/1024, cpuPercent, name)
			} else {
				output.Normal.Printf("%-10d %-10f %-30s %-6d %-6f  %-10s\n", pid, totalMemoryPercent, createTimeString, memoryInfoStat.RSS/1024/1024, cpuPercent, name)
			}
		}
	}
}

func ProcessPidExists(pid int32) bool {
	exists, err := process.PidExists(pid)
	if err != nil {
		fmt.Println("Get Process Pid Failure, err: ", err.Error())
		return false
	}
	return exists
}
