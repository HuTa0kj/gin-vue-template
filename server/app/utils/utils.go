package utils

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

// 获取项目的根目录
func GetProjectRoot() string {
	dir, _ := os.Getwd()
	root, _ := filepath.Abs(dir)
	return root
}

func SnakeCase(input string) string {
	// 蛇形命名法
	var result []rune
	for i, r := range input {
		if unicode.IsUpper(r) {
			// 如果不是第一个字符，添加下划线
			if i > 0 {
				result = append(result, '_')
			}
			// 将大写字母转换为小写
			result = append(result, unicode.ToLower(r))
		} else {
			// 如果是小写字母或其他字符，直接添加
			result = append(result, r)
		}
	}
	res := strings.Replace(string(result), "-", "_", -1)
	res = strings.Replace(res, " ", "", -1)

	return res
}

// 生成 UUIDv7
func GenerateUUIDv7() string {
	u, _ := uuid.NewV7()
	return u.String()
}

// 生成 bcrypt 密码
func HashAndStorePassword(password string) (string, error) {
	// 原始密码长度需要小于 72
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

// 生成随机字符串
func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#$%^"
	result := make([]byte, length)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}

	return string(result)
}
