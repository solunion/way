package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

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

	return config, nil
}
