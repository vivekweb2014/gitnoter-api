package db

import (
	"fmt"

	"github.com/git-noter/gitnoter-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect opens a new connection with database specified by configuration.
// It returns gorm db instance and any connection error encountered.
func Connect(config config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=%s TimeZone=UTC",
		config.Host, config.Username, config.Password, config.DBName, config.Port, "gitnoter", config.SSLMode)

	gormConfig := &gorm.Config{}
	if config.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	return gorm.Open(postgres.New(postgres.Config{
		DriverName: config.DriverName,
		DSN:        dsn,
	}), gormConfig)
}
