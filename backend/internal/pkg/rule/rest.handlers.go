package rule

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"go.uber.org/zap"
)

type Rest struct {
	handlers.Rest[any]
	service *Service
	log     *zap.SugaredLogger
}

func NewRest(service *Service, log *zap.SugaredLogger) *Rest {
	return &Rest{service: service, log: log}
}

func (r *Rest) Create(ctx fiber.Ctx) error {
	r.log.Debug("Rule - Create: API called...")

	request, err := buildRequest(ctx)

	if err != nil {
		r.log.Error("Failed to bind body request:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rule := new(Rule[any])

	if err := copier.Copy(rule, request); err != nil {
		r.log.Error("Failed to build request model:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.service.Create(ctx.Context(), rule); err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := buildResponse(rule)

	if err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := copier.Copy(response, rule); err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Rest) GetAll(ctx fiber.Ctx) error {
	r.log.Debug("Rule - GetAll API called...")

	ruleType := ctx.Params("type")
	r.log.Debugf("Optional rule type: %q", ruleType)

	requestCtx := ctx.Context()

	if ruleType != "" {
		requestCtx = context.WithValue(requestCtx, "rule_type", ruleType)
	}

	rules := make([]Rule[any], 0)

	if err := r.service.GetAll(requestCtx, &rules); err != nil {
		r.log.Error("Failed to find all rules:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := make([]any, 0)

	for _, rule := range rules {
		item, err := buildResponse(&rule)

		if err != nil {
			r.log.Error("Failed to build response:", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response = append(response, item)
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
