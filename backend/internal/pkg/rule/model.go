package rule

import (
	"github.com/solunion/way/backend/internal/pkg/database"
	"github.com/solunion/way/backend/internal/pkg/tenant"
	"github.com/uptrace/bun"
)

type Rule struct {
	database.WayBaseModel
	tenant.WithTenantModel
	bun.BaseModel `bun:"table:rules,alias:r"`
	Type          Type                   `bun:"type:rule_type,notnull" json:"type,omitempty"`
	Value         map[string]interface{} `bun:"type:jsonb,notnull" json:"value,omitempty"`
}

type RuleValues interface {
	HttpRuleValue | RouteRuleValue
	GetRuleType() Type
}
