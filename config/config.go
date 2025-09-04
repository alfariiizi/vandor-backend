package config

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/alfariiizi/vandor/pkg/validator"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	instance *config
	once     sync.Once
	validate = validator.New()
)

func LoadConfig() *config {
	once.Do(func() {
		v := viper.New()

		setDefaultValues(v)

		// Allow env var overrides: APP_NAME -> app.name
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// 1️⃣ Load non-sensitive defaults (ConfigMap in K8s)
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("./config")    // local
		v.AddConfigPath("/app/config") // Docker/K8s mount
		if err := v.MergeInConfig(); err != nil {
			log.Println("config.yaml not found, continuing...")
		}

		// 2️⃣ Load sensitive values (Secret in K8s)
		secretViper := viper.New()
		secretViper.SetConfigName("secret")
		secretViper.SetConfigType("yaml")
		secretViper.AddConfigPath("./config")
		secretViper.AddConfigPath("/app/secret") // K8s Secret mount path
		if err := secretViper.MergeInConfig(); err != nil {
			log.Println("secret.yaml not found, assuming env vars will provide secrets")
		}
		v.MergeConfigMap(secretViper.AllSettings())

		// 3️⃣ Load from .env (local dev convenience)
		envViper := viper.New()
		envViper.SetConfigFile(".env")
		if err := envViper.MergeInConfig(); err == nil {
			v.MergeConfigMap(envViper.AllSettings())
		}

		// 🔄 Hot reload for both config.yaml and secret.yaml
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			log.Println("Config changed:", e.Name)
			reload(v)
		})
		secretViper.WatchConfig()
		secretViper.OnConfigChange(func(e fsnotify.Event) {
			log.Println("Secret config changed:", e.Name)
			v.MergeConfigMap(secretViper.AllSettings())
			reload(v)
		})

		reload(v)

		fmt.Println("Config loaded successfully")
	})
	return instance
}

func setDefaultValues(v *viper.Viper) {
	// --- Default Values (safe for local dev) ---
	v.SetDefault("app.name", "Vandor Service")
	v.SetDefault("app.signature_response", "example.com")
	v.SetDefault("app.version", "1.0.0")

	v.SetDefault("http.app_url", "http://localhost:8000")
	v.SetDefault("http.port", 8000)

	v.SetDefault("db.driver", "postgres")
	v.SetDefault("db.url", "postgres://user:pass@localhost:5432/db")

	v.SetDefault("docs.username", "docsuser")
	v.SetDefault("docs.password", "docspass")

	v.SetDefault("redis.host", "localhost")
	v.SetDefault("redis.port", "6379")
	v.SetDefault("redis.username", "")
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)

	v.SetDefault("worker.concurrency", 10)

	v.SetDefault("auth.secret_key", "changeme")
	v.SetDefault("email.url", "https://api.emailservice.com")
	v.SetDefault("email.domain", "example.com")
	v.SetDefault("email.secret", "emailsecret")

	v.SetDefault("superadmin.name", "Admin")
	v.SetDefault("superadmin.email", "admin@example.com")
	v.SetDefault("superadmin.password", "password")
}

func reload(v *viper.Viper) {
	newCfg := &config{}
	if err := v.Unmarshal(newCfg); err != nil {
		log.Println("Error decoding config:", err)
		return
	}
	if err := validate.Validate(newCfg); err != nil {
		log.Println("Invalid config:", err)
		return
	}
	instance = newCfg
	log.Println("Config loaded successfully")
}
