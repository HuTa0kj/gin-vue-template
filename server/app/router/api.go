package router

import (
	"gintemplate/app/controller"
	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/middlewares/normal"
)

func ApiRouteGroup(c *gin.Engine) {
	apiGroup := c.Group(global.RoutePrefix)
	apiGroup.Use(normal.UserKeyAuth())
	apiGroupV1 := apiGroup.Group("v1")

	userGroup := apiGroupV1.Group("user")
	{
		userGroup.GET("info", controller.GetUserInfoFromKey)
	}

}
