package router

import (
	"github.com/gin-gonic/gin"

	"gintemplate/app/controller"
	"gintemplate/app/global"
)

func PublicRouteGroup(c *gin.Engine) {
	publicGroup := c.Group(global.RoutePrefix)

	loginGroup := publicGroup.Group("login")
	{
		loginGroup.POST("verify", controller.LoginCheck)
	}

	publicGroup.POST("/user/invite/check", controller.InviteUserCheck)

}
