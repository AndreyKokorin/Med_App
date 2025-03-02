package helps

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func RespWithError(ctx *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		slog.Error(message + ": " + err.Error())
		ctx.JSON(statusCode, gin.H{"error": message})
		return
	}

	ctx.JSON(statusCode, gin.H{"message": message})
}
