package tenant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"github.com/uptrace/bun"
)

type Tenant struct {
	bun.BaseModel `bun:"table:tenants,alias:t"`
	ID            uuid.UUID      `bun:"type:uuid,pk,default:uuid_generate_v4()"`
	Name          string         `bun:"type:text,notnull"`
	Description   sql.NullString `bun:"type:text"`
}

type WithTenantUserModel struct {
	internal.UserModel
	Tenant Tenant `bun:"rel:belongs-to,join:tenant_id=id"`
}
