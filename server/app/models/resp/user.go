package resp

import "gintemplate/app/models/sys"

type UserInfoResp struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Status   string `json:"status"` // ok error
	UserName string `json:"username"`
	UserID   int    `json:"id"`
	UserRole int    `json:"role"`
}

type UserTokenResp struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Status   string `json:"status"` // ok error
	Token    string `json:"token"`
	UserName string `json:"username"`
}

type AllUserResp struct {
	Code   int              `json:"code"`
	Msg    string           `json:"msg"`
	Users  []sys.SimpleUser `json:"users"`
	Status string           `json:"status"`
	Total  int64            `json:"total"`
}

type ResetPasswordResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Link   string `json:"link"`
	Status string `json:"status"`
}

type UpdateUserInfoResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status string `json:"status"`
}
