package monitor

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type SystemMetrics struct {
	CPUUsage     float64
	MemoryUsage  float64
	DiskUsage    float64
	NetRecvSpeed uint64
	NetSentSpeed uint64
}

// 컴퓨팅 리소스를 수집하기 위한 메서드
func (__this *SystemMetrics) GetSystemMetrics() (*SystemMetrics, error) {
	cpuStat, _ := cpu.Percent(time.Second, false)
	memStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	prevNetStat, _ := net.IOCounters(false)

	// 1초 간격
	time.Sleep(1 * time.Second)
	nextNetStat, _ := net.IOCounters(false)

	__this.CPUUsage = cpuStat[0]
	__this.MemoryUsage = memStat.UsedPercent
	__this.DiskUsage = diskStat.UsedPercent

	// bps
	__this.NetRecvSpeed = (nextNetStat[0].BytesRecv - prevNetStat[0].BytesRecv) * 8
	__this.NetSentSpeed = (nextNetStat[0].BytesSent - prevNetStat[0].BytesSent) * 8

	// mbps
	// __this.NetRecvSpeed = (nextNetStat[0].BytesRecv - prevNetStat[0].BytesRecv) * 8 / 1000000
	// __this.NetSentSpeed = (nextNetStat[0].BytesSent - prevNetStat[0].BytesSent) * 8 / 1000000

	// return &SystemMetrics{
	// 	CPUUsage:     cpuStat[0],
	// 	MemoryUsage:  memStat.UsedPercent,
	// 	DiskUsage:    diskStat.UsedPercent,
	// 	NetRecvSpeed: (nextNetStat[0].BytesRecv - prevNetStat[0].BytesRecv) * 8,
	// 	NetSentSpeed: (nextNetStat[0].BytesSent - prevNetStat[0].BytesSent) * 8,
	// }, nil

	return __this, nil
}
