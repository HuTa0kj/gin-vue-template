package resp

import "gintemplate/app/models/sys"

type UserInfoResp struct {
	Msg           string `json:"msg"`
	Status        string `json:"status"` // ok error
	UserName      string `json:"username"`
	UserID        int    `json:"id"`
	UserRole      int    `json:"role"`
	RegisterTime  string `json:"register_time"`
	LastLoginTime string `json:"last_login_time"`
}

type UserTokenResp struct {
	Msg      string `json:"msg"`
	Status   string `json:"status"` // ok error
	Token    string `json:"token"`
	UserName string `json:"username"`
}

type AllUserResp struct {
	Msg    string           `json:"msg"`
	Users  []sys.SimpleUser `json:"users"`
	Status string           `json:"status"`
	Total  int64            `json:"total"`
}

type ResetPasswordResp struct {
	Msg    string `json:"msg"`
	Link   string `json:"link"`
	Status string `json:"status"`
}

type UpdateUserInfoResp struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
}
