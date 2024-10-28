package rule

import (
	"github.com/solunion/way/internal/pkg/tenant"
)

type Rule struct {
	tenant.WithTenantUserModel `bun:"table:rules,alias:r,extends"`
	Type                       string `gorm:"type:text;not null;check:role_type,type in ('API')"`
	Data                       string `gorm:"type:jsonb;not null"`
}
