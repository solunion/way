package common

import "context"

type Service[T any] interface {
	Create(ctx context.Context, newModel *T) error
	FindAll(ctx context.Context, models *[]T) error
}
