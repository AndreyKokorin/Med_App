package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllUsers возвращает список всех пользователей
// @Summary Получение списка всех пользователей
// @Description Возвращает список всех пользователей из базы данных (доступно только для роли admin)
// @Tags Пользователи
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string][]models.User "Список пользователей"
// @Failure 401 {object} map[string]string "Доступ запрещён: отсутствует или неверный токен авторизации"
// @Failure 403 {object} map[string]string "Доступ запрещён: недостаточно прав (требуется роль admin)"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера (например, ошибка базы данных)"
// @Router /admin/users [get]
func GetAllUsers(ctx *gin.Context) {
	var users []models.User

	query := "SELECT id, name, age, email FROM users"

	rows, err := database.DB.Query(query)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		users = append(users, user)
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
