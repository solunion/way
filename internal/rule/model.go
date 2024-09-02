package rule

import (
	"github.com/solunion/way/internal"
	"github.com/solunion/way/internal/tenant"
)

//goland:noinspection GoNameStartsWithPackageName
type RuleType struct {
	internal.SystemModel `gorm:"embedded"`
}

type Rule struct {
	tenant.WithTenantUserModel `gorm:"embedded"`
	TypeID                     string   `gorm:"type:text;not null"`
	Type                       RuleType `gorm:"foreignKey:TypeID"`
	Data                       string   `gorm:"type:jsonb;not null"`
}
