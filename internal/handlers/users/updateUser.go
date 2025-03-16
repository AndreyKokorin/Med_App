package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/helps"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	id, ok := ctx.Get("user_id")

	if !ok {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Не удалось получить ID пользователя из контекста", errors.New("user_id not found in context"))
		return
	}

	var updateUserData models.UpdateUser
	if err := ctx.ShouldBindJSON(&updateUserData); err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Ошибка при разборе JSON данных пользователя", err)
		return
	}

	query, args := buildQueryUpdate(updateUserData, id.(int))

	if len(query) == 0 {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Не удалось сформировать SQL запрос для обновления", errors.New("empty query generated"))
		return
	}

	res, err := database.DB.Exec(query, args...)

	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Ошибка при выполнении SQL запроса в базе данных", err)
		return
	}

	rowssEffected, _ := res.RowsAffected()
	if rowssEffected == 0 {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Ни одна запись пользователя не была обновлена", errors.New("no rows affected"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "update user successfully!"})
}
func buildQueryUpdate(updateUserData models.UpdateUser, userId int) (string, []interface{}) {
	baseQuery := strings.Builder{}
	baseQuery.WriteString("UPDATE users SET ")

	var args []interface{}
	count := 1
	first := true // Флаг для управления запятыми

	if updateUserData.Name != "" {
		if !first {
			baseQuery.WriteString(", ")
		}
		baseQuery.WriteString(fmt.Sprintf("name=$%d", count))
		args = append(args, updateUserData.Name)
		count++
		first = false
	}

	if updateUserData.Email != "" {
		if !first {
			baseQuery.WriteString(", ")
		}
		baseQuery.WriteString(fmt.Sprintf("email=$%d", count))
		args = append(args, updateUserData.Email)
		count++
		first = false
	}

	if updateUserData.PhoneNumber != "" {
		if !first {
			baseQuery.WriteString(", ")
		}
		baseQuery.WriteString(fmt.Sprintf("phone_number=$%d", count))
		args = append(args, updateUserData.PhoneNumber)
		count++
		first = false
	}

	if updateUserData.Age != 0 {
		if !first {
			baseQuery.WriteString(", ")
		}
		baseQuery.WriteString(fmt.Sprintf("age=$%d", count))
		args = append(args, updateUserData.Age)
		count++
		first = false
	}

	if updateUserData.Address != "" {
		if !first {
			baseQuery.WriteString(", ")
		}
		baseQuery.WriteString(fmt.Sprintf("address=$%d", count))
		args = append(args, updateUserData.Address)
		count++
		first = false
	}

	if updateUserData.DateOfBirth != "" {
		if !first {
			baseQuery.WriteString(", ")
		}
		baseQuery.WriteString(fmt.Sprintf("date_of_birth=$%d", count))
		args = append(args, updateUserData.DateOfBirth)
		count++
		first = false
	}

	if first {
		return "", nil
	}

	baseQuery.WriteString(fmt.Sprintf(" WHERE id=$%d", count))
	args = append(args, userId)

	return baseQuery.String(), args
}
