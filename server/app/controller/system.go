package controller

import (
	"fmt"
	"gintemplate/app/models/resp"
	"gintemplate/app/services"
	"github.com/gin-gonic/gin"
)

func SystemStatisticsCPU(c *gin.Context) {
	cpuPercent, err := services.GetCPUUsage()
	if err != nil {
		c.JSON(500, resp.CPUStatisticsResp{
			CPUUsage:   "0%",
			UptimeDays: 0,
			LastBoot:   "1970-01-01 00:00:00",
			Msg:        fmt.Sprintf("Failed to get CPU usage: %v", err),
			Status:     "error",
		})
		return
	}

	bootTime, uptimeDays, err := services.GetUptimeInfo()
	if err != nil {
		c.JSON(500, resp.CPUStatisticsResp{
			CPUUsage:   fmt.Sprintf("%.1f%%", cpuPercent),
			UptimeDays: 0,
			LastBoot:   "1970-01-01 00:00:00",
			Msg:        fmt.Sprintf("Failed to get uptime: %v", err),
			Status:     "error",
		})
		return
	}

	response := resp.CPUStatisticsResp{
		CPUUsage:   fmt.Sprintf("%.1f%%", cpuPercent),
		UptimeDays: uptimeDays,
		LastBoot:   bootTime.Format("2006-01-02 15:04:05"),
		Msg:        "Successfully retrieved system information",
		Status:     "ok",
	}

	c.JSON(200, response)
}
