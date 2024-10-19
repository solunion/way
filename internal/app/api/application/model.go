package application

import (
	"github.com/solunion/way/internal/app/api/tenant"
)

type Application struct {
	tenant.WithTenantUserModel `bun:"table:applications,alias:a,extends"`
	version                    string `bun:"type:text,notnull"`
}
