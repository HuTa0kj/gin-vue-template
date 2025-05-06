package middlewares

import (
	"github.com/gin-gonic/gin"

	"gintemplate/app/middlewares/global"
)

func InitMiddleware(r *gin.Engine) {
	// 注册全局中间件
	r.Use(global.TestMiddleware())
}
