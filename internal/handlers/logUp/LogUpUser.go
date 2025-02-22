package logUp

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/internal/repositories/user"
	"awesomeProject/pkg/hash"
	"awesomeProject/pkg/validate"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const (
	admin_role   = "admin"
	defoult      = "user"
	doctor_role  = "doctor"
	admin_token  = "admin_token"
	doctor_token = "doctor_token"
)

func LogUpUser(ctx *gin.Context) {
	var logUpData models.LogUpUser
	if err := ctx.ShouldBindJSON(&logUpData); err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.ValidAndTrim(&logUpData); err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//задаем роль
	logUpData.Roles = determinateRole(logUpData.RoleToken)

	//Хэшируем пароль
	hashPassword, err := hash.PasswordHash(logUpData.Password)
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Проверяем есть ли такой пользователь в бд
	exists, err := repositories.UserExists(logUpData.Email, database.DB)
	if err != nil {
		slog.Error("Database error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if exists {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	//Добавляем пользователя в бд
	err = repositories.NewUser(logUpData, database.DB, hashPassword)
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
