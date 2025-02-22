package logIn

import (
	"awesomeProject/internal/models"
	repositories "awesomeProject/internal/repositories/user"
	"awesomeProject/pkg/jwt"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
)

func LogIn(ctx *gin.Context) {
	loginData, err := getLoginData(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user, err = repositories.GetUserByEmail(loginData.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong email or password"})
			return
		}

		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong email or password"})
		return
	}

	token, err := jwt.NewJWT(user.Id, user.Roles)
	if err != nil {
		slog.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func getLoginData(ctx *gin.Context) (models.LogInUser, error) {
	var logInData models.LogInUser

	if err := ctx.ShouldBindJSON(&logInData); err != nil {
		return models.LogInUser{}, err
	}

	return logInData, nil
}
