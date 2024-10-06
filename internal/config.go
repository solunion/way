package internal

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Module("config",
	fx.Provide(
		fx.Annotate(
			InitConfiguration,
			fx.As(new(Config)),
		),
	),
	fx.Provide(DatabaseConnection),
	fx.Invoke(MigrateDatabase),
)

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

func DatabaseConnection(config Config) (*gorm.DB, error) {
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

func MigrateDatabase(db *gorm.DB, config Config) {
	//FIXME: missing GO migration files (issue https://github.com/golang-migrate/migrate/issues/1177)
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
			config.Database().Type,
			config.Database().User,
			config.Database().Pass,
			config.Database().Host,
			config.Database().Port,
			config.Database().Name,
			config.Database().SSLMode),
	)

	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}
}
