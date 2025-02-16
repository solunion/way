package config

import (
	"errors"
	"github.com/spf13/viper"
)

func ViperConfiguration() (*Config, error) {
	config := NewConfiguration()

	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	viper.WatchConfig()

	return config, nil
}
