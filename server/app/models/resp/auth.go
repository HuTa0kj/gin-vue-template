package resp

type AuthResp struct {
	Code   int    `json:"code"`   // 业务代码
	Status string `json:"status"` // 状态 ok error
	Msg    string `json:"msg"`
}
