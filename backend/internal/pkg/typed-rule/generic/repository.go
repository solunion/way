package generic

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
	common.CRUDRepository[Rule, uuid.UUID]
	db *bun.DB
}

func (r *Repository) Create(ctx context.Context, rule *Rule) (sql.Result, error) {
	return r.db.NewInsert().Model(rule).Exec(ctx)
}

func (r *Repository) FindAll(ctx context.Context, rules *[]Rule) error {
	return r.db.NewSelect().Model(rules).Scan(ctx)
}

func (r *Repository) FindAllWithType(ctx context.Context, rules *[]Rule) error {
	return r.db.NewSelect().Model(rules).Where("type = ?", strings.ToUpper(ctx.Value("rule_type").(string))).Scan(ctx)
}
