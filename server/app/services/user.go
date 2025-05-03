package services

import (
	"fmt"
	"gintemplate/app/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"gintemplate/app/database"
	"gintemplate/app/global"
	"gintemplate/app/models/db"
	"gintemplate/app/models/resp"
	"gintemplate/app/utils"
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

func GetCurrentUserInfo(c *gin.Context) (db.User, bool) {
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

func ModifyUserPassword(c *gin.Context, op string, np string) error {
	var user db.User
	username, ok := c.Get("username")
	if !ok {
		return fmt.Errorf(global.CodeInformationNotFoundMsg)
	}
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return fmt.Errorf(global.CodeInformationNotFoundMsg)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(op))
	if err != nil {
		return fmt.Errorf(global.CodeOriginPasswordErrorMsg)
	}
	newHashPwd, _ := utils.HashAndStorePassword(np)
	if err := database.DB.Model(&user).Update("password", newHashPwd).Error; err != nil {
		return fmt.Errorf(global.CodeDatabaseUpdateErrorMsg)
	}
	return nil
}

func SelectAllUserInfo(page, pageSize int) ([]resp.SimpleUser, int64, error) {
	var users []resp.SimpleUser
	var total int64

	offset := (page - 1) * pageSize

	if err := database.DB.Model(&db.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.
		Model(&db.User{}).
		Select("id, username, role, " +
			"DATE_FORMAT(register_time, '%Y-%m-%d %H:%i:%s') as register_time, " +
			"DATE_FORMAT(last_login_time, '%Y-%m-%d %H:%i:%s') as last_login_time, " +
			"status").
		Offset(offset).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		logger.LogRus.Error(err.Error())
		return nil, 0, err
	}

	return users, total, nil
}
