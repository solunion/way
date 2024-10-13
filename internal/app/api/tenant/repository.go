package tenant

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"github.com/uptrace/bun"
)

//goland:noinspection GoNameStartsWithPackageName
type TenantRepository interface {
	internal.CRUDRepository[Tenant, uuid.UUID]
}

type tenantRepository struct {
	db  *bun.DB
	ctx context.Context
}

func (r *tenantRepository) FindAll(tenants *[]Tenant) error {
	return r.db.NewSelect().Model(&tenants).Scan(r.ctx)
}

func (r *tenantRepository) FindOne(tenant *Tenant, id uuid.UUID) error {
	return r.db.NewSelect().Model(&tenant).Where("id = @id", sql.Named("id", id)).Scan(r.ctx)
}

func (r *tenantRepository) Create(tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(&tenant).Exec(r.ctx)
}

func (r *tenantRepository) Update(tenant *Tenant) (sql.Result, error) {
	return r.db.NewUpdate().Model(&tenant).Exec(r.ctx)
}

func (r *tenantRepository) Save(tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(&tenant).On("CONFLICT (id) DO UPDATE").Exec(r.ctx)
}

func (r *tenantRepository) Delete(id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Tenant)(nil)).Where("?Pks", id).Exec(r.ctx)
}

func newTenantRepository(db *bun.DB, ctx context.Context) TenantRepository {
	return &tenantRepository{db: db, ctx: ctx}
}

// Interface checks
var _ = interface {
	TenantRepository
}(&tenantRepository{})
