package users

import (
	"awesomeProject/internal/Cash"
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func ChangePasswordSendEmail(ctx *gin.Context) {
	var emailAddress models.To

	if err := ctx.ShouldBindJSON(&emailAddress); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	row := database.DB.QueryRow("SELECT email FROM users WHERE email=$1", emailAddress.Email)
	var foundEmail string
	if err := row.Scan(&foundEmail); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": "Ошибка при восстановлении пароля"})

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": "Ошибка базы данных"})
		}
		return
	}

	secretCode := fmt.Sprintf("%04d", rand.Intn(10000))

	apiKey := os.Getenv("EMAIL_SENDER_API_KEY")
	if apiKey == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "Email API key не найден"})
		return
	}
	client := resend.NewClient(os.Getenv("EMAIL_SENDER_API_KEY"))

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{emailAddress.Email},
		Subject: "Смена пароля",
		Html:    fmt.Sprintf("Твой код: %s", secretCode),
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	err = Cash.RedisClient.Set(ctx.Request.Context(), emailAddress.Email, secretCode, time.Minute*5).Err()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sentId": sent.Id})
}
