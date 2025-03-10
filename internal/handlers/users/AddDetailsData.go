package users

import (
	"awesomeProject/internal/models"
	repositories "awesomeProject/internal/repositories/user"
	"awesomeProject/pkg/helps"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddDetailsData
// @Summary Добавление дополнительных данных пользователя
// @Description Обновляет дополнительные данные текущего пользователя
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param details body models.UserDetails true "Дополнительные данные пользователя"
// @Success 200 {object} map[string]string "Детали успешно обновлены"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/users/details [post]
func AddDetailsData(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		helps.RespWithError(ctx, http.StatusUnauthorized, "user_id not found", nil)
		return
	}

	var detailsData models.UserDetails
	if err := ctx.ShouldBindJSON(&detailsData); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "invalide json format", err)
		return
	}

	err := repositories.UpdateUserDetails(detailsData, userID.(int))
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error updating user details", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})

}
