package configs

type EnvironmentConfig struct {
	Type string `mapstructure:"ENVIRONMENT"`
}

type DatabaseConfig struct {
	Type     string `mapstructure:"DATABASE_TYPE"`
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"DATABASE_USER"`
	Pass     string `mapstructure:"DATABASE_PASSWORD"`
	Name     string `mapstructure:"DATABASE_NAME"`
	SSLMode  string `mapstructure:"DATABASE_SSLMODE"`
	TimeZone string `mapstructure:"DATABASE_TIMEZONE"`
}

type Config struct {
	DB  DatabaseConfig    `mapstructure:",squash"`
	Env EnvironmentConfig `mapstructure:",squash"`
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
	}
}
