package global

import (
	"fmt"

	"gintemplate/app/logger"
	"github.com/gin-gonic/gin"
)

func ReqUrlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求路径
		logger.LogRus.Info(fmt.Sprintf("[%s] %s", c.Request.Method, c.Request.URL))

		// 继续处理请求
		c.Next()
	}
}
