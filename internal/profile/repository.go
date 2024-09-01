package profile

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type ProfileRepository interface {
	internal.CRUDRepository[Profile, uuid.UUID]
}

type profileRepository struct {
	db *gorm.DB
}

func (r profileRepository) FindAll(profiles *[]Profile) (db *gorm.DB) {
	return r.db.Find(&profiles)
}

func (r profileRepository) FindOne(profile *Profile, id uuid.UUID) *gorm.DB {
	return r.db.Where("id = @id", sql.Named("id", id)).Take(&profile)
}

func (r profileRepository) Create(profile *Profile) (db *gorm.DB) {
	return r.db.Create(profile)
}

func (r profileRepository) Update(profile *Profile) (db *gorm.DB) {
	return r.db.Updates(&profile)
}

func (r profileRepository) Save(profile *Profile) (db *gorm.DB) {
	return r.Save(profile)
}

func (r profileRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return r.db.Delete(id)
}

func newProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{db: db}
}

// Interface checks
var _ = interface {
	ProfileRepository
}(&profileRepository{})
