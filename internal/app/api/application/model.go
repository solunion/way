package application

import (
	"github.com/solunion/way/internal/app/api/tenant"
	"github.com/uptrace/bun"
)

type Application struct {
	bun.BaseModel              `bun:"table:applications"`
	tenant.WithTenantUserModel `bun:", extends"`
	version                    string `bun:"type:text,notnull"`
}
