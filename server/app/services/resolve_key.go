package services

import (
	"gintemplate/app/database"
	"gintemplate/app/models/db"
)

// Resolve API KEY
func ResolveUserKey(t string) (db.User, bool) {
	if len(t) != 36 {
		return db.User{}, false
	}

	var user db.User
	err := database.DB.Where("api_key = ?", t).First(&user).Error
	if err != nil {
		return db.User{}, false
	}
	return user, true
}
