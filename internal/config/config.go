package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB_HOST     string `env:"DB_HOST" envDefault:"localhost"`
	DB_PORT     string `env:"DB_PORT" envDefault:"3306"`
	DB_USER     string `env:"DB_USER" envDefault:"root"`
	DB_PASSWORD string `env:"DB_PASSWORD" envDefault:"123456"`
	DB_NAME     string `env:"DB_NAME" envDefault:"golangz"`
	LOCAL_PORT  string `env:"LOCAL_PORT" envDefault:"8080"`
	JWT_SICRET  string `env:"JWT_SICRET" envDefault:""`
}

func LoadConfig() Config {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		LOCAL_PORT:  os.Getenv("LOCAL_PORT"),
		JWT_SICRET:  os.Getenv("JWT_SICRET"),
	}

	return cfg
}
