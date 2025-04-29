package resp

type UserInfoResp struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Status   string `json:"status"` // ok error
	UserName string `json:"username"`
	UserUid  string `json:"uid"`
}
