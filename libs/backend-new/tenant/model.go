package tenant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/database"
)

type Tenant struct {
	database.BaseModel `bun:"table:tenants,alias:t,extends"`
	ID                 uuid.UUID      `bun:"type:uuid,pk,default:uuid_generate_v4()"`
	Name               string         `bun:"type:text,notnull"`
	Description        sql.NullString `bun:"type:text"`
}

type WithTenantUserModel struct {
	database.BaseModel `bun:",extends"`
	Tenant             Tenant    `bun:"rel:belongs-to,join:tenant_id=id"`
	TenantId           uuid.UUID `bun:"type:uuid,notnull"`
}
