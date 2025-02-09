package common

type Service[T any] interface {
	Create(newModel *T) error
	FindAll(models *[]T) error
}
