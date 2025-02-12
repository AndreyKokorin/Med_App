package database

import (
	"awesomeProject/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func DbInit(cfg config.Config) {
	strCon := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT)
	var err error
	DB, err = sql.Open("postgres", strCon)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to database %s", cfg.DB_NAME)
}
