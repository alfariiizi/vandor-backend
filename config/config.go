package config

type Config struct {
	ServerPort int      `json:"server_port"`
	DB         dbConfig `json:"db"`
}

type dbConfig struct {
	Destination string
}

func NewConfig() Config {
	return Config{
		ServerPort: 8080,
		DB: dbConfig{
			Destination: "sqlite.db",
		},
	}
}
