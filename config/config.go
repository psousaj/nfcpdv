package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl      string
	RedisURL   string
	ServerPort string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		DBUrl:      os.Getenv("DB_URL"),
		RedisURL:   os.Getenv("REDIS_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
