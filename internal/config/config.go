package config

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/alfariiizi/vandor/internal/pkg/validator"
	"github.com/fsnotify/fsnotify"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

var (
	instance *config
	once     sync.Once
	validate = validator.New()
)

func GetConfig() *config {
	once.Do(func() {
		v := viper.New()

		// setDefaultValues(v)

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
		secretViper.AddConfigPath("/app/config") // K8s Secret mount path
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
			load(v)
		})
		secretViper.WatchConfig()
		secretViper.OnConfigChange(func(e fsnotify.Event) {
			log.Println("Secret config changed:", e.Name)
			v.MergeConfigMap(secretViper.AllSettings())
			load(v)
		})

		load(v)

		fmt.Println("Config loaded successfully")
	})
	return instance
}

func load(v *viper.Viper) {
	newCfg := &config{}

	defaults.SetDefaults(newCfg)

	if err := v.Unmarshal(newCfg); err != nil {
		log.Println("Error decoding config:", err)
		return
	}
	log.Println("Config loaded from Viper:", newCfg)
	if err := validate.Validate(newCfg); err != nil {
		log.Println("Invalid config:", err)
		panic(fmt.Sprintf("Config validation failed: %v", err))
	}

	// all process before assign to instance
	redisProcess(newCfg)

	instance = newCfg
	log.Println("Config loaded successfully")
}

func redisProcess(cfg *config) {
	cfg.Redis.Addr = fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
}
