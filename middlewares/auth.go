package middlewares

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/majoo-pos/utilities"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  "unauthorized",
				"message": "No token detected",
			})
			return
		}

		splitToken := strings.Split(tokenString, " ")

		if splitToken[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  "unauthorized",
				"message": "Invalid token",
			})
			return
		} else {
			tokenString = splitToken[1]
		}

		//Parse token menjadi return data
		token, err := utilities.ValidateToken(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  "unauthorized",
				"message": err.Error(),
			})
			return
		}
		if token == nil {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  "unauthorized",
				"message": err.Error(),
			})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if !token.Valid || !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  "unauthorized",
				"message": err.Error(),
			})
			return
		}
		payload := claims["payload"].(map[string]interface{})
		c.Set("id", int(payload["id"].(float64)))
		c.Set("username", payload["username"])
		c.Set("token", token.Raw)
		c.Next()
	}
}
