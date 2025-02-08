package database

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Model struct {
	bun.BaseModel
	CreatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `bun:",soft_delete,nullzero"`
}

type BaseModel struct {
	Model       `bun:",extend"`
	ID          uuid.UUID      `bun:"type:uuid,pk,default:uuid_generate_v4()"`
	Name        string         `bun:"type:text,notnull"`
	Description sql.NullString `bun:"type:text"`
}
