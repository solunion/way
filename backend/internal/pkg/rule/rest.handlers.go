package rule

import (
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"go.uber.org/zap"
)

type Rest struct {
	handlers.Rest[Rule]
	service *Service
	log     *zap.SugaredLogger
}

func newHttpRest(service *Service, log *zap.SugaredLogger) *Rest {
	return &Rest{service: service, log: log}
}

func (r *Rest) Create(ctx fiber.Ctx) error {
	r.log.Debug("Rule - Create API called...")

	// Utilizziamo il nuovo DTO CreateRuleRequest
	request := new(CreateRuleRequest)

	if err := ctx.Bind().Body(request); err != nil {
		r.log.Error("Failed to bind body request:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Utilizziamo il nuovo metodo CreateFromRequest che gestisce tutti i tipi di regole
	response, err := r.service.CreateFromRequest(ctx.Context(), *request)
	if err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Rest) GetAll(ctx fiber.Ctx) error {
	r.log.Debug("HttpRule - GetAll API called...")

	var rules []Rule

	if result, err := r.service.GetAll(ctx.Context()); err != nil {
		r.log.Error("Failed to find all rules:", err)
	} else {
		rules = result
	}

	r.log.Debug("Found rules: %+v", rules)

	return ctx.JSON(rules)
}
