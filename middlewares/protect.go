package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Protect() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("apiKey")
		if apiKey != "apikeydemo" {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  "unauthorized",
				"message": "Invalid API Key",
			})
			return
		}
		c.Next()
	}
}
