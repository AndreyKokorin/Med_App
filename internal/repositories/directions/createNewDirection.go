package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
	"fmt"
)

func CreateNewDirection(db *sql.DB, direction models.Direction, doctorId int) (models.Direction, error) {
	query := `INSERT INTO direction(doctor_id, patient_id, examination_type,executor_doctor_id) 
		VALUES ($1, $2, $3,$4) 
		RETURNING id, doctor_id, patient_id, examination_type, created_at,status,executor_doctor_id;`

	var NewDirection models.Direction
	err := db.QueryRow(query,
		doctorId,
		direction.PatientID,
		direction.ExaminationType,
		direction.ExecutorDoctorID,
	).Scan(&NewDirection.ID,
		&NewDirection.DoctorID,
		&NewDirection.PatientID,
		&NewDirection.ExaminationType,
		&NewDirection.CreatedAt,
		&NewDirection.Status,
		&NewDirection.ExecutorDoctorID,
	)

	if err != nil {
		return models.Direction{}, fmt.Errorf("ошибка при создании направления: %w", err)
	}

	return NewDirection, nil
}
