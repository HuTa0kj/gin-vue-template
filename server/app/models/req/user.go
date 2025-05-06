package req

type AllUserReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (r *AllUserReq) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 10
	}
}

type UserSearchReq struct {
	Username string `json:"username"`
}

type ResetPasswordReq struct {
	Username string `json:"username"`
}
