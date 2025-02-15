package tenant

import (
	"github.com/google/uuid"
	"github.com/solunion/way/backend/internal/pkg/database"
	"github.com/uptrace/bun"
)

type DatabaseBaseModel = database.BaseModel

type NewTenant struct {
	Tenant `bun:",extend"`
	ID     uuid.UUID `bun:"-"`
}

type Tenant struct {
	DatabaseBaseModel
	bun.BaseModel `bun:"table:tenants,alias:t"`
}

type WithTenantModel struct {
	Tenant   Tenant    `bun:"rel:belongs-to,join:tenant_id=id"`
	TenantId uuid.UUID `bun:"type:uuid,notnull"`
}
