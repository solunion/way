package tenant

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

//goland:noinspection GoNameStartsWithPackageName
type Rest struct {
	Http    *fiber.App
	service Service
	log     *zap.SugaredLogger
}

func (t *Rest) registerRoutes() {
	t.Create()
}

func (t *Rest) Create() {
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

		if err := t.service.Create(ctx.Context(), tenant); err != nil {
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

func newRest(app *fiber.App, service Service, logger *zap.SugaredLogger) *Rest {
	api := &Rest{Http: app, service: service, log: logger}
	api.registerRoutes()
	return api
}
