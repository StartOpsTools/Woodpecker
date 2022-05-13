package os

import (
	"github.com/StartOpsTools/Woodpecker/pkg/output"
	"github.com/shirou/gopsutil/host"
)

func Host() {

	uptime, _ := host.Uptime()
	info, _ := host.Info()
	platform, family, pver ,_ := host.PlatformInformation()
	kernelVersion, _ := host.KernelVersion()
	kernelArch , _ := host.KernelArch()
	hostId , _ := host.HostID()
	bootTime, _ := host.BootTime()
	users, _ := host.Users()

	output.Title.Print("主机信息:\n")
	output.Normal.Printf("\tuptime: %d", uptime)
	output.Normal.Printf("\tuptime: %s", info)
	output.Normal.Printf("\tuptime: %s", platform)
	output.Normal.Printf("\tuptime: %s", family)
	output.Normal.Printf("\tuptime: %s", pver)

	output.Normal.Printf("\tuptime: %s", kernelVersion)
	output.Normal.Printf("\tuptime: %s", kernelArch)

	output.Normal.Printf("\tuptime: %s", hostId)
	output.Normal.Printf("\tuptime: %s", bootTime)
	output.Normal.Printf("\tuptime: %s", users)

}
