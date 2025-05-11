package config

type config struct {
	Http httpConfig
	DB   dbConfig
}

type httpConfig struct {
	Port int
}

type dbConfig struct {
	URL string
}
