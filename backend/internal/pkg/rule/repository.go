package rule

import (
	"context"
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/solunion/way/backend/internal/pkg/common"
	"github.com/uptrace/bun"
)

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	common.CRUDRepository[Rule[any], uuid.UUID]
	db *bun.DB
}

func (r *Repository) Create(ctx context.Context, rule *Rule[any]) (sql.Result, error) {
	return r.db.NewInsert().Model(rule).Exec(ctx)
}

func (r *Repository) FindAll(ctx context.Context, rules *[]Rule[any]) error {
	if ctx.Value("rule_type") == nil {
		return r.db.NewSelect().Model(rules).Scan(ctx)
	} else {
		return r.db.NewSelect().Model(rules).Where("type = ?", strings.ToUpper(ctx.Value("rule_type").(string))).Scan(ctx)
	}
}

func (r *Repository) Update(ctx context.Context, rule *Rule[any]) (sql.Result, error) {
	return r.db.NewUpdate().Model(rule).OmitZero().WherePK().Returning("*").Exec(ctx)
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Rule[any])(nil)).Where("?PKs = ?", id).Exec(ctx)
}
