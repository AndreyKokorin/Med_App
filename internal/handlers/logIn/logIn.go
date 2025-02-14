package logIn

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/jwt"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
)

func LogIn(ctx *gin.Context) {
	var logInData models.LogInUser

	if err := ctx.ShouldBind(&logInData); err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userId int
	var role string
	var password string
	query := "SELECT id, roles,password FROM users WHERE email=$1"
	err := database.DB.QueryRow(query, logInData.Email).Scan(&userId, &role, &password)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong email or password"})
		return
	}
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(logInData.Password))
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong email and password"})
		return
	}

	token, err := jwt.NewJWT(userId, role)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
