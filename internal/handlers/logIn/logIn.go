package logIn

import (
	"awesomeProject/internal/models"
	repositories "awesomeProject/internal/repositories/user"
	"awesomeProject/pkg/helps"
	"awesomeProject/pkg/jwt"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// LogIn обрабатывает авторизацию пользователя
// @Summary Вход пользователя
// @Description Выполняет авторизацию пользователя на основе email и пароля, возвращает токены доступа и обновления
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param login body models.LogInUser true "Данные для входа (email, password)"
// @Success 200 {object} map[string]string "Успешная авторизация с токенами доступа и обновления"
// @Failure 400 {object} map[string]string "Неверный формат запроса"
// @Failure 401 {object} map[string]string "Неверный email или пароль"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера (например, ошибка базы данных или генерации токенов)"
// @Router /login [post]
func LogIn(ctx *gin.Context) {
	loginData, err := getLoginData(ctx)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Invalid request format", err)
		return
	}

	var user models.User
	user, err = repositories.GetUserByEmail(loginData.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helps.RespWithError(ctx, http.StatusUnauthorized, "Wrong email or password", err)
			return
		}

		helps.RespWithError(ctx, http.StatusInternalServerError, "Database access error", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		helps.RespWithError(ctx, http.StatusUnauthorized, "Wrong email or password", err)
		return
	}

	AccessToken, err := jwt.NewJWT(user.Id, user.Roles, time.Now().Add(jwt.AccessTokenTTL))
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to generate authentication token", err)
		return
	}

	RefreshToken, err := jwt.NewJWT(user.Id, user.Roles, time.Now().Add(jwt.RefreshTokenTTL))
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to generate authentication token", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Access-token": AccessToken, "Refresh-token": RefreshToken})
}

func getLoginData(ctx *gin.Context) (models.LogInUser, error) {
	var logInData models.LogInUser

	if err := ctx.ShouldBindJSON(&logInData); err != nil {
		return models.LogInUser{}, err
	}

	return logInData, nil
}
