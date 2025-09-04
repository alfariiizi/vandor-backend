package logger

import (
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/alfariiizi/vandor/internal/enum"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	EnableConsole bool
	LogDir        string
	ServiceName   string
	LogLevel      enum.LogLevel
}

var (
	once     sync.Once
	instance *zerolog.Logger
)

// parseLogLevel converts string to zerolog.Level
func parseLogLevel(level enum.LogLevel) zerolog.Level {
	switch level {
	case enum.LogLevelDebug:
		return zerolog.DebugLevel
	case enum.LogLevelInfo:
		return zerolog.InfoLevel
	case enum.LogLevelWarning, enum.LogLevelWarn:
		return zerolog.WarnLevel
	case enum.LogLevelError:
		return zerolog.ErrorLevel
	case enum.LogLevelFatal:
		return zerolog.FatalLevel
	case enum.LogLevelPanic:
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

// Init initializes the singleton logger (only once).
func Init() *zerolog.Logger {
	cfg := config.GetConfig()

	once.Do(func() {
		// Ensure log directory exists
		dir := filepath.Join("storage", "logs")
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic("failed to create log dir: " + err.Error())
		}

		fileName := filepath.Join(dir, time.Now().Format("2006-01-02")+".log")

		fileWriter := &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    100,
			MaxBackups: 30,
			MaxAge:     30,
			Compress:   false,
		}

		var writers []io.Writer
		if cfg.Log.EnableConsole {
			consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
			writers = append(writers, consoleWriter)
		}
		writers = append(writers, fileWriter)

		multi := io.MultiWriter(writers...)

		// Set level
		logLevel, err := enum.ParseLogLevel(cfg.Log.Level)
		if err != nil {
			panic("invalid log level: " + err.Error())
		}
		level := parseLogLevel(logLevel)
		zerolog.SetGlobalLevel(level)

		// Create logger
		logger := zerolog.New(multi).With().
			Timestamp().
			Logger()
		instance = &logger

		log.Logger = *instance
	})
	return instance
}

// Get returns the singleton logger (after Init is called).
func Get() *zerolog.Logger {
	if instance == nil {
		panic("logger not initialized, call logger.Init first")
	}
	return instance
}
