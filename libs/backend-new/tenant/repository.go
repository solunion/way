package tenant

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/common"
	"github.com/uptrace/bun"
)

func newRepository(db *bun.DB) Repository {
	return &repository{db: db}
}

type Repository interface {
	common.CRUDRepository[Tenant, uuid.UUID]
}

type repository struct {
	db *bun.DB
}

func (r *repository) Create(ctx context.Context, tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(tenant).Exec(ctx)
}

func (r *repository) FindAll(ctx context.Context, tenants *[]Tenant) error {
	return r.db.NewSelect().Model(tenants).Scan(ctx)
}

func (r *repository) FindOne(ctx context.Context, tenant *Tenant, id uuid.UUID) error {
	return r.db.NewSelect().Model(tenant).Where("?Pks", id).Scan(ctx)
}

func (r *repository) Update(ctx context.Context, tenant *Tenant) (sql.Result, error) {
	return r.db.NewUpdate().Model(tenant).Exec(ctx)
}

func (r *repository) Save(ctx context.Context, tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(tenant).On("CONFLICT (id) DO UPDATE").Exec(ctx)
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Tenant)(nil)).Where("?Pks", id).Exec(ctx)
}
