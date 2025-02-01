package config

type EnvironmentConfig struct {
	Type string `mapstructure:"ENVIRONMENT"`
}

type DatabaseConfig struct {
	Uri      string `mapstructure:"DATABASE_URI"`
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"DATABASE_USER"`
	Pass     string `mapstructure:"DATABASE_PASSWORD"`
	Name     string `mapstructure:"DATABASE_NAME"`
	SSLMode  string `mapstructure:"DATABASE_SSLMODE"`
	TimeZone string `mapstructure:"DATABASE_TIMEZONE"`
}

type WebConfig struct {
	Host string `mapstructure:"WEB_HOST"`
	Port int    `mapstructure:"WEB_PORT"`
}

type Config struct {
	DB  DatabaseConfig    `mapstructure:",squash"`
	Env EnvironmentConfig `mapstructure:",squash"`
	Web WebConfig         `mapstructure:",squash"`
}

func (w Config) Environment() *EnvironmentConfig {
	return &w.Env
}

func (w Config) Database() *DatabaseConfig {
	return &w.DB
}

func NewConfiguration() *Config {
	return &Config{
		Env: EnvironmentConfig{Type: "Development"},
		DB:  DatabaseConfig{},
		Web: WebConfig{Host: "localhost", Port: 3000},
	}
}
