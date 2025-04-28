package generic

import (
	"github.com/solunion/way/backend/internal/pkg/database"
	"github.com/solunion/way/backend/internal/pkg/tenant"
	"github.com/uptrace/bun"
)

type Rule[T any] struct {
	database.WayBaseModel
	tenant.WithTenantModel
	bun.BaseModel `bun:"table:rules,alias:r"`
	Type          Type `bun:"type:rule_type,notnull" json:"type,omitempty"`
	Value         T    `bun:"type:jsonb,notnull" json:"value,omitempty"`
}
