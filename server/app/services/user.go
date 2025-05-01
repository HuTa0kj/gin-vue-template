package services

import (
	"github.com/gin-gonic/gin"

	"gintemplate/app/database"
	"gintemplate/app/models/db"
)

// Select User Info
func GetKeyUserInfo(c *gin.Context) (db.User, bool) {
	var user db.User
	apiKey, ok := c.Get("api_key")
	if !ok {
		return db.User{}, false
	}
	err := database.DB.Where("api_key = ?", apiKey).First(&user).Error
	if err != nil {
		return db.User{}, false
	}
	return user, true
}

func GetUserInfo(c *gin.Context) (db.User, bool) {
	var user db.User
	user_id, ok := c.Get("user_id")
	if !ok {
		return db.User{}, false
	}
	err := database.DB.Where("id = ?", user_id).First(&user).Error
	if err != nil {
		return db.User{}, false
	}
	return user, true
}
