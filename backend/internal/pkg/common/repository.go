package common

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type CRUDRepository[T any, I uuid.UUID] interface {
	Create(ctx context.Context, model *T) (sql.Result, error)
	FindAll(ctx context.Context, models *[]T) error
	FindOne(ctx context.Context, model *T, id I) error
	Update(ctx context.Context, model *T) (sql.Result, error)
	Save(ctx context.Context, model *T) (sql.Result, error)
	Delete(ctx context.Context, id I) (sql.Result, error)
}
