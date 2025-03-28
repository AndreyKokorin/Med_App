package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
)

func GetResultById(id int, db *sql.DB) (models.DirectResult, error) {
	query := `SELECT 
    id, 
    direction_id, 
    file_path, 
    created_at 
FROM 
    examination_results 
WHERE 
    id = $1`

	var result models.DirectResult
	err := db.QueryRow(query, id).Scan(&result.Id, &result.DirectionId, &result.FilePath, &result.CreatedAt)

	if err != nil {
		return models.DirectResult{}, err
	}

	return result, nil
}
