package resp

type AuthResp struct {
	Status string `json:"status"` // 状态 ok error
	Msg    string `json:"msg"`
}
