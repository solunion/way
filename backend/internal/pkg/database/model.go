package database

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	CreatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt *time.Time `bun:",soft_delete,nullzero"`
}

type BaseModel struct {
	Model
	ID          uuid.UUID `bun:"type:uuid,pk,nullzero,notnull,default:uuid_generate_v4()"`
	Name        string    `bun:"type:text,notnull"`
	Description *string   `bun:"type:text"`
}
