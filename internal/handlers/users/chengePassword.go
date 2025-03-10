package users

import (
	"awesomeProject/internal/Cash"
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/hash"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
)

// ChangePassword изменяет пароль пользователя
// @Summary Смена пароля
// @Description Проверяет код из Redis и меняет пароль пользователя в базе данных
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param input body models.ChangeData true "Данные для смены пароля"
// @Success 200 {object} map[string]string "Пароль успешно изменен"
// @Failure 400 {object} map[string]string "Неверный код или некорректные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера или базы данных"
// @Router /password/change [post]
func ChangePassword(ctx *gin.Context) {
	var changeDataUser models.ChangeData

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
