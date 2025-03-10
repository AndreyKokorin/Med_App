package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/validate"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

// UpdateUser
// @Summary Обновление информации о пользователе
// @Description Позволяет обновить данные пользователя (имя, email, возраст)
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор пользователя"
// @Param user body models.UpdateUser true "Обновляемые данные пользователя"
// @Success 200 {object} map[string]models.User "Обновленный пользователь"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/users/{id} [put]
func UpdateUser(ctx *gin.Context) {
	var userid, _ = strconv.Atoi(ctx.Param("id"))

	var updateData models.UpdateUser
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("User request Failed")
		return
	}

	err := validate.ValidAndTrim(&updateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("User valid Failed")
		return
	}

	var userWithNewData models.User
	query := "UPDATE users SET name = COALESCE(NULLIF($1, ''), name), email = COALESCE(NULLIF($2, ''), email), age = COALESCE(NULLIF($3, 0), age) WHERE id = $4 RETURNING id, name, email, age;"
	err = database.DB.QueryRow(query, updateData.Name, updateData.Email, updateData.Age, userid).Scan(&userWithNewData.Id, &userWithNewData.Name, &userWithNewData.Email, &userWithNewData.Age)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error("User sql Failed")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userWithNewData})
}
