package application

import (
	"github.com/solunion/way/internal/app/api/tenant"
)

type Application struct {
	tenant.WithTenantUserModel `bun:",extend"`
	version                    string `bun:"type:text,notnull"`
}
