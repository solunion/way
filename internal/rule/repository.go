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

func (t ruleRepository) FindAll(Rules *[]Rule) (db *gorm.DB) {
	return t.db.Preload(clause.Associations).Find(&Rules)
}

func (t ruleRepository) FindOne(Rule *Rule, id uuid.UUID) *gorm.DB {
	return t.db.Preload(clause.Associations).Where("id = @id", sql.Named("id", id)).Take(&Rule)
}

func (t ruleRepository) Create(Rule *Rule) (db *gorm.DB) {
	return t.db.Create(Rule)
}

func (t ruleRepository) Update(Rule *Rule) (db *gorm.DB) {
	return t.db.Updates(&Rule)
}

func (t ruleRepository) Save(Rule *Rule) (db *gorm.DB) {
	return t.Save(Rule)
}

func (t ruleRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.db.Delete(id)
}

func newRuleRepository(DB *gorm.DB) *ruleRepository {
	return &ruleRepository{db: DB}
}
