package resp

type InviteUserResp struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
	Link   string `json:"link"`
}

type InviteCheckResp struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
}

type ResetCheckResp struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
}
