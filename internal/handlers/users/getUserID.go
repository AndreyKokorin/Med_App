package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetUserID(ctx *gin.Context) {
	id := ctx.Param("id")

	id = strings.TrimSpace(id)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var user models.User
	query := "SELECT id, name, age, email,roles from users where id=$1"
	err := database.DB.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Roles)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
