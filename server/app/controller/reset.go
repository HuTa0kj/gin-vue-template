package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/models/req"
	"gintemplate/app/models/resp"
	"gintemplate/app/services"
)

func InviteUser(c *gin.Context) {
	var inviteReq req.InviteUserReq
	if err := c.ShouldBindJSON(&inviteReq); err != nil {
		c.JSON(http.StatusBadRequest, resp.InviteUserResp{
			Code:   global.CodeParameterMissing,
			Msg:    global.CodeParameterMissingMsg,
			Link:   "",
			Status: "error",
		})
		return
	}
	l, code, err := services.CreateInviteLink(inviteReq.Username, inviteReq.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.InviteUserResp{
			Code:   code,
			Msg:    err.Error(),
			Status: "error",
			Link:   "",
		})
		return
	}
	c.JSON(http.StatusOK, resp.InviteUserResp{
		Code:   global.CodeSuccess,
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
			Code:   global.CodeParameterMissing,
			Msg:    global.CodeParameterMissingMsg,
			Status: "error",
		})
		return
	}
	code, err := services.UserPasswordReset(inviteCheck.Username, inviteCheck.Password, inviteCheck.InviteKey)
	if err != nil {
		c.JSON(http.StatusOK, resp.InviteCheckResp{
			Code:   code,
			Msg:    err.Error(),
			Status: "error",
		})
		return
	}
	c.JSON(http.StatusOK, resp.InviteCheckResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}

func ResetUserCheck(c *gin.Context) {
	var resetCheck req.ResetPasswordCheckReq
	if err := c.ShouldBindJSON(&resetCheck); err != nil {
		c.JSON(http.StatusBadRequest, resp.InviteCheckResp{
			Code:   global.CodeParameterMissing,
			Msg:    global.CodeParameterMissingMsg,
			Status: "error",
		})
		return
	}
	code, err := services.UserPasswordReset(resetCheck.Username, resetCheck.Password, resetCheck.ResetKey)
	if err != nil {
		c.JSON(http.StatusOK, resp.ResetCheckResp{
			Code:   code,
			Msg:    err.Error(),
			Status: "error",
		})
		return
	}
	c.JSON(http.StatusOK, resp.ResetCheckResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}
