package tenant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/database"
)

type NewTenant struct {
	database.BaseModel `bun:"table:tenants,alias:t,extends"`
	Name               string         `bun:"type:text,notnull"`
	Description        sql.NullString `bun:"type:text"`
}

type Tenant struct {
	NewTenant `bun:",extends"`
	ID        uuid.UUID `bun:"type:uuid,pk,default:uuid_generate_v4()"`
}

type WithTenantModel struct {
	Tenant   Tenant    `bun:"rel:belongs-to,join:tenant_id=id"`
	TenantId uuid.UUID `bun:"type:uuid,notnull"`
}
