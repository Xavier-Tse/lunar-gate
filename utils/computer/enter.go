package computer

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetCpuPercent() float64 {
	percentages, _ := cpu.Percent(time.Second, true)
	return percentages[0]
}

func GetMemPercent() float64 {
	v, _ := mem.VirtualMemory()
	return v.UsedPercent
}

func GetDiskPercent() float64 {
	usage, _ := disk.Usage("/")
	return usage.UsedPercent
}
