package tenant

import (
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"go.uber.org/zap"
)

type CreateTenantRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTenantResponse struct {
	//CreateTenantRequest
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

//goland:noinspection GoNameStartsWithPackageName
type TenantHttp struct {
	Http       *fiber.App
	Repository TenantRepository
	Log        *zap.SugaredLogger
}

func (t *TenantHttp) Create() {
	t.Http.Post("/tenants", func(ctx fiber.Ctx) error {
		log.Debug("Tenant - Create API called...")

		request := new(CreateTenantRequest)

		if err := ctx.Bind().Body(request); err != nil {
			log.Error("Failed to bind body request:", err)
		}

		tenant := Tenant{
			Name:        request.Name,
			Description: sql.NullString{String: request.Description},
		}

		_, err := t.Repository.Create(&tenant)

		if err != nil {
			log.Error("Failed to create tenant:", err)
		}

		response := CreateTenantResponse{
			Name:        tenant.Name,
			Description: tenant.Description.String,
			ID:          tenant.ID.String(),
		}

		err = ctx.JSON(response)

		if err != nil {
			log.Error("Failed to send response:", err)
		}

		return err
	})
}

func newTenantHttp(app *fiber.App, repository TenantRepository, logger *zap.SugaredLogger) *TenantHttp {
	return &TenantHttp{Http: app, Repository: repository, Log: logger}
}

func registerRoutes(http *TenantHttp) {
	http.Create()
}
