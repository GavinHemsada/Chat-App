package config

import (
	"log"
    "os"

	"github.com/joho/godotenv"
)

type Config struct {
    ServerPort     string
    DBHost         string
    DBPort         string
    DBUser         string
    DBPassword     string
    DBName         string
    RedisAddr      string
    RedisPassword  string
    JWTSecret      string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
    return &Config{
        ServerPort:    os.Getenv("SERVER_PORT"),
        DBHost:        os.Getenv("DB_HOST"),
        DBPort:        os.Getenv("DB_PORT"),
        DBUser:        os.Getenv("DB_USER"),
        DBPassword:    os.Getenv("DB_PASSWORD"),
        DBName:        os.Getenv("DB_NAME"),
        RedisAddr:     os.Getenv("REDIS_ADDR"),
        RedisPassword: os.Getenv("REDIS_PASSWORD"),
        JWTSecret:     os.Getenv("JWT_SECRET"),
    }
}