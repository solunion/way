package tenant

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/common"
	"go.uber.org/zap"
)

func newRest(service *Service, log *zap.SugaredLogger) *Rest {
	return &Rest{service: service, log: log}
}

type Rest struct {
	common.Rest[Tenant]
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
