package logUp

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/internal/repositories/user"
	"awesomeProject/pkg/hash"
	"awesomeProject/pkg/helps"
	"awesomeProject/pkg/validate"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	admin_role   = "admin"
	defoult      = "user"
	doctor_role  = "doctor"
	admin_token  = "admin_token"
	doctor_token = "doctor_token"
)

// LogUpUser регистрирует нового пользователя
// @Summary Регистрация пользователя
// @Description Создает нового пользователя в системе
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param input body models.LogUpUser true "Данные для регистрации"
// @Success 201 {object} map[string]interface{} "Пользователь успешно создан"
// @Failure 400 {object} map[string]interface{} "Некорректный формат запроса"
// @Failure 409 {object} map[string]interface{} "Пользователь с таким email уже существует"
// @Failure 500 {object} map[string]interface{} "Ошибка на стороне сервера"
// @Router /auth/register [post]
func LogUpUser(ctx *gin.Context) {
	var logUpData models.LogUpUser

	// Decode JSON
	if err := ctx.ShouldBindJSON(&logUpData); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Invalid request format", err)
		return
	}

	if err := validate.ValidAndTrim(&logUpData); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	logUpData.Roles = determinateRole(logUpData.RoleToken)

	// Hash password
	hashPassword, err := hash.PasswordHash(logUpData.Password)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Password processing error", err)
		return
	}

	exists, err := repositories.UserExists(logUpData.Email, database.DB)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Database access error", err)
		return
	}
	if exists {
		helps.RespWithError(ctx, http.StatusConflict, "A user with this email is already registered", errors.New("email already in use"))
		return
	}

	err = repositories.NewUser(logUpData, database.DB, hashPassword)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	ctx.Status(http.StatusCreated)

	ctx.JSON(http.StatusCreated, gin.H{"massage": "user created", "data": logUpData})
}

func determinateRole(TokenRole string) string {
	if TokenRole == admin_token {
		return admin_role
	}

	if TokenRole == doctor_token {
		return doctor_role
	}

	return defoult
}
