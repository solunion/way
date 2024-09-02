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

func (r tenantRepository) FindAll(tenants *[]Tenant) (db *gorm.DB) {
	return r.db.Find(&tenants)
}

func (r tenantRepository) FindOne(tenant *Tenant, id uuid.UUID) *gorm.DB {
	return r.db.Where("id = @id", sql.Named("id", id)).Take(&tenant)
}

func (r tenantRepository) Create(tenant *Tenant) (db *gorm.DB) {
	return r.db.Create(tenant)
}

func (r tenantRepository) Update(tenant *Tenant) (db *gorm.DB) {
	return r.db.Updates(&tenant)
}

func (r tenantRepository) Save(tenant *Tenant) (db *gorm.DB) {
	return r.Save(tenant)
}

func (r tenantRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return r.db.Delete(id)
}

func newTenantRepository(db *gorm.DB) *tenantRepository {
	return &tenantRepository{db: db}
}

// Interface checks
var _ = interface {
	TenantRepository
}(&tenantRepository{})
