package tenant

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

//goland:noinspection GoNameStartsWithPackageName
type TenantHttp struct {
	Http       *fiber.App
	Repository TenantRepository
	Log        *zap.SugaredLogger
}

func (t *TenantHttp) registerRoutes() {
	t.Create()
}

func (t *TenantHttp) Create() {
	t.Http.Post("/tenants", func(ctx fiber.Ctx) error {
		log.Debug("Tenant - Create API called...")

		request := new(CreateTenantRequest)

		if err := ctx.Bind().Body(request); err != nil {
			log.Error("Failed to bind body request:", err)
			return err
		}

		tenant := new(Tenant)

		if err := copier.Copy(&tenant, &request); err != nil {
			log.Error("Failed to convert tenant DTO:", err)
			return err
		}

		if _, err := t.Repository.Create(tenant); err != nil {
			log.Error("Failed to create tenant:", err)
			return err
		}

		response := new(CreateTenantResponse)

		if err := copier.Copy(&response, tenant); err != nil {
			log.Error("Failed to build response:", err)
			return err
		}

		if err := ctx.JSON(response); err != nil {
			log.Error("Failed to send response:", err)
		}

		return nil
	})
}

func registerTenantHttp(app *fiber.App, repository TenantRepository, logger *zap.SugaredLogger) *TenantHttp {
	api := &TenantHttp{Http: app, Repository: repository, Log: logger}
	api.registerRoutes()
	return api
}
