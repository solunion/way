package tenant

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/common"
	"github.com/uptrace/bun"
)

//goland:noinspection GoNameStartsWithPackageName
type Repository interface {
	common.CRUDRepository[Tenant, uuid.UUID]
}

type tenantRepository struct {
	db  *bun.DB
	ctx context.Context
}

func (r *tenantRepository) FindAll(tenants *[]Tenant) error {
	return r.db.NewSelect().Model(tenants).Scan(r.ctx)
}

func (r *tenantRepository) FindOne(tenant *Tenant, id uuid.UUID) error {
	return r.db.NewSelect().Model(tenant).Where("?Pks", id).Scan(r.ctx)
}

func (r *tenantRepository) Create(tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(tenant).Exec(r.ctx)
}

func (r *tenantRepository) Update(tenant *Tenant) (sql.Result, error) {
	return r.db.NewUpdate().Model(tenant).Exec(r.ctx)
}

func (r *tenantRepository) Save(tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(tenant).On("CONFLICT (id) DO UPDATE").Exec(r.ctx)
}

func (r *tenantRepository) Delete(id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Tenant)(nil)).Where("?Pks", id).Exec(r.ctx)
}

func newRepository(ctx context.Context, db *bun.DB) Repository {
	return &tenantRepository{db: db, ctx: ctx}
}

// Interface checks
var _ = interface {
	common.CRUDRepository[Tenant, uuid.UUID]
}(&tenantRepository{})
