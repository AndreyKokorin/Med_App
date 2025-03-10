package repositories

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"database/sql"
)

func GetMedicalRecordsUserId(userId string) ([]models.FullMedicalRecord, error) {
	var fullMedicalRecords []models.FullMedicalRecord
	query := `
		SELECT 
		m.id,
		m.patient_id,
		p.name AS patient_name,
		m.doctor_id,
		d.name AS doctor_name,
		m.diagnosis,
		m.recomendation,
		m.created_time
	FROM Medical_Records m
	JOIN Users p ON m.patient_id = p.id
	JOIN Users d ON m.doctor_id = d.id
	WHERE m.patient_id = $1;
	`

	rows, err := database.DB.Query(query, userId)
	if err != nil {
		return fullMedicalRecords, err
	}
	defer rows.Close()
	for rows.Next() {
		var fullMedicalRecord models.FullMedicalRecord

		err = rows.Scan(&fullMedicalRecord.ID,
			&fullMedicalRecord.Patient.ID,
			&fullMedicalRecord.Patient.Name,
			&fullMedicalRecord.Doctor.ID,
			&fullMedicalRecord.Doctor.Name,
			&fullMedicalRecord.Diagnosis,
			&fullMedicalRecord.Recommendation,
			&fullMedicalRecord.CreatedTime)

		if err != nil {
			return fullMedicalRecords, err
		}

		fullMedicalRecords = append(fullMedicalRecords, fullMedicalRecord)
	}

	if len(fullMedicalRecords) == 0 {
		return fullMedicalRecords, sql.ErrNoRows
	}

	return fullMedicalRecords, nil
}
