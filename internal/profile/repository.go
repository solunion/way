package profile

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type ProfileRepository struct {
	DB *gorm.DB
}

func (t ProfileRepository) FindAll(Profiles *[]Profile) (db *gorm.DB) {
	return t.DB.Find(&Profiles)
}

func (t ProfileRepository) FindOne(Profile *Profile, id uuid.UUID) *gorm.DB {
	return t.DB.Where("id = @id", sql.Named("id", id)).Take(&Profile)
}

func (t ProfileRepository) Create(Profile *Profile) (db *gorm.DB) {
	return t.DB.Create(Profile)
}

func (t ProfileRepository) Update(Profile *Profile) (db *gorm.DB) {
	return t.DB.Updates(&Profile)
}

func (t ProfileRepository) Save(Profile *Profile) (db *gorm.DB) {
	return t.Save(Profile)
}

func (t ProfileRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.DB.Delete(id)
}

func NewProfileRepository(DB *gorm.DB) internal.CRUDRepository[Profile, uuid.UUID] {
	return &ProfileRepository{DB: DB}
}
