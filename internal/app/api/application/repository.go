package application

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type ApplicationRepository interface {
	internal.CRUDRepository[Application, uuid.UUID]
}

type applicationRepository struct {
	db *gorm.DB
}

func (r *applicationRepository) FindAll(applications *[]Application) (db *gorm.DB) {
	return r.db.Find(&applications)
}

func (r *applicationRepository) FindOne(application *Application, id uuid.UUID) (db *gorm.DB) {
	return r.db.Where("id = @id", sql.Named("id", id)).Take(&application)
}

func (r *applicationRepository) Create(application *Application) (db *gorm.DB) {
	return r.db.Create(&application)
}

func (r *applicationRepository) Update(application *Application) (db *gorm.DB) {
	return r.db.Save(&application)
}

func (r *applicationRepository) Save(application *Application) (db *gorm.DB) {
	return r.db.Save(&application)
}

func (r *applicationRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return r.db.Delete(id)
}

func newApplicationRepository(db *gorm.DB) ApplicationRepository {
	return nil
	//return &applicationRepository{db: db}
}

// Interface checks
//var _ = interface {
//	ApplicationRepository
//}(&applicationRepository{})
