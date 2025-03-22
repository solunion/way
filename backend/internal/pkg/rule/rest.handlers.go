package rule

import (
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"go.uber.org/zap"
)

type Rest struct {
	handlers.Rest[Rule]
	service            *Service
	log                *zap.SugaredLogger
	typeHandlerService *TypeHandlerService
}

func newHttpRest(service *Service, log *zap.SugaredLogger, typeHandlerService *TypeHandlerService) *Rest {
	return &Rest{service: service, log: log, typeHandlerService: typeHandlerService}
}

func (r *Rest) Create(ctx fiber.Ctx) error {
	r.log.Debug("Rule - Create API called...")

	// Utilizziamo il nuovo DTO CreateRuleRequest
	request := new(CreateRuleRequest)

	if err := ctx.Bind().Body(request); err != nil {
		r.log.Error("Failed to bind body request:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rule, err := r.typeHandlerService.CreateRuleFromRequest(ctx.Context(), *request)

	if err != nil {
		r.log.Error("Failed to create rule from request:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Utilizziamo il nuovo metodo CreateFromRequest che gestisce tutti i tipi di regole
	created, err := r.service.Create(ctx.Context(), rule)

	if err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	handler, exists := r.typeHandlerService.GetHandler(created.GetType())
	if !exists {
		r.log.Error("Failed to create rule from request:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	details, err := handler.ToDTO(created)
	if err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Prepara la risposta
	getType := created.GetType()

	response := RuleResponse{
		ID:          created.GetInfo().ID,
		Type:        getType.String(),
		Name:        created.GetInfo().Name,
		Description: *created.GetInfo().Description,
		Details:     details,
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
