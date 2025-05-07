package database

import (
	"fmt"
	"gintemplate/app/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	redisConfig := config.ConfigInfo.DB.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
}
