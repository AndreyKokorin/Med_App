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
// @Summary Изменение пароля пользователя
// @Description Изменяет пароль пользователя на основе email, кода (отправленного на email) и нового пароля
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param changeData body models.СhangeData true "Данные для изменения пароля (email, code, newPassword)"
// @Success 200 {object} map[string]string "Успешное изменение пароля"
// @Failure 400 {object} map[string]string "Неверный формат запроса или неверный код"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера (например, ошибка базы данных, Redis или хеширования пароля)"
// @Router /changePassword [post]
func ChangePassword(ctx *gin.Context) {
	var changeDataUser models.СhangeData

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
