package router

import (
	"github.com/gin-gonic/gin"

	"gintemplate/app/controller"
	"gintemplate/app/global"
	"gintemplate/app/middlewares"
)

func ApiRouteGroup(c *gin.Engine) {
	apiGroup := c.Group(global.RoutePrefix)
	apiGroup.Use(middlewares.UserKeyAuth())
	apiGroupV1 := apiGroup.Group("v1")

	userGroup := apiGroupV1.Group("user")
	{
		userGroup.GET("info", controller.GetUserInfoFromKey)
	}

}
