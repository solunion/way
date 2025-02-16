package rule

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
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

func (r *Rest) GetAll(ctx fiber.Ctx) error {
	r.log.Debug("Rule - GetAll API called...")

	rules := make([]Rule, 0)

	if err := r.service.GetAll(ctx.Context(), &rules); err != nil {
		r.log.Error("Failed to find all rules:", err)
	}

	r.log.Debug("Found rules: %+v", rules)

	response := make([]ResponseDto, 0)

	if err := copier.Copy(&response, rules); err != nil {
		r.log.Error("Failed to build response:", err)
	}

	return ctx.JSON(response)
}
