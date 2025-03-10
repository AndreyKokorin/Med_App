package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

// GetAllDoctors
// @Summary Получение всех докторов
// @Description Возвращает список всех пользователей с ролью "doctor"
// @Tags doctors
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.User "Список докторов"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/doctors [get]
func GetAllDoctors(ctx *gin.Context) {
	query := `
			SELECT 
    user_id,
    specialty,
    experience,
    education,
    languages,
    name,
    age,
    email,
    roles,
    gender,
    date_of_birth,
    phone_number,
    address,
    avatar_url
FROM 
    user_profile_view
`

	rows, err := database.DB.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	var doctors []models.FullDoctorProfile
	for rows.Next() {
		var doctor models.FullDoctorProfile
		err := rows.Scan(&doctor.UserID,
			&doctor.Specialty,
			&doctor.Experience,
			&doctor.Education,
			pq.Array(&doctor.Languages), // Используем pq.Array для маппинга массива
			&doctor.Name,
			&doctor.Age,
			&doctor.Email,
			&doctor.Roles,
			&doctor.Gender,
			&doctor.DateOfBirth,
			&doctor.PhoneNumber,
			&doctor.Address,
			&doctor.Avatar_url)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		doctors = append(doctors, doctor)
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "Doctors get successful", "doctors": doctors})
}
