package controller

import (
	"gintemplate/app/services"
	"net/http"

	"gintemplate/app/global"
	"gintemplate/app/models/resp"
	"github.com/gin-gonic/gin"
)

// Get User Info
func GetCurrentUserInfo(c *gin.Context) {
	userInfo, ok := services.GetUserInfo(c)
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			resp.UserInfoResp{
				Code:     global.CodeInformationNotFound,
				Msg:      global.CodeInformationNotFoundMsg,
				Status:   "error",
				UserName: "",
			})
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
