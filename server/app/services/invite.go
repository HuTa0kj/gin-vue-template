package services

import (
	"fmt"
	"net/url"
	"path"

	"gintemplate/app/config"
	"gintemplate/app/database"
	"gintemplate/app/global"
	"gintemplate/app/models/db"
	"gintemplate/app/utils"
)

func CreateInviteLink(u string, r int) (string, int, error) {
	var existingUser db.User
	if err := database.DB.Where("username = ?", u).First(&existingUser).Error; err == nil {
		return "", global.CodeDuplicateUsername, fmt.Errorf(global.CodeDuplicateUsernameMsg)
	}

	// Create tmp invite key
	inviteKey := utils.GenerateUUIDv7()

	newUser := db.User{
		UserName: u,
		Password: inviteKey,
		ApiKey:   utils.GenerateUUIDv7(),
		Role:     r,
	}
	if r > 10 || r < 1 {
		return "", global.CodeInviteRoleError, fmt.Errorf(global.CodeInviteRoleErrorMsg)
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		return "", global.CodeDatabaseInsertError, fmt.Errorf(global.CodeDatabaseInsertErrorMsg)
	}

	baseUrl := config.ConfigInfo.Server.BaseUrl
	base, _ := url.Parse(baseUrl)
	base.Path = path.Join(base.Path, "/invite")
	query := url.Values{}
	query.Set("username", u)
	query.Set("invite_key", inviteKey)
	base.RawQuery = query.Encode()

	l := base.String()
	return l, global.CodeSuccess, nil
}

func InviteUserRegister(u string, p string, k string) (int, error) {
	// Invite key is tmp password
	var user db.User
	result := database.DB.Where("username = ? AND password = ?", u, k).First(&user)

	if result.Error != nil {
		return global.CodeInviteCodeInvalid, fmt.Errorf(global.CodeInviteCodeInvalidMsg)
	}

	hashPwd, _ := utils.HashAndStorePassword(p)

	if err := database.DB.Model(&user).Update("password", hashPwd).Error; err != nil {
		return global.CodeDatabaseUpdateError, fmt.Errorf(global.CodeDatabaseUpdateErrorMsg)
	}

	return global.CodeSuccess, nil
}
