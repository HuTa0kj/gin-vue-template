package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"gintemplate/app/controller"
	"gintemplate/app/global"
	"gintemplate/app/middlewares"
)

func PrivateRouteGroup(c *gin.Engine) {
	privateGroup := c.Group(global.RoutePrefix)
	privateGroup.Use(middlewares.UserAuth())

	privateGroup.GET("/login/status", controller.CheckLoginStatus)
	privateGroup.GET("/logout", controller.Logout)

	userGroup := privateGroup.Group("user")
	{
		userGroup.GET("token", controller.GetUserToken)
		userGroup.GET("info", controller.GetUserInfoFromCookie)
		userGroup.PATCH("/password/update", controller.UpdatePassword)
	}

	adminGroup := privateGroup.Group("admin")
	adminGroup.Use(middlewares.AdminAuth())
	{
		adminGroup.POST("/user/invite", controller.InviteUser)
		adminGroup.GET("/user/all", middlewares.CacheMiddleware(120*time.Second), controller.GetAllUserInfo)
		adminGroup.POST("/user/search", controller.SearchUserInfo)
		adminGroup.PUT("/user/reset/password", controller.ResetPassword)
		adminGroup.PUT("/user/update", controller.UpdateUserInfo)

	}

	rootGroup := privateGroup.Group("root")
	rootGroup.Use(middlewares.RootAuth())

	statisticsGroup := privateGroup.Group("statistics")
	statisticsGroup.Use(middlewares.UserAuth())
	{
		statisticsGroup.GET("/cpu", controller.SystemStatisticsCPU)
	}

}
