package resp

type UpdatePasswordResp struct {
	Msg    string `json:"msg"`
	Status string `json:"status"` // ok error
}
