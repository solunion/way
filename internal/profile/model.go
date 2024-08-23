package profile

import "github.com/solunion/way/internal/tenant"

type Profile struct {
	tenant.WithTenantUserModel `gorm:"embedded"`
}
