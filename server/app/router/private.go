package router

import (
	"github.com/gin-gonic/gin"

	"gintemplate/app/controller"
	"gintemplate/app/global"
	"gintemplate/app/middlewares/normal/auth"
)

func PrivateRouteGroup(c *gin.Engine) {
	privateGroup := c.Group(global.RoutePrefix)
	privateGroup.Use(auth.UserAuth())

	privateGroup.GET("/login/status", controller.CheckLoginStatus)

	userGroup := privateGroup.Group("user")
	{
		userGroup.GET("info", controller.GetCurrentUserInfo)
	}

}
