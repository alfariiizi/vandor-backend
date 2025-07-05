package config

import "github.com/redis/go-redis/v9"

type config struct {
	App        appInfoConfig
	Superadmin superadminConfig
	Http       httpConfig
	DB         dbConfig
	Docs       docsConfig
	Redis      redis.Options
	Auth       authConfig
	Email      emailConfig
}

type appInfoConfig struct {
	Name              string
	SignatureResponse string
	Version           string
}

type httpConfig struct {
	AppURL string
	Port   int
}

type dbConfig struct {
	Driver string
	URL    string
}

type authConfig struct {
	SecretKey string
}

type superadminConfig struct {
	Name     string
	Email    string
	Password string
}

type docsConfig struct {
	Username string
	Password string
}

type emailConfig struct {
	Url    string
	Domain string
	Secret string
}
