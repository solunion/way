package rule

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

//goland:noinspection GoNameStartsWithPackageName
type RuleRepository struct {
	DB *gorm.DB
}

func (t RuleRepository) FindAll(Rules *[]Rule) (db *gorm.DB) {
	return t.DB.Find(&Rules)
}

func (t RuleRepository) FindOne(Rule *Rule, id uuid.UUID) *gorm.DB {
	return t.DB.Where("id = @id", sql.Named("id", id)).Take(&Rule)
}

func (t RuleRepository) Create(Rule *Rule) (db *gorm.DB) {
	return t.DB.Create(Rule)
}

func (t RuleRepository) Update(Rule *Rule) (db *gorm.DB) {
	return t.DB.Updates(&Rule)
}

func (t RuleRepository) Save(Rule *Rule) (db *gorm.DB) {
	return t.Save(Rule)
}

func (t RuleRepository) Delete(id uuid.UUID) (db *gorm.DB) {
	return t.DB.Delete(id)
}

func NewRuleRepository(DB *gorm.DB) internal.CRUDRepository[Rule, uuid.UUID] {
	return &RuleRepository{DB: DB}
}
