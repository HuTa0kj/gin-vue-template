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
			Msg:    global.CodeParameterMissingMsg,
			Token:  "",
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}

	token, err := services.Login(loginJson.Username, loginJson.Password) // 具体验证逻辑
	if err != nil {
		logger.LogRus.Warning(err)
		c.JSON(http.StatusUnauthorized, resp.LoginResp{
			Msg:    err.Error(),
			Token:  "",
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	logger.LogRus.Infof("用户 %s 登录成功", loginJson.Username)
	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, resp.LoginResp{
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
				Msg:    global.CodeUnauthorizedMsg,
				Status: "error",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.LoginStatusResp{
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
