package profile

import (
	"github.com/solunion/way/internal/app/api/tenant"
	"github.com/uptrace/bun"
)

type Profile struct {
	bun.BaseModel              `bun:"table:profiles"`
	tenant.WithTenantUserModel `bun:",extends"`
}
