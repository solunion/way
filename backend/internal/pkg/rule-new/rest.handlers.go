package rule_new

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"go.uber.org/zap"
)

type Rest struct {
	handlers.Rest[Rule]
	service *Service
	log     *zap.SugaredLogger
}

func newRest(service *Service, log *zap.SugaredLogger) *Rest {
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
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := r.service.Create(ctx.Context(), rule)

	if err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := r.fromModel(created)

	if err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Rest) toModel(request *CreateRequest) (Rule, error) {
	rule := new(BasicRule)

	if ruleType, err := TypeFromString(request.Type); err != nil {
		return nil, err
	} else {
		if request.ID != "" {
			rule.ID = uuid.MustParse(request.ID)
		}
		rule.Name = request.Name
		rule.Description = request.Description
		rule.Type = ruleType
		rule.Value = request.Value
	}

	return rule, nil
}

func (r *Rest) fromModel(rule Rule) (*Response, error) {
	response := new(Response)

	if value, err := rule.GetValue(); err != nil {
		return nil, err
	} else {
		response.Name = rule.GetName()
		response.Description = rule.GetDescription()
		response.Type = rule.GetType()
		response.Value = value
	}

	return response, nil
}
