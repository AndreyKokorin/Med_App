package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

const BaseQuery = `SELECT id, 
       doctor_id, 
       patient_id, 
       examination_type, 
       executor_doctor_id,
       created_at, 
       status 
FROM direction 
WHERE 1=1`

func GetFilterDirections(db *sql.DB, filters map[string]interface{}) ([]models.Direction, error) {
	builder := strings.Builder{}
	builder.WriteString(BaseQuery)
	var params []interface{}
	counter := 1

	if doctorId, ok := filters["doctor_id"]; ok && doctorId != "" {
		doctorIdInt, err := strconv.Atoi(doctorId.(string))
		if err != nil {
			return []models.Direction{}, err
		}

		builder.WriteString(fmt.Sprintf(" AND doctor_id = $%d", counter))
		params = append(params, doctorIdInt)
		counter++
	}

	if patientId, ok := filters["patient_id"]; ok && patientId != "" {
		doctorIdInt, err := strconv.Atoi(patientId.(string))
		if err != nil {
			return []models.Direction{}, err
		}

		builder.WriteString(fmt.Sprintf(" AND patient_id = $%d", counter))
		params = append(params, doctorIdInt)
		counter++
	}

	if examinationType, ok := filters["examination_type_id"]; ok && examinationType != "" {
		examinationTypeIdInt, err := strconv.Atoi(examinationType.(string))
		if err != nil {
			return []models.Direction{}, err
		}

		builder.WriteString(fmt.Sprintf(" AND examination_type = $%d", counter))
		params = append(params, examinationTypeIdInt)
		counter++
	}

	if status, ok := filters["status"]; ok && status != "" {
		builder.WriteString(fmt.Sprintf(" AND status = $%d", counter))
		params = append(params, status)
		counter++
	}

	rows, err := db.Query(builder.String(), params...)
	if err != nil {
		return []models.Direction{}, err
	}

	var directions []models.Direction
	for rows.Next() {
		var direction models.Direction

		err := rows.Scan(
			&direction.ID,
			&direction.DoctorID,
			&direction.PatientID,
			&direction.ExaminationType,
			&direction.ExecutorDoctorID,
			&direction.CreatedAt,
			&direction.Status)

		if err != nil {
			return []models.Direction{}, err
		}

		directions = append(directions, direction)
	}
	if err = rows.Err(); err != nil {
		return []models.Direction{}, err
	}

	defer rows.Close()

	slog.Info(builder.String())

	return directions, nil
}
