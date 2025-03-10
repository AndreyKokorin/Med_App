package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllUsers
// @Summary Получение всех пользователей
// @Description Возвращает список всех пользователей в системе
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} map[string][]models.User "Список пользователей"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/users [get]
func GetAllUsers(ctx *gin.Context) {
	var users []models.User

	query := "SELECT id, email, age, name, roles, date_of_birth, phone_number, address, gender, avatar_url FROM users"

	rows, err := database.DB.Query(query)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Email, &user.Age, &user.Name, &user.Roles, &user.DateOfBirth, &user.PhoneNumber, &user.Address, &user.Gender, &user.Avatar_url)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		users = append(users, user)
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
