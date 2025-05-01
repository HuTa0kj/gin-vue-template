package resp

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
