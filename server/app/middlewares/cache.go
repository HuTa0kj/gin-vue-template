package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"gintemplate/app/database"
	"gintemplate/app/logger"
)

// responseWriter 自定义 ResponseWriter 用于捕获响应体
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func CacheMiddleware(expireTime time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 只缓存GET请求，非幂等不缓存
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		userID, _ := c.Get("user_id")
		url := c.Request.URL
		path := c.Request.URL.Path
		cacheKey := fmt.Sprintf("%s:%s:%d", c.Request.Method, url, userID)

		cachedData, err := database.RedisClient.Get(c, cacheKey).Bytes()
		if err == nil {
			var result struct {
				Route string          `json:"route"`
				Data  json.RawMessage `json:"data"` // 使用RawMessage保留原始JSON
			}

			if err := json.Unmarshal(cachedData, &result); err == nil && result.Route == path {
				logger.LogRus.Info(fmt.Sprintf("cache hit %s", cacheKey))

				c.Writer.Header().Set("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusOK)
				c.Writer.Write(result.Data)
				c.Abort()
				return
			}
		} else if err != redis.Nil {
			logger.LogRus.WithError(err).WithField("cache_key", cacheKey).Error("redis error")
		}

		// 创建自定义 ResponseWriter 来捕获响应
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		c.Next()

		if c.Writer.Status() != http.StatusOK {
			return
		}

		contentType := c.Writer.Header().Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			return
		}

		cacheValue := map[string]interface{}{
			"route": path,
			"data":  json.RawMessage(writer.body.Bytes()),
		}

		jsonData, err := json.Marshal(cacheValue)
		if err != nil {
			logger.LogRus.WithError(err).WithField("cache_key", cacheKey).Error("marshal cache value failed")
			return
		}

		if err := database.RedisClient.Set(c, cacheKey, jsonData, expireTime).Err(); err != nil {
			logger.LogRus.WithError(err).WithField("cache_key", cacheKey).Error("set cache failed")
			return
		}

		logger.LogRus.Info(fmt.Sprintf("Cache set %s", cacheKey))
	}
}
