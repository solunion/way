package profile

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal/pkg/database"
	"github.com/uptrace/bun"
)

//goland:noinspection GoNameStartsWithPackageName
type ProfileRepository interface {
	database.CRUDRepository[Profile, uuid.UUID]
}

type profileRepository struct {
	db  *bun.DB
	ctx context.Context
}

func (r *profileRepository) FindAll(profiles *[]Profile) error {
	return r.db.NewSelect().Model(profiles).Scan(r.ctx)
}

func (r *profileRepository) FindOne(profile *Profile, id uuid.UUID) error {
	return r.db.NewSelect().Model(profile).Where("?Pks", id).Scan(r.ctx)
}

func (r *profileRepository) Create(profile *Profile) (sql.Result, error) {
	return r.db.NewInsert().Model(profile).Exec(r.ctx)
}

func (r *profileRepository) Update(profile *Profile) (sql.Result, error) {
	return r.db.NewUpdate().Model(profile).Exec(r.ctx)
}

func (r *profileRepository) Save(profile *Profile) (sql.Result, error) {
	return r.db.NewInsert().Model(profile).On("CONFLICT (id) DO UPDATE").Exec(r.ctx)
}

func (r *profileRepository) Delete(id uuid.UUID) (sql.Result, error) {
	return r.db.NewDelete().Model((*Profile)(nil)).Where("?Pks", id).Exec(r.ctx)
}

func newProfileRepository(ctx context.Context, db *bun.DB) ProfileRepository {
	return &profileRepository{db: db, ctx: ctx}
}

// Interface checks
var _ = interface {
	database.CRUDRepository[Profile, uuid.UUID]
}(&profileRepository{})
