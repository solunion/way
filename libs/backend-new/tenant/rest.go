package tenant

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func newRest(app *fiber.App, service Service, logger *zap.SugaredLogger) {
	handlers := &Rest{service: service, log: logger}
	app.Post("/tenants", handlers.Create)
	app.Get("/tenants", handlers.FindAll)
}

type Rest struct {
	service Service
	log     *zap.SugaredLogger
}

func (r *Rest) Create(ctx fiber.Ctx) error {
	r.log.Debug("Tenant - Create API called...")

	request := new(CreateRequestDto)

	if err := ctx.Bind().Body(request); err != nil {
		r.log.Error("Failed to bind body request:", err)
		return err
	}

	tenant := new(Tenant)

	if err := copier.Copy(tenant, request); err != nil {
		r.log.Error("Failed to convert tenant DTO:", err)
		return err
	}

	if err := r.service.Create(ctx.Context(), tenant); err != nil {
		r.log.Error("Failed to create tenant:", err)
		return err
	}

	response := new(ResponseDto)

	if err := copier.Copy(response, tenant); err != nil {
		r.log.Error("Failed to build response:", err)
		return err
	}

	return ctx.JSON(response)
}

func (r *Rest) FindAll(ctx fiber.Ctx) error {
	r.log.Debug("Tenant - FindAll API called...")

	tenants := make([]Tenant, 0)

	if err := r.service.FindAll(ctx.Context(), &tenants); err != nil {
		r.log.Error("Failed to find all tenants:", err)
	}

	response := make([]ResponseDto, 0)

	if err := copier.Copy(&response, tenants); err != nil {
		r.log.Error("Failed to build response:", err)
	}

	return ctx.JSON(response)
}
