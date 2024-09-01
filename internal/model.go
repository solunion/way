package internal

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SystemModel struct {
	gorm.Model
	ID          string         `gorm:"type:text;primary_key"`
	Name        string         `gorm:"type:text;not null"`
	Description sql.NullString `gorm:"type:text"`
}

type UserModel struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string         `gorm:"type:text;not null"`
	Description sql.NullString `gorm:"type:text"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"DATABASE_USER"`
	Pass     string `mapstructure:"DATABASE_PASSWORD"`
	Name     string `mapstructure:"DATABASE_NAME"`
	SSLMode  string `mapstructure:"DATABASE_SSLMODE"`
	TimeZone string `mapstructure:"DATABASE_TIMEZONE"`
}

type Config interface {
	Database() *DatabaseConfig
}

type WayConfig struct {
	DB DatabaseConfig `mapstructure:",squash"`
}

func (w WayConfig) Database() *DatabaseConfig {
	return &w.DB
}

func NewConfiguration() *WayConfig {
	return &WayConfig{
		DB: DatabaseConfig{},
	}
}
