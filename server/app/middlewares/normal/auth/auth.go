package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/models/resp"
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
	fmt.Println(userName, userID, userRole)
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

	c.Set("username", userName)
	c.Set("userID", intUserID)
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
