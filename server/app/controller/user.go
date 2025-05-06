package controller

import (
	"gintemplate/app/logger"
	"net/http"

	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/models/req"
	"gintemplate/app/models/resp"
	"gintemplate/app/models/sys"
	"gintemplate/app/services"
)

// Get User Info
func GetUserInfoFromKey(c *gin.Context) {
	userInfo, ok := services.GetKeyUserInfo(c)
	if !ok {
		c.JSON(
			http.StatusNotFound,
			resp.UserInfoResp{
				Code:     global.CodeInformationNotFound,
				Msg:      global.CodeInformationNotFoundMsg,
				Status:   "error",
				UserName: "",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.UserInfoResp{
			Code:     global.CodeSuccess,
			Msg:      global.CodeSuccessMsg,
			Status:   "ok",
			UserName: userInfo.UserName,
			UserID:   userInfo.ID,
			UserRole: userInfo.Role,
		})
	return
}

func GetUserInfoFromCookie(c *gin.Context) {
	userInfo, ok := services.GetCurrentUserInfo(c)
	if !ok {
		c.JSON(
			http.StatusNotFound,
			resp.UserInfoResp{
				Code:     global.CodeInformationNotFound,
				Msg:      global.CodeInformationNotFoundMsg,
				Status:   "error",
				UserName: "",
				UserRole: 0,
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.UserInfoResp{
			Code:     global.CodeSuccess,
			Msg:      global.CodeSuccessMsg,
			Status:   "ok",
			UserName: userInfo.UserName,
			UserID:   userInfo.ID,
			UserRole: userInfo.Role,
		})
	return
}

func GetUserToken(c *gin.Context) {
	userInfo, ok := services.GetCurrentUserInfo(c)
	if !ok {
		c.JSON(
			http.StatusNotFound,
			resp.UserInfoResp{
				Code:     global.CodeInformationNotFound,
				Msg:      global.CodeInformationNotFoundMsg,
				Status:   "error",
				UserName: "",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.UserTokenResp{
			Code:     global.CodeSuccess,
			Msg:      global.CodeSuccessMsg,
			Status:   "ok",
			UserName: userInfo.UserName,
			Token:    userInfo.ApiKey,
		})
	return
}

// Update user password
func UpdatePassword(c *gin.Context) {
	var ur req.UpdatePasswordReq
	if len(ur.NewPassword) < 6 {
		c.JSON(http.StatusBadRequest, resp.UpdatePasswordResp{
			Code:   global.CodeParameterIllegal,
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
		})
		return
	}
	if err := c.ShouldBindJSON(&ur); err != nil {
		c.JSON(http.StatusBadRequest, resp.UpdatePasswordResp{
			Code:   global.CodeParameterMissing,
			Msg:    global.CodeParameterMissingMsg,
			Status: "error",
		})
		return
	}
	if err := services.ModifyUserPassword(c, ur.OldPassword, ur.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, resp.UpdatePasswordResp{
			Code:   global.CodeDatabaseUpdateError,
			Msg:    err.Error(),
			Status: "error",
		})
		return
	}
	c.JSON(http.StatusOK, resp.UpdatePasswordResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}

func GetAllUserInfo(c *gin.Context) {
	var aur req.AllUserReq

	if err := c.ShouldBindQuery(&aur); err != nil {
		c.JSON(http.StatusBadRequest, resp.AllUserResp{
			Code:   global.CodeParameterIllegal,
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
			Users:  []sys.SimpleUser{},
			Total:  0,
		})
		return
	}
	aur.SetDefaults()
	users, total, err := services.SelectAllUserInfo(aur.Page, aur.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.AllUserResp{
			Code:   global.CodeParameterIllegal,
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
			Users:  []sys.SimpleUser{},
			Total:  0,
		})
		return
	}
	c.JSON(http.StatusOK, resp.AllUserResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
		Users:  users,
		Total:  total,
	})
	return
}

func SearchUserInfo(c *gin.Context) {
	var ur req.UserSearchReq
	if err := c.ShouldBindJSON(&ur); err != nil {
		c.JSON(http.StatusBadRequest, resp.SingleUserResp{
			Code:   global.CodeParameterIllegal,
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
			User:   sys.SimpleUser{},
		})
		return
	}
	user, err := services.SearchSingleUserInfo(ur.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.SingleUserResp{
			Code:   global.CodeDatabaseSelectError,
			Msg:    global.CodeDatabaseSelectErrorMsg,
			Status: "error",
			User:   user,
		})
		return
	}
	c.JSON(http.StatusOK, resp.SingleUserResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
		User:   user,
	})
	return
}

func ResetPassword(c *gin.Context) {
	var ur req.ResetPasswordReq
	if err := c.ShouldBindJSON(&ur); err != nil {
		c.JSON(http.StatusBadRequest, resp.ResetPasswordResp{
			Code:   global.CodeParameterIllegal,
			Msg:    global.CodeParameterIllegalMsg,
			Link:   "",
			Status: "error",
		})
		return
	}
	l, code, err := services.CreatePasswordResetLink(ur.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.ResetPasswordResp{
			Code:   code,
			Msg:    err.Error(),
			Link:   "",
			Status: "error",
		})
		return
	}
	c.JSON(http.StatusOK, resp.ResetPasswordResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Link:   l,
		Status: "ok",
	})
	return
}

func UpdateUserInfo(c *gin.Context) {
	var uq req.UpdateUserInfoReq
	if err := c.ShouldBindJSON(&uq); err != nil {
		c.JSON(http.StatusBadRequest, resp.UpdateUserInfoResp{
			Code:   global.CodeParameterIllegal,
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	code, err := services.ModifyUserInfo(uq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.UpdateUserInfoResp{
			Code:   code,
			Msg:    err.Error(),
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.UpdateUserInfoResp{
		Code:   global.CodeSuccess,
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}
