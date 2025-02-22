package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllDoctors(ctx *gin.Context) {
	query := "SELECT id, name, age, email,roles FROM users where roles='doctor'"

	rows, err := database.DB.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	var doctors []models.User
	for rows.Next() {
		var doctor models.User
		err := rows.Scan(&doctor.Id, &doctor.Name, &doctor.Age, &doctor.Email, &doctor.Roles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		doctors = append(doctors, doctor)
	}

	ctx.JSON(http.StatusOK, doctors)
}
