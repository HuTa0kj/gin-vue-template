package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/logger"
	"gintemplate/app/models/req"
	"gintemplate/app/models/resp"
	"gintemplate/app/services"
)

func LoginCheck(c *gin.Context) {
	var loginJson req.LoginReq
	if err := c.ShouldBindJSON(&loginJson); err != nil {
		c.JSON(400, resp.LoginResp{
			Code:   global.CodeParameterMissing,
			Msg:    global.CodeParameterMissingMsg,
			Token:  "",
			Status: "error",
		})
		return
	}

	token, err := services.Login(loginJson.Username, loginJson.Password) // 具体验证逻辑
	if err != nil {
		c.JSON(http.StatusUnauthorized, resp.LoginResp{
			Code:   global.CodeLoginFail,
			Msg:    global.CodeLoginFailMsg,
			Token:  "",
			Status: "error",
		})
		return
	}
	logger.LogRus.Infof("用户 %s 登录成功", loginJson.Username)
	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, resp.LoginResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Token:  token,
		Status: "ok",
	})
	return
}

// 检验登录状态
func CheckLoginStatus(c *gin.Context) {
	_, NameExists := c.Get("username")
	r, RoleNameExists := c.Get("role")
	if !NameExists || !RoleNameExists {
		c.JSON(
			http.StatusUnauthorized,
			resp.LoginStatusResp{
				Code:   global.CodeUnauthorized,
				Msg:    global.CodeUnauthorizedMsg,
				Status: "error",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.LoginStatusResp{
			Code:   global.CodeSuccess,
			Msg:    global.CodeSuccessMsg,
			Status: "ok",
			Auth:   true,
			Role:   r.(int),
		})
	return
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"code":   2000,
		"status": "ok",
	})
}
