package rule

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
	common.CRUDRepository[Rule, uuid.UUID]
	db *bun.DB
}

func (r *Repository) Create(ctx context.Context, rule *Rule) (sql.Result, error) {
	return r.db.NewInsert().Model(rule).Exec(ctx)
}

func (r *Repository) FindAll(ctx context.Context, rules *[]Rule) error {
	return r.db.NewSelect().Model(rules).Scan(ctx)
}

func (r *Repository) FindOne(ctx context.Context, rule *Rule, id uuid.UUID) error {
	return r.db.NewSelect().Model(rule).Where("id = ?", id).Scan(ctx)
}

func (r *Repository) Save(ctx context.Context, rule *Rule) (sql.Result, error) {
	return r.db.NewInsert().Model(rule).On("CONFLICT (id) DO UPDATE").Exec(ctx)
}

func (r *Repository) Update(ctx context.Context, rule *Rule) (sql.Result, error) {
	return r.db.NewUpdate().Model(rule).OmitZero().WherePK().Returning("*").Exec(ctx)
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Rule)(nil)).Where("?PKs = ?", id).Exec(ctx)
}
