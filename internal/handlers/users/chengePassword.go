package users

import (
	"awesomeProject/internal/Cash"
	"awesomeProject/internal/database"
	"awesomeProject/pkg/hash"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
)

type changeData struct {
	Code        string `json:"code"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

func ChangePassword(ctx *gin.Context) {
	var changeDataUser changeData

	if err := ctx.ShouldBindJSON(&changeDataUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	slog.Info(changeDataUser.Code)
	slog.Info(changeDataUser.Email)

	userCode, err := Cash.RedisClient.Get(ctx.Request.Context(), changeDataUser.Email).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": "Код не найден или истёк"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if userCode != changeDataUser.Code {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "код не верный"})
		return
	}

	hashPassword, err := hash.PasswordHash(changeDataUser.NewPassword)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	_, err = database.DB.Exec("UPDATE users SET password=$1 WHERE email=$2", hashPassword, changeDataUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "пароль заменен"})
}
