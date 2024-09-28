package tenant

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/solunion/way/internal"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string         `gorm:"type:text;not null"`
	Description sql.NullString `gorm:"type:text"`
}

type WithTenantUserModel struct {
	internal.UserModel
	TenantID uuid.UUID `gorm:"type:uuid;not null"`
	Tenant   Tenant    `gorm:"foreignKey:TenantID"`
}
