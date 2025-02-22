package repositories

import (
	"database/sql"
	"log"
	"log/slog"
	"time"
)

func ArchiveExpiredSchedules(db *sql.DB) error {
	// SQL-запрос для изменения статуса на 'archived' для записей, у которых время прошло
	query := `
		UPDATE schedules
		SET status = 'archived'
		WHERE end_time < $1 AND status = 'active'
	`

	// Выполнение запроса
	_, err := db.Exec(query, time.Now())
	if err != nil {
		log.Printf("Failed to update expired schedules: %v", err)
		return err
	}

	slog.Info("Successfully updated expired schedules to archived")

	return nil
}
