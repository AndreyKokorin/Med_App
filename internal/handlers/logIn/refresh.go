package logIn

import (
	"awesomeProject/pkg/helps"
	"awesomeProject/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RefreshToken struct {
	Refresh string `json:"refresh"`
}

// Refresh обновляет токен доступа
// @Summary Обновление access-токена
// @Description Принимает refresh-токен и возвращает новый access-токен
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param input body RefreshToken true "Refresh-токен"
// @Success 200 {object} map[string]string "Новый access-токен"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 401 {object} map[string]string "Недействительный или истекший токен"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /refresh [post]
func Refresh(ctx *gin.Context) {
	var refresh RefreshToken

	if err := ctx.ShouldBindJSON(&refresh); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if refresh.Refresh == "" {
		helps.RespWithError(ctx, http.StatusBadRequest, "Refresh token is required", errors.New("Refresh token is required"))
		return
	}

	dataFromToken, err := jwt.ParseJWT(refresh.Refresh)
	if err != nil {
		helps.RespWithError(ctx, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	userID, ok := dataFromToken["user_id"].(int)
	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "Invalid token payload ID", errors.New("user_id is not a valid number"))
		return
	}

	role, ok := dataFromToken["role"].(string)
	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "Invalid token payload ROLE", errors.New("role is not a valid string"))
		return
	}

	newAccessToken, err := jwt.NewJWT(userID, role, time.Now().Add(jwt.AccessTokenTTL))
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to generate new access token", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}
