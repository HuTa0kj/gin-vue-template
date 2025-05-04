package sys

// Display User Model
type SimpleUser struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Role          int    `json:"role"`
	RegisterTime  string `json:"register_time"`
	LastLoginTime string `json:"last_login_time"`
	Status        bool   `json:"status"`
}
