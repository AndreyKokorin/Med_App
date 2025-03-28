package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
	"log/slog"
)

func CreateResultExamination(directId int, doctor_id int, urlFile string, db *sql.DB) (models.DirectResult, error) {
	query := `SELECT doctor_id FROM direction WHERE id=$1`

	var doctorIdDirection int
	err := db.QueryRow(query, directId).Scan(&doctorIdDirection)
	if err != nil {
		slog.Info("doctorIdDirection error")
		return models.DirectResult{}, err
	}

	if doctorIdDirection != doctor_id {
		return models.DirectResult{}, errors.New("You cannot upload results for this direction.")
	}

	query = `INSERT INTO examination_results(direction_id, doctor_id, file_path)
				VALUES ($1, $2, $3)
				RETURNING id, direction_id, doctor_id, file_path, created_at`

	var result models.DirectResult
	err = db.QueryRow(query, directId, doctor_id, urlFile).Scan(&result.Id, &result.DirectionId, &result.DoctorID, &result.FilePath, &result.CreatedAt)
	if err != nil {
		slog.Info("result error")
		return models.DirectResult{}, err
	}

	return result, nil
}
