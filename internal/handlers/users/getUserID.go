package users

import (
	"awesomeProject/internal/database"
	repositories "awesomeProject/internal/repositories/user"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetUserID
// @Summary Получение информации о пользователе по ID
// @Description Возвращает данные пользователя по его идентификатору
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Идентификатор пользователя"
// @Success 200 {object} models.User "Данные пользователя"
// @Failure 400 {object} map[string]string "ID не указан"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/users/{id} [get]
func GetUserID(ctx *gin.Context) {
	id := ctx.Param("id")

	id = strings.TrimSpace(id)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	idInt, err := strconv.Atoi(id)
	user, err := repositories.GetUserId(database.DB, idInt)

	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
