package rule

import (
	"github.com/solunion/way/internal/app/api/tenant"
)

type Rule struct {
	tenant.WithTenantUserModel `gorm:"embedded"`
	Type                       string `gorm:"type:text;not null;check:role_type,type in ('api')"`
	Data                       string `gorm:"type:jsonb;not null"`
}
