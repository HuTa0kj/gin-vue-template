package services

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

// 获取CPU使用率
func GetCPUUsage() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

// 获取系统启动时间和运行天数
func GetUptimeInfo() (time.Time, int, error) {
	bootTime, err := host.BootTime()
	if err != nil {
		return time.Time{}, 0, err
	}

	uptime := time.Since(time.Unix(int64(bootTime), 0))
	uptimeDays := int(uptime.Hours() / 24)

	return time.Unix(int64(bootTime), 0), uptimeDays, nil
}
