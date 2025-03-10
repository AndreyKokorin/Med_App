package doctorRep

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"errors"
	"github.com/lib/pq"
)

func GetFullDoctorProfile(userId int) (models.FullDoctorProfile, error) {
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
WHERE 
    user_id = $1;
`

	var doctorProfile models.FullDoctorProfile
	err := database.DB.QueryRow(query, userId).Scan(
		&doctorProfile.UserID,
		&doctorProfile.Specialty,
		&doctorProfile.Experience,
		&doctorProfile.Education,
		pq.Array(&doctorProfile.Languages), // Используем pq.Array для маппинга массива
		&doctorProfile.Name,
		&doctorProfile.Age,
		&doctorProfile.Email,
		&doctorProfile.Roles,
		&doctorProfile.Gender,
		&doctorProfile.DateOfBirth,
		&doctorProfile.PhoneNumber,
		&doctorProfile.Address,
		&doctorProfile.Avatar_url,
	)

	if err != nil {
		return models.FullDoctorProfile{}, err
	}

	if doctorProfile.Roles != "doctor" {
		return models.FullDoctorProfile{}, errors.New("user isnt doctor")
	}

	return doctorProfile, nil
}
