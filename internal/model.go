package internal

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Model struct {
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete"`
}

type SystemModel struct {
	Model       `bun:",extend"`
	ID          string         `bun:"type:text,pk"`
	Name        string         `bun:"type:text,notnull"`
	Description sql.NullString `bun:"type:text"`
}

type UserModel struct {
	Model       `bun:",extend"`
	ID          uuid.UUID      `bun:"type:uuid,pk,default:uuid_generate_v4()"`
	Name        string         `bun:"type:text,notnull"`
	Description sql.NullString `bun:"type:text"`
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
