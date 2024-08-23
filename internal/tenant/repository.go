package tenant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type TenantRepository struct {
	DB *gorm.DB
}

func (t TenantRepository) FindAll(tenants *[]Tenant) (db *gorm.DB) {
	return t.DB.Find(&tenants)
}

func (t TenantRepository) FindOne(tenant *Tenant, id uuid.UUID) *gorm.DB {
	return t.DB.Where("id = @id", sql.Named("id", id)).Take(&tenant)
}

func (t TenantRepository) Create(tenant *Tenant) (db *gorm.DB) {
	return t.DB.Create(tenant)
}

func (t TenantRepository) Update(tenant *Tenant) (db *gorm.DB) {
	return t.DB.Updates(&tenant)
}

func (t TenantRepository) Save(tenant *Tenant) (db *gorm.DB) {
	return t.Save(tenant)
}

func (t TenantRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.DB.Delete(id)
}

func NewTenantRepository(DB *gorm.DB) internal.CRUDRepository[Tenant, uuid.UUID] {
	return &TenantRepository{DB: DB}
}
