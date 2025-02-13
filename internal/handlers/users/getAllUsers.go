package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(ctx *gin.Context) {
	var users []models.User

	query := "SELECT id, name, age, email FROM users"

	rows, err := database.DB.Query(query)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		users = append(users, user)
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
