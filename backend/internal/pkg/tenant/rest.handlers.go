package tenant

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"go.uber.org/zap"
)

func newRest(service *Service, log *zap.SugaredLogger) *Rest {
	return &Rest{service: service, log: log}
}

type Rest struct {
	handlers.Rest[Tenant]
	service *Service
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

func (r *Rest) GetAll(ctx fiber.Ctx) error {
	r.log.Debug("Tenant - GetAll API called...")

	tenants := make([]Tenant, 0)

	if err := r.service.GetAll(ctx.Context(), &tenants); err != nil {
		r.log.Error("Failed to find all tenants:", err)
	}

	response := make([]ResponseDto, 0)

	if err := copier.Copy(&response, tenants); err != nil {
		r.log.Error("Failed to build response:", err)
	}

	return ctx.JSON(response)
}

func (r *Rest) GetById(ctx fiber.Ctx) error {
	r.log.Debug("Tenant - GetById API called...")

	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		r.log.Error("Failed to parse id:", err)
		return err
	}

	tenant := new(Tenant)

	if err := r.service.GetById(ctx.Context(), tenant, id); err != nil {
		r.log.Error("Failed to find tenant by id:", err)
		return err
	}

	response := new(ResponseDto)

	if err := copier.Copy(response, tenant); err != nil {
		r.log.Error("Failed to build response:", err)
		return err
	}

	return ctx.JSON(response)
}

func (r *Rest) Update(ctx fiber.Ctx) error {
	r.log.Debug("Tenant - Update API called...")

	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		r.log.Error("Failed to parse id:", err)
		return err
	}

	request := new(UpdateRequestDto)

	if err := ctx.Bind().Body(request); err != nil {
		r.log.Error("Failed to bind body request:", err)
		return err
	}

	tenant := new(Tenant)

	if err := copier.CopyWithOption(tenant, request, copier.Option{IgnoreEmpty: true}); err != nil {
		r.log.Error("Failed to convert tenant DTO:", err)
		return err
	}

	tenant.ID = id

	if err := r.service.Update(ctx.Context(), tenant); err != nil {
		r.log.Error("Failed to update tenant:", err)
		return err
	}

	response := make([]ResponseDto, 0)

	if err := copier.Copy(&response, tenant); err != nil {
		r.log.Error("Failed to build response:", err)
	}

	return ctx.JSON(response)
}

func (r *Rest) Delete(ctx fiber.Ctx) error {
	r.log.Debug("Tenant - Delete API called...")

	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		r.log.Error("Failed to parse id:", err)
		return err
	}

	if err := r.service.Delete(ctx.Context(), id); err != nil {
		r.log.Error("Failed to delete tenant:", err)
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
