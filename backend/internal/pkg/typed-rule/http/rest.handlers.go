package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
	"go.uber.org/zap"
)

type Rest struct {
	handlers.Rest[HttpRule]
	service *Service
	log     *zap.SugaredLogger
}

func NewRest(service *Service, log *zap.SugaredLogger) *Rest {
	return &Rest{service: service, log: log}
}

func (r *Rest) Create(ctx fiber.Ctx) error {
	r.log.Debug("Rule - Create: API called...")

	request := new(CreateRequest)

	if err := ctx.Bind().Body(request); err != nil {
		r.log.Error("Failed to bind body request:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rule, err := r.toModel(request)

	if err != nil {
		r.log.Error("Failed to build request model:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = r.service.Create(ctx.Context(), rule)

	if err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := r.fromModel(rule)

	if err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Rest) toModel(request *CreateRequest) (*HttpRule, error) {
	rule := new(HttpRule)
	rule.Name = request.Name
	rule.Description = request.Description
	rule.Type = generic.Http
	rule.Method = request.Method
	rule.Path = request.Path
	return rule, nil
}

func (r *Rest) fromModel(rule *HttpRule) (*Response, error) {
	response := new(Response)

	value, err := rule.Type.Value()
	if err != nil {
		return nil, err
	}

	response.ID = rule.ID.String()
	response.Name = rule.Name
	response.Description = rule.Description
	response.Type = (value).(string)
	response.Method = rule.Method
	response.Path = rule.Path

	return response, nil
}
