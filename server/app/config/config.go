package config

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/spf13/viper"
)

// MySQL
type MySQL struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// Database
type DB struct {
	MySQL MySQL
	Redis Redis
}

// Auth
type Auth struct {
	JwtKey         string
	EffectiveHours time.Duration
}

// Log
type Log struct {
	Debug    bool
	SaveFile bool
}

// Route
type Route struct {
	Prefix string
}

// Server
type Server struct {
	BaseUrl        string
	Log            Log
	Route          Route
	Port           int
	TrustedProxies []string
}

// Config
type Config struct {
	Server Server
	DB     DB
	Auth   Auth
}

var ConfigInfo Config

// ReadConfig
func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("app/config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("读取配置文件错误：%s", err))
		return err
	}

	if err := viper.Unmarshal(&ConfigInfo); err != nil {
		fmt.Println(fmt.Sprintf("解析配置文件错误：%s", err))
		return err
	}
	// 未设置秘钥则使用随机秘钥
	if ConfigInfo.Auth.JwtKey == "jwt_key" {
		jk := generateRandomKey()
		fmt.Println("未设置JWT秘钥，使用随机秘钥：", jk)
		ConfigInfo.Auth.JwtKey = jk
	}

	fmt.Println("读取配置文件成功")
	return nil
}

func generateRandomKey() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#$%^"
	result := make([]byte, 12)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}
