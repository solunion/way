package tenant

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/internal/pkg/common"
	"github.com/uptrace/bun"
)

func newRepository(db *bun.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	common.CRUDRepository[Tenant, uuid.UUID]
	db *bun.DB
}

func (r *Repository) Create(ctx context.Context, tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(tenant).Exec(ctx)
}

func (r *Repository) FindAll(ctx context.Context, tenants *[]Tenant) error {
	return r.db.NewSelect().Model(tenants).Scan(ctx)
}

func (r *Repository) FindOne(ctx context.Context, tenant *Tenant, id uuid.UUID) error {
	return r.db.NewSelect().Model(tenant).Where("id = ?", id).Scan(ctx)
}

func (r *Repository) Update(ctx context.Context, tenant *Tenant) (sql.Result, error) {
	return r.db.NewUpdate().Model(tenant).OmitZero().WherePK().Returning("*").Exec(ctx)
}

func (r *Repository) Save(ctx context.Context, tenant *Tenant) (sql.Result, error) {
	return r.db.NewInsert().Model(tenant).On("CONFLICT (id) DO UPDATE").Exec(ctx)
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Tenant)(nil)).Where("?PKs = ?", id).Exec(ctx)
}
