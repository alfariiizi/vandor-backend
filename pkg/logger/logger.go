package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/alfariiizi/vandor/internal/enum"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config holds logger settings
type Config struct {
	EnableConsole bool   // show logs in terminal
	LogDir        string // directory for log files
	ServiceName   string // service name for log prefix
	LogLevel      string // "debug", "info", "warn", "error"
}

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

// NewLogger returns a zerolog.Logger configured for daily file rotation + optional console logging
func NewLogger(cfg Config) zerolog.Logger {
	// Ensure log directory exists
	if err := os.MkdirAll(cfg.LogDir, 0755); err != nil {
		panic("failed to create log dir: " + err.Error())
	}

	// Daily file name (yyyy-mm-dd.log)
	fileName := filepath.Join(cfg.LogDir, time.Now().Format("2006-01-02")+".log")

	// Setup file writer
	fileWriter := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    100, // MB
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   false,
	}

	// Choose writers
	var writers []io.Writer
	if cfg.EnableConsole {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		writers = append(writers, consoleWriter)
	}
	writers = append(writers, fileWriter)

	multi := io.MultiWriter(writers...)

	// Set global level from config
	level := parseLogLevel(cfg.LogLevel)
	zerolog.SetGlobalLevel(level)

	// Create logger
	logger := zerolog.New(multi).With().
		Timestamp().
		Str("service", cfg.ServiceName).
		Logger()

	log.Logger = logger
	return logger
}

// Fx module for Vandor
var Module = fx.Module("logger",
	fx.Provide(NewLogger),
)
