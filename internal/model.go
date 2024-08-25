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
	Host     string
	Port     int
	User     string
	Pass     string
	Name     string
	SSLMode  string
	TimeZone string
}

type Config interface {
	Database() *DatabaseConfig
}

type WayConfig struct {
	DB DatabaseConfig
}

func (w WayConfig) Database() *DatabaseConfig {

	return &w.DB
}

func NewConfiguration() *WayConfig {
	return &WayConfig{
		DB: DatabaseConfig{},
	}
}
