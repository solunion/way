package profile

import (
	"github.com/solunion/way/internal/app/api/tenant"
)

type Profile struct {
	tenant.WithTenantUserModel `bun:",extend"`
}
