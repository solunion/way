package rule

import (
	"database/sql/driver"
	"fmt"
	"github.com/solunion/way/backend/internal/pkg/database"
	"github.com/solunion/way/backend/internal/pkg/tenant"
	"github.com/uptrace/bun"
)

type Type struct {
	slug string
}

func (r *Type) String() string {
	return r.slug
}

func (r *Type) Value() (driver.Value, error) {
	return r.slug, nil
}

func (r *Type) Scan(src interface{}) error {
	byteArray := src.([]uint8)
	if value, err := TypeFromString(string(byteArray)); err != nil {
		return err
	} else {
		*r = value
	}
	return nil
}

var (
	Unknown = Type{""}
	Http    = Type{"HTTP"}
	Route   = Type{"ROUTE"}
)

func TypeFromString(s string) (Type, error) {
	switch s {
	case "HTTP":
		return Http, nil
	case "ROUTE":
		return Route, nil
	default:
		return Unknown, fmt.Errorf("unknown rule type: %s", s)
	}
}

type Rule struct {
	database.WayBaseModel
	tenant.WithTenantModel
	bun.BaseModel `bun:"table:rules,alias:r"`
	Type          Type   `bun:"type:rule_type,notnull"`
	Value         string `bun:"type:jsonb,notnull"`
}
