package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *config
	once     sync.Once
)

func GetConfig() *config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(".env file not found, using system environment variables")
		}
		instance = &config{
			Http: httpConfig{
				Port: getEnvAsInt("HTTP_PORT", false, 8000),
			},
			DB: dbConfig{
				URL: getEnv("DB_URL", true, ""),
			},
		}
		fmt.Println("Config loaded from environment variables", instance)
	})
	return instance
}
