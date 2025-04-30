package router

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {
	// API 路由
	PublicRouteGroup(r)
	PrivateRouteGroup(r)
	ApiRouteGroup(r)
}
