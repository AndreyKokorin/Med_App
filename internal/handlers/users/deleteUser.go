package users

import (
	"awesomeProject/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteUser Удаление пользователя (только для администраторов)
// @Summary Удаление пользователя
// @Description Удаляет пользователя по его ID. Доступно только администраторам.
// @Tags users
// @Security ApiKeyAuth
// @Param id path string true "ID пользователя"
// @Success 200 {object} map[string]string "Пользователь удален"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 403 {object} map[string]string "Недостаточно прав"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	query := "DELETE FROM users WHERE id=$1"
	res, err := database.DB.Exec(query, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении количества удаленных строк"})
		return
	}
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"deleted user": userId})
}
