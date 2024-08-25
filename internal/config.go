package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func DatabaseConnection(lf fx.Lifecycle, config Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	fmt.Println(config)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Database().Host,
		config.Database().User,
		config.Database().Pass,
		config.Database().Name,
		config.Database().Port,
		config.Database().SSLMode,
		config.Database().TimeZone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	lf.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				db.WithContext(ctx)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				sqlDB, err := db.DB()
				if err != nil {
					return err
				}
				return sqlDB.Close()
			},
		},
	)

	return db, nil
}

func InitConfiguration() (*WayConfig, error) {
	config := NewConfiguration()

	err := godotenv.Load()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error loading .env file: %s", err))
	}
	log.Printf(".env file loaded")

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("DATABASE_PORT env variable not set or has invalid value: %q", os.Getenv("DATABASE_PORT")))
	}

	config.DB.Host = os.Getenv("DATABASE_HOST")
	config.DB.Port = port
	config.DB.User = os.Getenv("DATABASE_USER")
	config.DB.Pass = os.Getenv("DATABASE_PASSWORD")
	config.DB.Name = os.Getenv("DATABASE_NAME")
	config.DB.SSLMode = os.Getenv("DATABASE_SSL_MODE")
	config.DB.TimeZone = os.Getenv("DATABASE_TIMEZONE")

	return config, nil
}
