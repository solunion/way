package rule

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"go.uber.org/zap"
)

type Rest[T HttpRule | RouteRule] struct {
	handlers.Rest[RuleDao]
	service *Service[T]
	log     *zap.SugaredLogger
}

func newHttpRest[T HttpRule | RouteRule](service *Service[T], log *zap.SugaredLogger) *Rest[T] {
	return &Rest[T]{service: service, log: log}
}

func (r *Rest[T]) Create(ctx fiber.Ctx) error {
	r.log.Debug("HttpRule - Create API called...")

	request := new(CreateRequestDto)

	if err := ctx.Bind().Body(request); err != nil {
		r.log.Error("Failed to bind body request:", err)
		return err
	}

	rule := new(GenericRule[T])

	if err := copier.Copy(rule, request); err != nil {
		r.log.Error("Failed to convert rule DTO:", err)
		return err
	}

	if err := r.service.Create(ctx.Context(), rule); err != nil {
		r.log.Error("Failed to create rule:", err)
		return err
	}

	response := new(ResponseDto)

	if err := copier.Copy(response, rule); err != nil {
		r.log.Error("Failed to build response:", err)
		return err
	}

	return ctx.JSON(response)
}

func (r *Rest[T]) GetAll(ctx fiber.Ctx) error {
	r.log.Debug("HttpRule - GetAll API called...")

	rules := make([]GenericRule[T], 0)

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
