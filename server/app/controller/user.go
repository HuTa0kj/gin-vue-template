package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gintemplate/app/global"
	"gintemplate/app/models/resp"
	"gintemplate/app/services"
)

// Get User Info
func GetCurrentUserInfo(c *gin.Context) {
	userInfo, ok := services.GetUserInfo(c)
	if !ok {
		c.JSON(
			http.StatusNotFound,
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
