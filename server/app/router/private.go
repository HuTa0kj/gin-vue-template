package router

import (
	"github.com/gin-gonic/gin"

	"gintemplate/app/controller"
	"gintemplate/app/global"
	"gintemplate/app/middlewares/normal"
)

func PrivateRouteGroup(c *gin.Engine) {
	privateGroup := c.Group(global.RoutePrefix)
	privateGroup.Use(normal.UserAuth())

	privateGroup.GET("/login/status", controller.CheckLoginStatus)

}
