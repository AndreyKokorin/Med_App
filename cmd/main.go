package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	repositories "awesomeProject/internal/repositories/shedules"
	"awesomeProject/internal/router"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	// Загрузка конфигурации и инициализация базы данных
	cfg := config.LoadConfig()
	database.DbInit(cfg)

	// Запуск фоновой задачи для обновления записей в базе данных каждые 5 секунд
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			// Вызов функции архивации
			if err := repositories.ArchiveExpiredSchedules(database.DB); err != nil {
				log.Printf("Error archiving expired schedules: %v", err)
			}
		}
	}()

	// Настройка и запуск HTTP-сервера с Gin
	r := gin.Default()
	router.SetupRouter(r)

	// Запуск сервера на указанном порту
	err := r.Run(cfg.LOCAL_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
