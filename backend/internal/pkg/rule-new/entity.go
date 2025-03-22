package rule_new

import (
	"encoding/json"
	"github.com/solunion/way/backend/internal/pkg/database"
	"github.com/solunion/way/backend/internal/pkg/tenant"
	"github.com/uptrace/bun"
)

type Entity struct {
	database.WayBaseModel
	tenant.WithTenantModel
	bun.BaseModel `bun:"table:rules,alias:r"`
	Type          Type            `bun:"type:rule_type,notnull" json:",omitempty"`
	Value         json.RawMessage `bun:"type:jsonb,notnull" json:"value,omitempty"`
}
