package utils

import (
	"fmt"
	"time"
	
	"github.com/dgrijalva/jwt-go"

	"gintemplate/app/global"
)

func GenerateJWT(username string, userId int, userRole int, userStatus bool) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"id":       userId,
		"role":     userRole,
		"status":   userStatus,
		"sub":      "auth",
		"exp":      time.Now().UTC().Add(time.Hour * global.EffectiveHours).Unix(), // 设置 token 的过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(global.JwtKey)
}

// 验证 JWT 信息并返回 jwt.MapClaims 结构
func CheckJWT(tokenString string) (jwt.MapClaims, int, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用 HS256 签名方法
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Not the expected signature: %v", token.Method)
		}
		// 返回密钥，用于验证 Token 的签名
		return global.JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, global.CodeTokenExpired, global.CodeTokenExpiredMsg, fmt.Errorf("Invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, global.CodeTokenInvalid, global.CodeTokenInvalidMsg, fmt.Errorf("Invalid Token")
	}
	return claims, global.CodeSuccess, global.CodeSuccessMsg, nil
}
