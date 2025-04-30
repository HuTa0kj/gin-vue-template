package router

import (
	"gintemplate/app/controller"
	"gintemplate/app/global"
	"gintemplate/app/middlewares/normal"
	"github.com/gin-gonic/gin"
)

func ApiRouteGroup(c *gin.Engine) {
	apiGroup := c.Group(global.RoutePrefix)
	apiGroup.Use(normal.UserKeyAuth())

	userGroup := apiGroup.Group("user")
	{
		userGroup.GET("info", controller.GetCurrentUserInfo)
	}

}
