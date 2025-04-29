package resp

type LoginResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
	Status string `json:"status"` // ok error
}

type LoginStatusResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status string `json:"status"` // ok error
	Auth   bool   `json:"auth"`
	Role   int    `json:"role"`
}
