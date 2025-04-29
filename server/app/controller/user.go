package controller

import (
	"net/http"

	"gintemplate/app/global"
	"gintemplate/app/models/resp"
	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetCurrentUserInfo(c *gin.Context) {
	userName, NameExists := c.Get("username")
	userUid, uidExists := c.Get("userUid")
	if !NameExists || !uidExists {
		c.JSON(
			http.StatusUnauthorized,
			resp.UserInfoResp{
				Code:     global.CodeUnauthorized,
				Msg:      global.CodeUnauthorizedMsg,
				Status:   "error",
				UserName: "",
				UserUid:  "",
			})
		return
	}
	c.JSON(
		http.StatusOK,
		resp.UserInfoResp{
			Code:     global.CodeSuccess,
			Msg:      global.CodeSuccessMsg,
			Status:   "ok",
			UserName: userName.(string),
			UserUid:  userUid.(string),
		})
	return
}
