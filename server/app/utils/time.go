package utils

import "time"

func CurrentTime() string {
	// 获取当前时间
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}

func CurrentDay() string {
	// 获取当前日期
	t := time.Now().Format("2006-01-02")
	return t
}
