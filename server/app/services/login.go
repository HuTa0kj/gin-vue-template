package services

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gintemplate/app/database"
	"gintemplate/app/global"
	"gintemplate/app/logger"
	"gintemplate/app/models/db"
	"gintemplate/app/utils"
)

func Login(username, password string) (string, int, error) {
	if username == "" || password == "" {
		return "", http.StatusBadRequest, fmt.Errorf(global.CodeParameterMissingMsg)
	}
	var user db.User

	// 从数据库查询用户
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return "", http.StatusUnauthorized, fmt.Errorf(global.CodeLoginFailMsg)
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", http.StatusUnauthorized, fmt.Errorf(global.CodeLoginFailMsg)
	}

	// 密码正确，生成 JWT Token
	userName := user.GetUserName()
	userRole := user.GetRole()
	userID := user.GetID()
	token, err := utils.GenerateJWT(userName, userID, userRole)
	if err != nil {
		logger.LogRus.Error("Failed to generate JWT Token: %s", err)
		return "", http.StatusInternalServerError, fmt.Errorf("Failed to generate JWT Token")
	}

	// 返回生成的 JWT Token
	return token, http.StatusOK, nil
}
