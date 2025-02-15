package config

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"strings"
)

func LoggerConfiguration(config *Config) (*zap.Logger, *zap.SugaredLogger) {
	var logger *zap.Logger
	var err error

	switch strings.ToLower(config.Environment().Type) {
	case "development", "dev":
		fmt.Printf("Logger config to: 'dev' mode.\n")
		logger, err = zap.NewDevelopment()
	case "production", "prod":
		fmt.Printf("Logger config to: 'prod' mode.\n")
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
