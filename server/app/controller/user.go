package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/logger"
	"gintemplate/app/models/req"
	"gintemplate/app/models/resp"
	"gintemplate/app/models/sys"
	"gintemplate/app/services"
)

// Get User Info
func GetUserInfoFromKey(c *gin.Context) {
	userInfo, ok := services.GetCurrentUserInfo(c)
	if !ok {
		c.JSON(
			http.StatusNotFound,
			resp.UserInfoResp{
				Msg:      global.CodeInformationNotFoundMsg,
				Status:   "error",
				UserName: "",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.UserInfoResp{
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
				Msg:           global.CodeInformationNotFoundMsg,
				Status:        "error",
				UserName:      "",
				UserRole:      0,
				LastLoginTime: "",
				RegisterTime:  "",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.UserInfoResp{
			Msg:           global.CodeSuccessMsg,
			Status:        "ok",
			UserName:      userInfo.UserName,
			UserID:        userInfo.ID,
			UserRole:      userInfo.Role,
			RegisterTime:  userInfo.RegisterTime,
			LastLoginTime: userInfo.LastLoginTime,
		})
	return
}

func GetUserToken(c *gin.Context) {
	userInfo, ok := services.GetCurrentUserInfo(c)
	if !ok {
		c.JSON(
			http.StatusNotFound,
			resp.UserTokenResp{
				Msg:      global.CodeInformationNotFoundMsg,
				Status:   "error",
				UserName: "",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.UserTokenResp{
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
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
		})
		return
	}
	if err := c.ShouldBindJSON(&ur); err != nil {
		c.JSON(http.StatusBadRequest, resp.UpdatePasswordResp{
			Msg:    global.CodeParameterMissingMsg,
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	if err := services.ModifyUserPassword(c, ur.OldPassword, ur.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, resp.UpdatePasswordResp{
			Msg:    err.Error(),
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.UpdatePasswordResp{
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}

func GetAllUserInfo(c *gin.Context) {
	var aur req.AllUserReq

	if err := c.ShouldBindQuery(&aur); err != nil {
		c.JSON(http.StatusBadRequest, resp.AllUserResp{
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
			Users:  []sys.SimpleUser{},
			Total:  0,
		})
		logger.LogRus.Error(err)
		return
	}
	aur.SetDefaults()
	users, total, err := services.SelectAllUserInfo(aur.Page, aur.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.AllUserResp{
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
			Users:  []sys.SimpleUser{},
			Total:  0,
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.AllUserResp{
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
		c.JSON(http.StatusBadRequest, resp.AllUserResp{
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
			Users:  []sys.SimpleUser{},
		})
		logger.LogRus.Error(err)
		return
	}
	users, err := services.SearchKeywordUserInfo(ur.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.AllUserResp{
			Msg:    global.CodeDatabaseSelectErrorMsg,
			Status: "error",
			Users:  users,
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.AllUserResp{
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
		Users:  users,
	})
	return
}

func ResetPassword(c *gin.Context) {
	var ur req.ResetPasswordReq
	if err := c.ShouldBindJSON(&ur); err != nil {
		c.JSON(http.StatusBadRequest, resp.ResetPasswordResp{
			Msg:    global.CodeParameterIllegalMsg,
			Link:   "",
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	l, err := services.CreatePasswordResetLink(ur.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.ResetPasswordResp{
			Msg:    err.Error(),
			Link:   "",
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.ResetPasswordResp{
		Msg:    global.CodeSuccessMsg,
		Link:   l,
		Status: "ok",
	})
	return
}

// 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	var uq req.UpdateUserInfoReq
	if err := c.ShouldBindJSON(&uq); err != nil {
		c.JSON(http.StatusBadRequest, resp.UpdateUserInfoResp{
			Msg:    global.CodeParameterIllegalMsg,
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	err := services.ModifyUserInfo(uq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.UpdateUserInfoResp{
			Msg:    err.Error(),
			Status: "error",
		})
		logger.LogRus.Error(err)
		return
	}
	c.JSON(http.StatusOK, resp.UpdateUserInfoResp{
		Msg:    global.CodeSuccessMsg,
		Status: "ok",
	})
	return
}
