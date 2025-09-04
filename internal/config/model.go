package config

type config struct {
	App   appInfoConfig
	HTTP  httpConfig `mapstructure:"http"`
	DB    dbConfig   `mapstructure:"database"`
	Docs  docsConfig
	Jobs  jobsConfig
	Auth  authConfig
	Redis redisConfig `mapstructure:"redis"`
	Log   logConfig   `mapstructure:"log"`
}

type appInfoConfig struct {
	Name              string `mapstructure:"name" default:"Vandor"`
	SignatureResponse string `mapstructure:"signature_response" default:"vandor.com"`
	Version           string `mapstructure:"version" default:"1.0.0"`
}

type httpConfig struct {
	URL  string `mapstructure:"url" validate:"required,url"`
	Port int    `mapstructure:"port" validate:"required,min=1,max=65535"`
}

type dbConfig struct {
	Driver string `mapstructure:"driver" validate:"required"`
	URL    string `mapstructure:"url" validate:"required"`
}

type authConfig struct {
	SecretKey              string `mapstructure:"secret_key" validate:"required"`
	TokenDurationInMinutes int    `mapstructure:"token_duration_in_minutes" validate:"required,min=1"`
	SessionDurationInDays  int    `mapstructure:"session_duration_in_days" validate:"required,min=1"`
}

type jobsConfig struct {
	QueuePrefix      string `mapstructure:"queue_prefix"`
	HTTPHeaderSecret string `mapstructure:"http_header_secret" validate:"required"`
	Concurrency      int    `mapstructure:"concurrency" validate:"required,min=1"`
}

type redisConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required,min=1,max=65535"`
	Addr     string `mapstructure:"-" default:""`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db" validate:"min=0,max=15" default:"0"`
}

type docsConfig struct {
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

type logConfig struct {
	Level         string `mapstructure:"level" validate:"required"`
	EnableConsole bool   `mapstructure:"enable_console" default:"true"`
}
