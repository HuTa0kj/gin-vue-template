package resp

type UpdatePasswordResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status string `json:"status"` // ok error
}
