package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/models/resp"
	"gintemplate/app/services"
	"gintemplate/app/utils"
)

// 基础验证与权限校验
func baseAuthHelper(c *gin.Context, minRole int) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			resp.AuthResp{
				Code:   global.CodeUnauthorized,
				Msg:    global.CodeUnauthorizedMsg,
				Status: "error",
			})
		c.Abort()
		return
	}

	claims, code, m, err := utils.CheckJWT(tokenString)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			resp.AuthResp{
				Code:   code,
				Msg:    m,
				Status: "error",
			})
		c.Abort()
		return
	}

	userName, uok := claims["username"].(string)
	userID, iok := claims["id"].(float64)
	userRole, rok := claims["role"].(float64)
	userStatus, _ := claims["status"].(bool)
	if !uok || !rok || !iok {
		c.JSON(
			http.StatusUnauthorized,
			resp.AuthResp{
				Code:   global.CodeTokenParsingFailed,
				Msg:    global.CodeTokenParsingFailedMsg,
				Status: "error",
			})
		c.Abort()
		return
	}

	intUserID := int(userID)
	intUserRole := int(userRole)

	if intUserRole < minRole {
		c.JSON(http.StatusUnauthorized, resp.AuthResp{
			Code:   global.CodeUnauthorized,
			Msg:    global.CodeUnauthorizedMsg,
			Status: "error",
		})
		c.Abort()
		return
	}
	if !userStatus {
		c.JSON(http.StatusUnauthorized, resp.AuthResp{
			Code:   global.CodeUserDisable,
			Msg:    global.CodeUserDisableMsg,
			Status: "error",
		})
		c.Abort()
		return
	}

	c.Set("username", userName)
	c.Set("user_id", intUserID)
	c.Set("role", intUserRole)

	c.Next()
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		baseAuthHelper(c, global.RoleCommonUser)
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		baseAuthHelper(c, global.RoleAdminUser)
	}
}

func RootAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		baseAuthHelper(c, global.RoleRootUser)
	}
}

func apiKeyAuthHelper(c *gin.Context, minRole int) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	u, ok := services.ResolveUserKey(tokenString)
	if !ok {
		c.JSON(
			http.StatusUnauthorized,
			resp.AuthResp{
				Code:   global.CodeUnauthorized,
				Msg:    global.CodeUnauthorizedMsg,
				Status: "error",
			})
		c.Abort()
		return
	}
	if u.Role < minRole {
		c.JSON(http.StatusUnauthorized, resp.AuthResp{
			Code:   global.CodeUnauthorized,
			Msg:    global.CodeUnauthorizedMsg,
			Status: "error",
		})
		c.Abort()
		return
	}
	c.Set("user_id", u.ID)
	c.Next()
}

func UserKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyAuthHelper(c, global.RoleCommonUser)
	}
}

func AdminKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyAuthHelper(c, global.RoleAdminUser)
	}
}

func RootKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyAuthHelper(c, global.RoleRootUser)
	}
}
