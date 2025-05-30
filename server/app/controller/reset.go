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

func InviteUser(c *gin.Context) {
	var inviteReq req.InviteUserReq
	if err := c.ShouldBindJSON(&inviteReq); err != nil {
		c.JSON(http.StatusBadRequest, resp.InviteUserResp{
			Msg:    global.CodeParameterMissingMsg,
			Link:   "",
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	l, err := services.CreateInviteLink(inviteReq.Username, inviteReq.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.InviteUserResp{
			Msg:    err.Error(),
			Status: "error",
			Link:   "",
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.InviteUserResp{
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
		Link:   l,
	})
	return
}

func InviteUserCheck(c *gin.Context) {
	var inviteCheck req.InviteUserRegisterReq
	if err := c.ShouldBindJSON(&inviteCheck); err != nil {
		c.JSON(http.StatusBadRequest, resp.InviteCheckResp{
			Msg:    global.CodeParameterMissingMsg,
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	err := services.UserTmpPasswordReset(inviteCheck.Username, inviteCheck.Password, inviteCheck.InviteKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.InviteCheckResp{
			Msg:    err.Error(),
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.InviteCheckResp{
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}

func ResetUserCheck(c *gin.Context) {
	var resetCheck req.ResetPasswordCheckReq
	if err := c.ShouldBindJSON(&resetCheck); err != nil {
		c.JSON(http.StatusBadRequest, resp.InviteCheckResp{
			Msg:    global.CodeParameterMissingMsg,
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	err := services.UserTmpPasswordReset(resetCheck.Username, resetCheck.Password, resetCheck.ResetKey)
	if err != nil {
		c.JSON(http.StatusOK, resp.ResetCheckResp{
			Msg:    err.Error(),
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.ResetCheckResp{
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}
