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

func DatabaseConnection2(lifecycle fx.Lifecycle) *gorm.DB {
	var db *gorm.DB

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				err := godotenv.Load()
				if err != nil {
					return errors.New(fmt.Sprintf("error loading .env file: %s", err))
				}
				log.Printf(".env file loaded")

				port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

				if err != nil {
					return errors.New(fmt.Sprintf("DATABASE_PORT env variable not set or has invalid value: %q", os.Getenv("DATABASE_PORT")))
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
					return errors.New(fmt.Sprintf("failed to open db connection: %s", err))
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				if db != nil {

					if conn, err := db.DB(); err != nil {
						return err
					} else {
						if err := conn.Close(); err != nil {
							return err
						}
					}
				}
				return nil
			},
		})

	return db
}
