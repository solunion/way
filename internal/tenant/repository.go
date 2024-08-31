package tenant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type TenantRepository interface {
	internal.CRUDRepository[Tenant, uuid.UUID]
}

type tenantRepository struct {
	db *gorm.DB
}

func (t tenantRepository) FindAll(tenants *[]Tenant) (db *gorm.DB) {
	return t.db.Find(&tenants)
}

func (t tenantRepository) FindOne(tenant *Tenant, id uuid.UUID) *gorm.DB {
	return t.db.Where("id = @id", sql.Named("id", id)).Take(&tenant)
}

func (t tenantRepository) Create(tenant *Tenant) (db *gorm.DB) {
	return t.db.Create(tenant)
}

func (t tenantRepository) Update(tenant *Tenant) (db *gorm.DB) {
	return t.db.Updates(&tenant)
}

func (t tenantRepository) Save(tenant *Tenant) (db *gorm.DB) {
	return t.Save(tenant)
}

func (t tenantRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.db.Delete(id)
}

func newTenantRepository(DB *gorm.DB) *tenantRepository {
	return &tenantRepository{db: DB}
}
