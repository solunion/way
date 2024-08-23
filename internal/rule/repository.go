package rule

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//goland:noinspection GoNameStartsWithPackageName
type RuleRepository struct {
	db *gorm.DB
}

func (t RuleRepository) FindAll(Rules *[]Rule) (db *gorm.DB) {
	return t.db.Preload(clause.Associations).Find(&Rules)
}

func (t RuleRepository) FindOne(Rule *Rule, id uuid.UUID) *gorm.DB {
	return t.db.Preload(clause.Associations).Where("id = @id", sql.Named("id", id)).Take(&Rule)
}

func (t RuleRepository) Create(Rule *Rule) (db *gorm.DB) {
	return t.db.Create(Rule)
}

func (t RuleRepository) Update(Rule *Rule) (db *gorm.DB) {
	return t.db.Updates(&Rule)
}

func (t RuleRepository) Save(Rule *Rule) (db *gorm.DB) {
	return t.Save(Rule)
}

func (t RuleRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.db.Delete(id)
}

func NewRuleRepository(DB *gorm.DB) internal.CRUDRepository[Rule, uuid.UUID] {
	return &RuleRepository{db: DB}
}
