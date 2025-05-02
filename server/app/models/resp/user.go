package resp

import (
	"time"
)

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
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Users  []SimpleUser `json:"users"`
	Status string       `json:"status"`
	Total  int64        `json:"total"`
}

// Display User Model
type SimpleUser struct {
	ID            uint      `json:"id"`
	Username      string    `json:"username"`
	Role          string    `json:"role"`
	RegisterTime  time.Time `json:"register_time"`
	LastLoginTime time.Time `json:"last_login_time"`
	Status        bool      `json:"status"`
}
