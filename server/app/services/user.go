package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/url"
	"path"

	"gintemplate/app/config"
	"gintemplate/app/database"
	"gintemplate/app/global"
	"gintemplate/app/logger"
	"gintemplate/app/models/db"
	"gintemplate/app/models/req"
	"gintemplate/app/models/sys"
	"gintemplate/app/utils"
)

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

func SelectAllUserInfo(page, pageSize int) ([]sys.SimpleUser, int64, error) {
	var users []sys.SimpleUser
	var total int64

	offset := (page - 1) * pageSize

	if err := database.DB.Model(&db.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.
		Model(&db.User{}).
		Select("id, username, role, register_time, last_login_time, status").
		Offset(offset).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		logger.LogRus.Error(err.Error())
		return nil, 0, err
	}

	return users, total, nil
}

func SearchKeywordUserInfo(username string) ([]sys.SimpleUser, error) {
	var dbUsers []db.User
	if len(username) < 2 {
		return nil, fmt.Errorf("username too short")
	}

	err := database.DB.
		Where("username LIKE ?", "%"+username+"%").
		Find(&dbUsers).
		Error

	if err != nil {
		logger.LogRus.Error(err)
		return nil, err
	}

	if len(dbUsers) == 0 {
		return []sys.SimpleUser{}, nil
	}

	var simpleUsers []sys.SimpleUser
	for _, dbUser := range dbUsers {
		simpleUsers = append(simpleUsers, sys.SimpleUser{
			ID:            uint(dbUser.ID),
			Username:      dbUser.UserName,
			Role:          dbUser.Role,
			RegisterTime:  dbUser.RegisterTime,
			LastLoginTime: dbUser.LastLoginTime,
			Status:        dbUser.Status,
		})
	}

	return simpleUsers, nil
}

func CreatePasswordResetLink(u string) (string, error) {
	var user db.User
	if err := database.DB.Where("username = ?", u).First(&user).Error; err != nil {
		return "", fmt.Errorf(global.CodeInformationNotFoundMsg)
	}
	resetKey := utils.GenerateUUIDv7()
	if err := database.DB.Model(&user).Update("password", resetKey).Error; err != nil {
		return "", fmt.Errorf(global.CodeDatabaseUpdateErrorMsg)
	}
	baseUrl := config.ConfigInfo.Server.BaseUrl
	base, _ := url.Parse(baseUrl)
	base.Path = path.Join(base.Path, "/password-reset")
	query := url.Values{}
	query.Set("username", u)
	query.Set("reset_key", resetKey)
	base.RawQuery = query.Encode()

	l := base.String()
	return l, nil
}

func ModifyUserInfo(u req.UpdateUserInfoReq) error {
	if u.Role > 10 || u.Role < 1 {
		return fmt.Errorf(global.CodeParameterIllegalMsg)
	}
	updates := map[string]interface{}{
		"role":   u.Role,
		"status": u.Status,
	}
	result := database.DB.Model(&db.User{}).Where("id = ?", u.ID).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf(global.CodeDatabaseUpdateErrorMsg)
	}
	return nil
}
