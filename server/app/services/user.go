package services

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/url"
	"path"

	"gintemplate/app/config"
	"gintemplate/app/database"
	"gintemplate/app/global"
	"gintemplate/app/logger"
	"gintemplate/app/models/db"
	"gintemplate/app/models/sys"
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
func SearchSingleUserInfo(username string) (sys.SimpleUser, error) {
	var dbUser db.User
	err := database.DB.Where("username = ?", username).First(&dbUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return sys.SimpleUser{}, nil
	}
	if err != nil {
		logger.LogRus.Error(err)
		return sys.SimpleUser{}, err
	}

	// db.User → sys.SimpleUser
	simpleUser := sys.SimpleUser{
		ID:            uint(dbUser.ID),
		Username:      dbUser.UserName,
		Role:          dbUser.Role,
		RegisterTime:  dbUser.RegisterTime,
		LastLoginTime: dbUser.LastLoginTime,
		Status:        dbUser.Status,
	}
	return simpleUser, nil
}

func CreatePasswordResetLink(u string) (string, int, error) {
	var user db.User
	if err := database.DB.Where("username = ?", u).First(&user).Error; err != nil {
		return "", global.CodeInformationNotFound, fmt.Errorf(global.CodeInformationNotFoundMsg)
	}
	resetKey := utils.GenerateUUIDv7()
	if err := database.DB.Model(&user).Update("password", resetKey).Error; err != nil {
		return "", global.CodeDatabaseUpdateError, fmt.Errorf(global.CodeDatabaseUpdateErrorMsg)
	}
	baseUrl := config.ConfigInfo.Server.BaseUrl
	base, _ := url.Parse(baseUrl)
	base.Path = path.Join(base.Path, "/password-reset")
	query := url.Values{}
	query.Set("username", u)
	query.Set("reset_key", resetKey)
	base.RawQuery = query.Encode()

	l := base.String()
	return l, global.CodeSuccess, nil
}
