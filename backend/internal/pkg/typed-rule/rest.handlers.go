package rule

import (
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/http"
	"strings"

	"go.uber.org/zap"
)

type Rest struct {
	handlers.Rest[generic.Rule]
	log      *zap.SugaredLogger
	httpRest *http.Rest
}

func NewRest(log *zap.SugaredLogger, httpRest *http.Rest) *Rest {
	return &Rest{log: log, httpRest: httpRest}
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
