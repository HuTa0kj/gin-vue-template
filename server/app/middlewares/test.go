package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Test Global Middleware
func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
