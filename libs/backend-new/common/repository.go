package common

import (
	"database/sql"
	"github.com/google/uuid"
)

type CRUDRepository[T any, I uuid.UUID] interface {
	FindAll(models *[]T) error
	FindOne(model *T, id I) error
	Create(model *T) (sql.Result, error)
	Update(model *T) (sql.Result, error)
	Save(model *T) (sql.Result, error)
	Delete(id I) (sql.Result, error)
}
