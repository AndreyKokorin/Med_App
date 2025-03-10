package users

import (
	"awesomeProject/internal/database"
	repositories "awesomeProject/internal/repositories/user"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// GetProfile
// @Summary Получение профиля текущего пользователя
// @Description Возвращает информацию о текущем пользователе на основе его токена
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.User "Данные пользователя"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/profile [get]
func GetProfile(ctx *gin.Context) {
	id, ok := ctx.Get("user_id")

	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "user_id not found", errors.New("user_id not found"))
		return
	}

	slog.Info("user_id: ", id)

	idInt, ok := id.(int)

	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "user_id must be int", errors.New("user_id must be int"))
		return
	}

	user, err := repositories.GetUserId(database.DB, idInt)

	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
		}
		helps.RespWithError(ctx, status, "failed to fetch user", err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
