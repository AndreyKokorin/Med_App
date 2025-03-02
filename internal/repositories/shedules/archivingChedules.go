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

	loc, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		log.Print("Ошибка загрузки временной зоны:", err)
	}

	// Получаем текущее время в Алматы
	currentTime := time.Now().In(loc)

	// Выполнение запроса
	_, err = db.Exec(query, currentTime)
	if err != nil {
		log.Printf("Failed to update expired schedules: %v", err)
		return err
	}
	slog.Info(currentTime.String())
	slog.Info("Successfully updated expired schedules to archived")

	return nil
}
