package req

type InviteUserReq struct {
	Username string `json:"username" binding:"required"`
	Role     int    `json:"role" binding:"required"`
}

type InviteUserRegisterReq struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	InviteKey string `json:"invite_key" binding:"required"`
}
