package doctorRep

import (
	"awesomeProject/internal/models"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"strings"
)

// GetFilteredDoctors возвращает отфильтрованный список докторов
func GetFilteredDoctors(db *sql.DB, filters map[string]interface{}) ([]models.FullDoctorProfile, error) {
	// Базовый запрос
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

	// Динамически добавляем условия фильтрации
	var conditions []string
	var args []interface{}
	argCounter := 1

	for key, value := range filters {
		switch key {
		case "specialty":
			conditions = append(conditions, fmt.Sprintf("specialty = $%d", argCounter))
			args = append(args, value)
			argCounter++
		case "experience":
			conditions = append(conditions, fmt.Sprintf("experience >= $%d", argCounter))
			args = append(args, value)
			argCounter++
		case "languages":
			conditions = append(conditions, fmt.Sprintf("$%d = ANY(languages)", argCounter))
			args = append(args, value)
			argCounter++
		case "gender":
			conditions = append(conditions, fmt.Sprintf("gender = $%d", argCounter))
			args = append(args, value)
			argCounter++
		case "min_age":
			conditions = append(conditions, fmt.Sprintf("age >= $%d", argCounter))
			args = append(args, value)
			argCounter++
		case "max_age":
			conditions = append(conditions, fmt.Sprintf("age <= $%d", argCounter))
			args = append(args, value)
			argCounter++
		}
	}

	// Добавляем условия в запрос, если они есть
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Выполняем запрос
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var doctors []models.FullDoctorProfile
	for rows.Next() {
		var doctor models.FullDoctorProfile
		var languagesStr string

		err := rows.Scan(
			&doctor.UserID,
			&doctor.Specialty,
			&doctor.Experience,
			&doctor.Education,
			pq.Array(&doctor.Languages),
			&doctor.Name,
			&doctor.Age,
			&doctor.Email,
			&doctor.Roles,
			&doctor.Gender,
			&doctor.DateOfBirth,
			&doctor.PhoneNumber,
			&doctor.Address,
			&doctor.Avatar_url,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Преобразуем languages из строки в массив
		doctor.Languages = strings.Split(languagesStr, ",")
		doctors = append(doctors, doctor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return doctors, nil
}
