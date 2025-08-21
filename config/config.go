package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl      string
	RedisURL   string
	ServerPort string
	NumWorkers int
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		DBUrl:      os.Getenv("DB_URL"),
		RedisURL:   os.Getenv("REDIS_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),
		NumWorkers: func() int {
			num, _ := os.LookupEnv("NUM_WORKERS")
			if num == "" {
				num = "4"
			}
			val, _ := strconv.Atoi(num)
			return val
		}(),
	}
}
