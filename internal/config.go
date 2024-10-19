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
	"go.uber.org/zap"
	"log"
	"strings"
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
	fx.Provide(LoggerConfiguration),
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

	fmt.Println("Configuration successfully loaded!!!")

	return config, nil
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

func LoggerConfiguration(config Config) (*zap.Logger, *zap.SugaredLogger) {
	var logger *zap.Logger
	var err error

	switch strings.ToLower(config.Environment().Type) {
	case "development", "dev":
		logger, err = zap.NewDevelopment()
	case "production", "prod":
		logger, err = zap.NewProduction()
	default:
		fmt.Printf("Unknown environment type: %s.\n", config.Environment().Type)
		logger = zap.NewExample()
	}

	if err != nil {
		log.Fatalf("Could not initialize zap logger: %v", err)
	}

	return logger, logger.Sugar()
}
