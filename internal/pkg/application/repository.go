package application

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal/pkg/common/database"
	"github.com/uptrace/bun"
)

//goland:noinspection GoNameStartsWithPackageName
type ApplicationRepository interface {
	database.CRUDRepository[Application, uuid.UUID]
}

type applicationRepository struct {
	db  *bun.DB
	ctx context.Context
}

func (r *applicationRepository) FindAll(applications *[]Application) error {
	return r.db.NewSelect().Model(applications).Scan(r.ctx)
}

func (r *applicationRepository) FindOne(application *Application, id uuid.UUID) error {
	return r.db.NewSelect().Model(application).Where("?Pks", id).Scan(r.ctx)
}

func (r *applicationRepository) Create(application *Application) (sql.Result, error) {
	return r.db.NewInsert().Model(application).Exec(r.ctx)
}

func (r *applicationRepository) Update(application *Application) (sql.Result, error) {
	return r.db.NewUpdate().Model(application).Exec(r.ctx)
}

func (r *applicationRepository) Save(application *Application) (sql.Result, error) {
	return r.db.NewInsert().Model(application).On("CONFLICT (id) DO UPDATE").Exec(r.ctx)
}

func (r *applicationRepository) Delete(id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Application)(nil)).Where("?Pks", id).Exec(r.ctx)
}

func newApplicationRepository(db *bun.DB, ctx context.Context) ApplicationRepository {
	return &applicationRepository{db: db, ctx: ctx}
}

// Interface checks
var _ = interface {
	database.CRUDRepository[Application, uuid.UUID]
}(&applicationRepository{})
