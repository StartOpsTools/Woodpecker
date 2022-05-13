package main

import (
	"fmt"
	"github.com/StartOpsTools/Woodpecker/pkg/os"
)

func main() {
	os.GetOpenFile()

	os.OutputCpuInfo()

	fmt.Print("\n\n")

	os.OutputDiskUsageInfo()

	fmt.Print("\n\n")

	os.MemoryInfo()

	os.Connection("all")

	os.Process()
}
