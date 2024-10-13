package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Module("config",
	fx.Provide(ContextConfiguration),
	fx.Provide(
		fx.Annotate(
			InitConfiguration,
			fx.As(new(Config)),
		),
	),
	fx.Provide(BunDatabaseConnection),
	//fx.Provide(GORMDatabaseConnection),
)

func ContextConfiguration() context.Context {
	return context.Background()
}

func InitConfiguration() (*WayConfig, error) {
	config := NewConfiguration()

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	fmt.Printf("Configuration loaded: %+v\n", config)

	return config, nil
}

func GORMDatabaseConnection(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Database().Host,
		config.Database().User,
		config.Database().Pass,
		config.Database().Name,
		config.Database().Port,
		config.Database().SSLMode,
		config.Database().TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func BunDatabaseConnection(config Config) *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Database().User,
		config.Database().Pass,
		config.Database().Host,
		config.Database().Port,
		config.Database().Name,
		config.Database().SSLMode,
	)

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return bun.NewDB(conn, pgdialect.New())
}
