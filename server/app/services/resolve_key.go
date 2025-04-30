package services

import (
	"gintemplate/app/database"
	"gintemplate/app/models/db"
)

// Resolve API KEY
func ResolveUserKey(t string) (int, bool) {
	if len(t) != 36 {
		return 0, false
	}

	var user db.User
	err := database.DB.Where("api_key = ?", t).First(&user).Error
	if err != nil {
		return 0, false
	}
	return user.Role, true
}
