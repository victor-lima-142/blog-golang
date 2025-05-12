package db

import (
	"blog/core/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDialect() *gorm.Dialector {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get variables
	dbURL := os.Getenv("DATABASE_URL")
	dialector := postgres.Open(dbURL)

	return &dialector
}

func GetConnection() (*gorm.DB, error) {
	dialector := getDialect()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			ParameterizedQueries:      true,        // Don't include arguments in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(*dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err) // Use %w for wrapping errors in Go 1.13+
	}

	err = db.AutoMigrate(&models.User{}, &models.Profile{}, &models.UserFollowing{}, &models.Article{})
	if err != nil {
		return nil, fmt.Errorf("error in database prepare: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
