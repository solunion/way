package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CRUDRepository[T any, I string | uuid.UUID] interface {
	FindAll(models *[]T) *gorm.DB
	FindOne(model *T, id I) *gorm.DB
	Create(model *T) *gorm.DB
	Update(model *T) *gorm.DB
	Save(model *T) *gorm.DB
	Delete(id I) *gorm.DB
}
