package helps

import (
	"awesomeProject/internal/database"
	repositories "awesomeProject/internal/repositories/shedules"
	"log"
	"time"
)

func ArchiveSchedules() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		// Вызов функции архивации
		if err := repositories.ArchiveExpiredSchedules(database.DB); err != nil {
			log.Printf("Error archiving expired schedules: %v", err)
		}
	}
}
