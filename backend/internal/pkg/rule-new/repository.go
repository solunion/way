package rule_new

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

func (r *Repository) Create(ctx context.Context, rule Rule) (sql.Result, error) {
	if entity, err := r.toEntity(rule); err != nil {
		return nil, err
	} else {
		return r.db.NewInsert().Model(entity).Exec(ctx)
	}
}

func (r *Repository) FindAll(ctx context.Context, rules *[]Rule) error {
	entities := make([]Entity, 0)
	result := make([]Rule, 0)

	if err := r.db.NewSelect().Model(&entities).Scan(ctx); err != nil {
		return err
	}

	for _, entity := range entities {
		rule, err := r.fromEntity(&entity)
		if err != nil {
			return err
		}
		result = append(result, rule)
	}

	*rules = result

	return nil
}

//
//func (r *Repository) FindOne(ctx context.Context, rule *Rule, id uuid.UUID) error {
//	return r.db.NewSelect().Model(rule).Where("id = ?", id).Scan(ctx)
//}
//
//func (r *Repository) Save(ctx context.Context, rule *Rule) (sql.Result, error) {
//	return r.db.NewInsert().Model(rule).On("CONFLICT (id) DO UPDATE").Exec(ctx)
//}
//
//func (r *Repository) Update(ctx context.Context, rule *Rule) (sql.Result, error) {
//	return r.db.NewUpdate().Model(rule).OmitZero().WherePK().Returning("*").Exec(ctx)
//}
//
//func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (sql.Result, error) {
//	return r.db.NewDelete().Model((*Entity)(nil)).Where("?PKs = ?", id).Exec(ctx)
//}

func (r *Repository) toEntity(rule Rule) (*Entity, error) {
	entity := new(Entity)

	entity.ID = rule.GetID()
	entity.Type = rule.GetType()
	entity.Name = rule.GetName()
	entity.Description = rule.GetDescription()

	if value, err := rule.GetValue(); err != nil {
		return nil, err
	} else {
		entity.Value = value
	}

	return entity, nil
}

func (r *Repository) fromEntity(entity *Entity) (Rule, error) {
	rule := new(BasicRule)

	rule.ID = entity.ID
	rule.Name = entity.Name
	rule.Description = entity.Description
	rule.Type = entity.Type
	rule.Value = entity.Value

	return rule, nil
}
