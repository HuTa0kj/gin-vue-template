package logger

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"gintemplate/app/config"
	"gintemplate/app/utils"
)

var LogRus *logrus.Logger

func InitLog() {
	LogRus = logrus.New()
	debug := config.ConfigInfo.Server.Log.Debug
	save := config.ConfigInfo.Server.Log.SaveFile
	// 默认使用 INFO
	LogRus.SetLevel(logrus.InfoLevel)
	if debug {
		LogRus.SetLevel(logrus.DebugLevel)
	}
	LogRus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 如果需要保存日志到文件中
	if save {
		fileName := "log_" + utils.SnakeCase(utils.CurrentDay()) + ".log"

		logFilePath := path.Join(utils.GetProjectRoot(), "log", fileName)
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		LogRus.SetOutput(file)
	}
	LogRus.SetReportCaller(true)
}
