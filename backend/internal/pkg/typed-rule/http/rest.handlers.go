package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
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

	rule := new(HttpRule)

	if err := copier.Copy(rule, request); err != nil {
		r.log.Error("Failed to build request model:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.service.Create(ctx.Context(), rule); err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := new(Response)

	if err := copier.Copy(response, rule); err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
