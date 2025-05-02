package db

import "time"

// User 用户模型
type User struct {
	ID            int       `gorm:"primaryKey;column:id"`
	UserName      string    `gorm:"size:100;not null;column:username"`
	Password      string    `gorm:"size:100;not null;column:password"`
	ApiKey        string    `gorm:"size:100;not null;column:api_key"`
	Role          int       `gorm:"size:20;not null;column:role"`
	RegisterTime  time.Time `gorm:"column:;not null;register_time"`
	LastLoginTime time.Time `gorm:"column:last_login_time"`
	Status        bool      `gorm:"size:20;not null;column:status"`
}

func (u User) GetUserName() string {
	return u.UserName
}

func (u User) GetRole() int {
	return u.Role
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetStatus() bool {
	return u.Status
}
