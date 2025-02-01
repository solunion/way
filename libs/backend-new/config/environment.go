package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

// FIXME: create better configuration function
func EnvironmentConfiguration() (*Config, error) {
	config := NewConfiguration()

	config.DB.Type = os.Getenv("DATABASE_TYPE")
	config.DB.Host = os.Getenv("DATABASE_HOST")
	config.DB.Port, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	config.DB.User = os.Getenv("DATABASE_USER")
	config.DB.Pass = os.Getenv("DATABASE_PASSWORD")
	config.DB.Name = os.Getenv("DATABASE_NAME")
	config.DB.SSLMode = os.Getenv("DATABASE_SSLMODE")
	config.DB.TimeZone = os.Getenv("DATABASE_TIMEZONE")

	return config, nil
}

func DotEnvConfiguration() (*Config, error) {
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
	fmt.Printf("%+v\n", config)

	return config, nil
}
