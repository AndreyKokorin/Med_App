package users

import (
	"awesomeProject/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteUser удаляет пользователя по ID
// @Summary Удаление пользователя
// @Description Удаляет пользователя из базы данных по указанному ID (доступно только для роли admin)
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя для удаления"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string "Успешное удаление пользователя"
// @Failure 401 {object} map[string]string "Доступ запрещён: отсутствует или неверный токен авторизации"
// @Failure 403 {object} map[string]string "Доступ запрещён: недостаточно прав (требуется роль admin)"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера (например, ошибка базы данных)"
// @Router /admin/delete/user/:id [delete]
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
