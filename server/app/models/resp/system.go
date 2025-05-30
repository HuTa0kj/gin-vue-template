package resp

type CPUStatisticsResp struct {
	CPUUsage   string `json:"cpu_usage"`   // CPU使用率百分比
	UptimeDays int    `json:"uptime_days"` // 系统运行天数(整数)
	LastBoot   string `json:"last_boot"`   // 最后启动时间
	Msg        string `json:"msg"`         // 状态消息
	Status     string `json:"status"`      // 状态标识
}
