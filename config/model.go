package config

import "github.com/redis/go-redis/v9"

type config struct {
	App        appInfoConfig
	Superadmin superadminConfig
	HTTP       httpConfig
	DB         dbConfig
	Docs       docsConfig
	Redis      redis.Options
	Worker     workerConfig
	Auth       authConfig
	Email      emailConfig
}

type appInfoConfig struct {
	Name              string `mapstructure:"name"`
	SignatureResponse string `mapstructure:"signature_response"`
	Version           string `mapstructure:"version"`
}

type httpConfig struct {
	AppURL string `mapstructure:"app_url" validate:"required,url"`
	Port   int    `mapstructure:"port" validate:"required,min=1,max=65535"`
}

type dbConfig struct {
	Driver string `mapstructure:"driver" validate:"required"`
	URL    string `mapstructure:"url" validate:"required"`
}

type authConfig struct {
	SecretKey string `mapstructure:"secret_key" validate:"required"`
}

type superadminConfig struct {
	Name     string `mapstructure:"name" validate:"required"`
	Email    string `mapstructure:"email" validate:"required,email"`
	Password string `mapstructure:"password" validate:"required,min=8"`
}

type workerConfig struct {
	Concurrency int `mapstructure:"concurrency" validate:"required,min=1"`
}

type docsConfig struct {
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

type emailConfig struct {
	URL    string `mapstructure:"url" validate:"required,url"`
	Domain string `mapstructure:"domain" validate:"required"`
	Secret string `mapstructure:"secret" validate:"required"`
}
