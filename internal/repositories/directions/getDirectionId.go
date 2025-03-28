package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
)

const query = `SELECT id, 
       doctor_id, 
       patient_id, 
       examination_type, 
       executor_doctor_id,
       created_at, 
       status 
FROM direction 
WHERE id = $1`

func GetDirectionId(id int, db *sql.DB) (models.Direction, error) {
	var direction models.Direction
	err := db.QueryRow(query, id).Scan(
		&direction.ID,
		&direction.DoctorID,
		&direction.PatientID,
		&direction.ExaminationType,
		&direction.ExecutorDoctorID,
		&direction.CreatedAt,
		&direction.Status,
	)

	if err != nil {
		return models.Direction{}, err
	}

	return direction, nil
}
