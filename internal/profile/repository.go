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

func (t profileRepository) FindAll(Profiles *[]Profile) (db *gorm.DB) {
	return t.db.Find(&Profiles)
}

func (t profileRepository) FindOne(Profile *Profile, id uuid.UUID) *gorm.DB {
	return t.db.Where("id = @id", sql.Named("id", id)).Take(&Profile)
}

func (t profileRepository) Create(Profile *Profile) (db *gorm.DB) {
	return t.db.Create(Profile)
}

func (t profileRepository) Update(Profile *Profile) (db *gorm.DB) {
	return t.db.Updates(&Profile)
}

func (t profileRepository) Save(Profile *Profile) (db *gorm.DB) {
	return t.Save(Profile)
}

func (t profileRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.db.Delete(id)
}

func newProfileRepository(DB *gorm.DB) *profileRepository {
	return &profileRepository{db: DB}
}
