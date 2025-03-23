package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
    DBUser     string
    DBPassword string
    DBName     string
    DBHost     string
    DBPort     string
}

func LoadConfig() (Config, error) {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file") // Не критическая ошибка, если .env нет
    }

    cfg := Config{
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
    }

    return cfg, nil
}