package config

import "sync"

type Config struct {
	ServerPort int      `json:"server_port"`
	DB         dbConfig `json:"db"`
}

type dbConfig struct {
	Destination string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			ServerPort: 8080,
			DB: dbConfig{
				Destination: "sqlite.db",
			},
		}
	})
	return instance
}
