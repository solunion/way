package common

import (
	"context"
	"github.com/google/uuid"
)

type Service[T any] interface {
	Create(ctx context.Context, newModel *T) error
	GetAll(ctx context.Context, models *[]T) error
	GetById(ctx context.Context, model *T, id uuid.UUID) error
}
