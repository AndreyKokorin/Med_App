package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetFilterUsers получает список пользователей с фильтрацией (доступно врачам и администраторам)
// @Summary Фильтрация пользователей
// @Description Возвращает список пользователей с возможностью фильтрации по возрасту, email и роли
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param age query int false "Возраст пользователя"
// @Param email query string false "Email пользователя"
// @Param role query string false "Роль пользователя"
// @Param limit query int false "Количество записей (по умолчанию 10)"
// @Param offset query int false "Смещение записей (по умолчанию 0)"
// @Success 200 {array} models.User "Список пользователей"
// @Failure 400 {object} map[string]string "Ошибка валидации параметров"
// @Failure 404 {object} map[string]string "Пользователи не найдены"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /users/filter [get]
func GetFilterUsers(ctx *gin.Context) {
	ageQuery := ctx.Query("age")
	emailQuery := ctx.Query("email")
	roleQuery := ctx.Query("role")
	limitQuery := ctx.DefaultQuery("limit", "10")
	offsetQuery := ctx.DefaultQuery("offset", "0")

	query := strings.Builder{}
	query.WriteString("SELECT id, name, age, email, roles FROM users WHERE 1=1")
	args := []interface{}{}
	argID := 1

	if ageQuery != "" {
		query.WriteString(fmt.Sprintf(" AND age=$%d", argID))
		args = append(args, ageQuery)
		argID++
	}

	if emailQuery != "" {
		query.WriteString(fmt.Sprintf(" AND email=$%d", argID))
		args = append(args, emailQuery)
		argID++
	}

	if roleQuery != "" {
		query.WriteString(fmt.Sprintf(" AND roles=$%d", argID))
		args = append(args, roleQuery)
		argID++
	}

	// Преобразуем limit и offset в int
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	query.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1))
	args = append(args, limit, offset)

	rows, err := database.DB.Query(query.String(), args...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Roles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
