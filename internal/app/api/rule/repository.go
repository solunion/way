package rule

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//goland:noinspection GoNameStartsWithPackageName
type RuleRepository interface {
	internal.CRUDRepository[Rule, uuid.UUID]
}

type ruleRepository struct {
	db *gorm.DB
}

func (r *ruleRepository) FindAll(rules *[]Rule) (db *gorm.DB) {
	return r.db.Preload(clause.Associations).Find(&rules)
}

func (r *ruleRepository) FindOne(rule *Rule, id uuid.UUID) *gorm.DB {
	return r.db.Preload(clause.Associations).Where("id = @id", sql.Named("id", id)).Take(&rule)
}

func (r *ruleRepository) Create(rule *Rule) (db *gorm.DB) {
	return r.db.Create(rule)
}

func (r *ruleRepository) Update(rule *Rule) (db *gorm.DB) {
	return r.db.Updates(&rule)
}

func (r *ruleRepository) Save(rule *Rule) (db *gorm.DB) {
	return r.Save(rule)
}

func (r *ruleRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return r.db.Delete(id)
}

func newRuleRepository(db *gorm.DB) RuleRepository {
	return nil
	//return &ruleRepository{db: db}
}

// Interface checks
//var _ = interface {
//	RuleRepository
//}(&ruleRepository{})
