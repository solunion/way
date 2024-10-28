package profile

import (
	"github.com/solunion/way/internal/pkg/tenant"
)

type Profile struct {
	tenant.WithTenantUserModel `bun:"table:profiles,alias:p,extends"`
}
