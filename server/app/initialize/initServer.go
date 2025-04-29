package initialize

import (
	"fmt"
	
	"github.com/gin-gonic/gin"

	"gintemplate/app/config"
	"gintemplate/app/database"
	"gintemplate/app/global"
	"gintemplate/app/logger"
	"gintemplate/app/middlewares"
	"gintemplate/app/router"
)

func InitServer() {
	r := gin.Default()

	// 初始化配置
	err := config.InitConfig()
	if err != nil {
		logger.LogRus.Error(err)
		return
	}

	// 配置信任代理
	_ = r.SetTrustedProxies(config.ConfigInfo.Server.TrustedProxies)

	// 加载全局变量
	global.InitGlobal()

	// 初始化日志管理器
	logger.InitLog()

	// 初始化全局中间件
	middlewares.InitMiddleware(r)

	// 初始化路由
	router.InitRoute(r)

	// 数据库初始化
	err = database.InitDatabase()
	if err != nil {
		logger.LogRus.Error(err)
		return
	}

	err = r.Run(fmt.Sprintf(":%d", config.ConfigInfo.Server.Port))
	if err != nil {
		logger.LogRus.Error(err)
		return
	}
}
