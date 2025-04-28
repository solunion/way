package rule

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/http"
	"strings"

	"go.uber.org/zap"
)

type Rest struct {
	handlers.Rest[generic.Rule[any]]
	log      *zap.SugaredLogger
	service  *Service
	httpRest *http.Rest
}

func NewRest(log *zap.SugaredLogger, service *Service, httpRest *http.Rest) *Rest {
	return &Rest{log: log, service: service, httpRest: httpRest}
}

func (r *Rest) Create(ctx fiber.Ctx) error {
	r.log.Debug("Rule - Create: API called...")

	request := &struct {
		Type string `json:"type"`
	}{}

	if err := ctx.Bind().Body(request); err != nil {
		r.log.Error("Failed to bind body request type:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	switch strings.ToLower(request.Type) {
	case "http":
		return r.httpRest.Create(ctx)
	default:
		r.log.Error("Unsupported type:", request.Type)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported type"})
	}
}

func (r *Rest) GetAll(ctx fiber.Ctx) error {
	r.log.Debug("Rule - GetAll API called...")

	ruleType := ctx.Params("type")
	r.log.Debugf("Optional rule type: %q", ruleType)

	switch strings.ToLower(ruleType) {
	case "":
		roles := make([]generic.Rule[any], 0)

		if err := r.service.GetAll(ctx.Context(), &roles); err != nil {
			r.log.Error("Failed to find all rules:", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response := make([]generic.Response, 0)

		if err := copier.Copy(&response, roles); err != nil {
			r.log.Error("Failed to build response:", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	case "http":
		return r.httpRest.GetAll(ctx)
	default:
		r.log.Error("Unsupported type:", ruleType)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported type"})
	}
}
