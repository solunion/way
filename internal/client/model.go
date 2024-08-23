package client

import "github.com/solunion/way/internal/tenant"

type Client struct {
	tenant.WithTenantUserModel `gorm:"embedded"`
}
