package os

import (
	"fmt"
	"github.com/StartOpsTools/Woodpecker/pkg/output"
	"github.com/shirou/gopsutil/net"
)

func Connection(king string) {
	connectionsStatus, err := net.Connections(king)
	if err == nil {
		formatOutputConnectionsStatus(connectionsStatus)
	}
}

func PidConnection(king string, pid int32) {
	connectionsStatus, err := net.ConnectionsPid(king, pid)
	if err == nil {
		formatOutputConnectionsStatus(connectionsStatus)
	}
}

// 发送字节数，发送包数， 丢弃包等
func DeviceIOCounts() {
	ist, err := net.IOCounters(true)
	if err == nil {
		for _, i := range ist {
			fmt.Println("i: ", i)
		}
	}
}
// 列出所有网卡
func Device(){
	net.Interfaces()
}



func formatOutputConnectionsStatus(connectionsStatus []net.ConnectionStat) {
	output.Title.Print("Socket连接信息:\n")
	_, _ = output.Title.Printf("%-16s %-10s %-16s %-10s %-12s %s\n", "LocalIP", "LocalPort", "RemoteIP", "RemotePort", "Status", "Pid")
	for _, connectionStatus := range connectionsStatus {

		fmt.Printf("%-16s %-10d %-16s %-10d %-12s %d\n", connectionStatus.Laddr.IP, connectionStatus.Laddr.Port,
			connectionStatus.Raddr.IP,connectionStatus.Raddr.Port, connectionStatus.Status, connectionStatus.Pid)
	}
}
