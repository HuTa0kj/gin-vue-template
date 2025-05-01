package router

import (
	"gintemplate/app/controller"
	"gintemplate/app/global"
	"github.com/gin-gonic/gin"
)

func PublicRouteGroup(c *gin.Engine) {
	publicGroup := c.Group(global.RoutePrefix)

	loginGroup := publicGroup.Group("/login")
	{
		loginGroup.POST("verify", controller.LoginCheck)
	}
	
	publicGroup.POST("/user/invite/check", controller.InviteUserCheck)

}
