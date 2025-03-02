package main

import (
	_ "awesomeProject/docs"
	"awesomeProject/internal/Cash"
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	_ "awesomeProject/internal/handlers/logUp"
	_ "awesomeProject/internal/handlers/users"
	"awesomeProject/internal/router"
	"awesomeProject/pkg/helps"
	"github.com/gin-gonic/gin"
	"log"
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

	// Запуск фоновой задачи для обновления актуального расписания в базе
	go helps.ArchiveSchedules()

	// Настройка и запуск HTTP-сервера с Gin
	r := gin.Default()
	router.SetupRouter(r)

	// Запуск сервера на указанном порту
	err = r.Run(cfg.LOCAL_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
