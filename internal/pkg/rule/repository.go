package rule

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal/pkg/database"
	"github.com/uptrace/bun"
)

//goland:noinspection GoNameStartsWithPackageName
type RuleRepository interface {
	database.CRUDRepository[Rule, uuid.UUID]
}

type ruleRepository struct {
	db  *bun.DB
	ctx context.Context
}

func (r *ruleRepository) FindAll(rules *[]Rule) error {
	return r.db.NewSelect().Model(rules).Scan(r.ctx)
}

func (r *ruleRepository) FindOne(rule *Rule, id uuid.UUID) error {
	return r.db.NewSelect().Model(rule).Where("?Pks", id).Scan(r.ctx)
}

func (r *ruleRepository) Create(rule *Rule) (sql.Result, error) {
	return r.db.NewInsert().Model(rule).Exec(r.ctx)
}

func (r *ruleRepository) Update(rule *Rule) (sql.Result, error) {
	return r.db.NewUpdate().Model(rule).Exec(r.ctx)
}

func (r *ruleRepository) Save(rule *Rule) (sql.Result, error) {
	return r.db.NewInsert().Model(rule).On("CONFLICT (id) DO UPDATE").Exec(r.ctx)
}

func (r *ruleRepository) Delete(id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Rule)(nil)).Where("?Pks", id).Exec(r.ctx)
}

func newRuleRepository(ctx context.Context, db *bun.DB) RuleRepository {
	return &ruleRepository{db: db, ctx: ctx}
}

// Interface checks
var _ = interface {
	database.CRUDRepository[Rule, uuid.UUID]
}(&ruleRepository{})
