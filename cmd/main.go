package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers/register"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.DbInit(cfg)

	r := gin.Default()

	r.POST("/user/register", register.LogUpUser)

	err := r.Run(cfg.LOCAL_PORT)
	if err != nil {
		panic(err)
	}
}
