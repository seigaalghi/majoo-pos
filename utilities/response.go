package utilities

import "github.com/gin-gonic/gin"

func ErrorResponse(ctx *gin.Context, err error, code int) {
	ctx.JSON(code, gin.H{
		"status":  "failed",
		"message": err.Error(),
		"result":  gin.H{},
	})
}
