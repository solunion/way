package database

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Model struct {
	bun.BaseModel
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
