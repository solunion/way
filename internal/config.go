package internal

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func DatabaseConnection() (*gorm.DB, error) {
	var db *gorm.DB

	err := godotenv.Load()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error loading .env file: %s", err))
	}
	log.Printf(".env file loaded")

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("DATABASE_PORT env variable not set or has invalid value: %q", os.Getenv("DATABASE_PORT")))
	}

	var (
		host     = os.Getenv("DATABASE_HOST")
		user     = os.Getenv("DATABASE_USER")
		password = os.Getenv("DATABASE_PASSWORD")
		name     = os.Getenv("DATABASE_NAME")
		sslMode  = os.Getenv("DATABASE_SSL_MODE")
		timeZone = os.Getenv("DATABASE_TIMEZONE")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, name, port, sslMode, timeZone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open db connection: %s", err))
	}

	return db, nil
}
