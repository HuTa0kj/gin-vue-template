package resp

type LoginResp struct {
	Msg    string `json:"msg"`
	Token  string `json:"token"`
	Status string `json:"status"` // ok error
}

type LoginStatusResp struct {
	Msg    string `json:"msg"`
	Status string `json:"status"` // ok error
	Auth   bool   `json:"auth"`
	Role   int    `json:"role"`
}
