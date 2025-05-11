package database

import (
	"log"
	"os"
	"time"

	"github.com/alfariiizi/go-service/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB() *GormDB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	cfg := config.GetConfig()

	db, err := gorm.Open(sqlite.Open(cfg.DB.Destination), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	return &GormDB{
		DB: db,
	}
}
