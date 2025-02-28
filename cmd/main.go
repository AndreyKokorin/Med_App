package main

import (
	_ "awesomeProject/docs"
	"awesomeProject/internal/Cash"
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	_ "awesomeProject/internal/handlers/logUp"
	_ "awesomeProject/internal/handlers/users"
	repositories "awesomeProject/internal/repositories/shedules"
	"awesomeProject/internal/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"time"
)

// @title API для медицинского приложения
// @version 1.0
// @description API для медицинского приложения
// @host localhost:8088
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer токен авторизации (например, "Bearer <token>")
func main() {

	// Загрузка конфигурации и инициализация базы данных
	cfg := config.LoadConfig()
	//подключаем кэш

	err := Cash.InitRedis()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	//доключаем бд
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера на указанном порту
	err = r.Run(cfg.LOCAL_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
