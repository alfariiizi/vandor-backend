package config

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *config
	once     sync.Once
)

func GetConfig() *config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println(".env file not found, assuming environment variables are set by Docker")
		}
		appUrl, err := url.Parse(getEnv("APP_URL", true, ""))
		if err != nil {
			log.Fatal("Invalid APP_URL format, must be a valid URL")
		}
		port, err := strconv.Atoi(appUrl.Port())
		if err != nil {
			log.Println("Using default port 8000, as APP_URL does not specify a port")
			port = 8000
		}
		redisOption, err := parseRedisURL(getEnv("REDIS_URL", true, ""))
		if err != nil {
			log.Fatal("Invalid REDIS_URL format, must be a valid URL")
		}
		instance = &config{
			Superadmin: superadminConfig{
				Name:     getEnv("SUPERADMIN_NAME", true, ""),
				Email:    getEnv("SUPERADMIN_EMAIL", true, ""),
				Password: getEnv("SUPERADMIN_PASSWORD", true, ""),
			},
			App: appInfoConfig{
				Name:              getEnv("APP_NAME", false, "Go Service"),
				SignatureResponse: getEnv("APP_SIGNATURE_RESPONSE", false, "rizalalfarizi.com"),
				Version:           getEnv("APP_VERSION", false, "1.0.0"),
			},
			Http: httpConfig{
				AppURL: appUrl.String(),
				Port:   port,
			},
			DB: dbConfig{
				Driver: getEnv("DB_DRIVER", true, ""),
				URL:    getEnv("DB_URL", true, ""),
			},
			Docs: docsConfig{
				Username: getEnv("DOCS_USERNAME", true, ""),
				Password: getEnv("DOCS_PASSWORD", true, ""),
			},
			Redis: *redisOption,
			Auth: authConfig{
				SecretKey: getEnv("AUTH_SECRET_KEY", true, ""),
			},
			Email: emailConfig{
				Url:    getEnv("EMAIL_URL", true, ""),
				Domain: getEnv("EMAIL_DOMAIN", true, ""),
				Secret: getEnv("EMAIL_SECRET", true, ""),
			},
		}
		fmt.Println("Config loaded from environment variables", instance)
	})
	return instance
}
