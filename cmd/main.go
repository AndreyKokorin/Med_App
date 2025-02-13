package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	"awesomeProject/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.DbInit(cfg)

	r := gin.Default()
	router.SetupRouter(r)
	err := r.Run(cfg.LOCAL_PORT)

	if err != nil {
		panic(err)
	}
}
