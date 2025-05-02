package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"gintemplate/app/database"
	"gintemplate/app/global"
	"gintemplate/app/logger"
	"gintemplate/app/models/db"
	"gintemplate/app/utils"
)

func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", fmt.Errorf(global.CodeParameterMissingMsg)
	}
	var user db.User

	// 从数据库查询用户
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return "", fmt.Errorf(global.CodeLoginFailMsg)
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf(global.CodeLoginFailMsg)
	}

	// 密码正确，生成 JWT Token
	userName := user.GetUserName()
	userRole := user.GetRole()
	userID := user.GetID()
	userStatus := user.GetStatus()
	if !userStatus {
		logger.LogRus.Error("User is disabled: %s", userName)
		return "", fmt.Errorf(global.CodeUserDisableMsg)
	}
	token, err := utils.GenerateJWT(userName, userID, userRole, userStatus)
	if err != nil {
		logger.LogRus.Error("Failed to generate JWT Token: %s", err)
		return "", fmt.Errorf("Failed to generate JWT Token")
	}

	// 返回生成的 JWT Token
	return token, nil
}
