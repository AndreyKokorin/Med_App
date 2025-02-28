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

// LogUpUser обрабатывает регистрацию нового пользователя
// @Summary Регистрация нового пользователя
// @Description Создает нового пользователя в системе на основе предоставленных данных (email, пароль, роль)
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param user body models.LogUpUser true "Данные для регистрации пользователя (email, password, RoleToken)"
// @Success 201 {object} nil "Успешная регистрация (пустой ответ)"
// @Failure 400 {object} map[string]string "Неверный формат запроса или данные"
// @Failure 409 {object} map[string]string "Пользователь с таким email уже зарегистрирован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера (например, ошибка базы данных или хеширования пароля)"
// @Router /logup [post]
func LogUpUser(ctx *gin.Context) {
	var logUpData models.LogUpUser

	// Decode JSON
	if err := ctx.ShouldBindJSON(&logUpData); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Invalid request format", err)
		return
	}

	// Validate input data
	if err := validate.ValidAndTrim(&logUpData); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	// Determine user role
	logUpData.Roles = determinateRole(logUpData.RoleToken)

	// Hash password
	hashPassword, err := hash.PasswordHash(logUpData.Password)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Password processing error", err)
		return
	}

	// Check if user already exists
	exists, err := repositories.UserExists(logUpData.Email, database.DB)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Database access error", err)
		return
	}
	if exists {
		helps.RespWithError(ctx, http.StatusConflict, "A user with this email is already registered", errors.New("email already in use"))
		return
	}

	// Create new user
	err = repositories.NewUser(logUpData, database.DB, hashPassword)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	ctx.Status(http.StatusCreated)
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
