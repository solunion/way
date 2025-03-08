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

//func (r *Rest) Create(ctx fiber.Ctx) error {
//	r.log.Debug("HttpRule - Create API called...")
//
//	request := new(CreateRequestDto)
//
//	if err := ctx.Bind().Body(request); err != nil {
//		r.log.Error("Failed to bind body request:", err)
//		return err
//	}
//
//	rule := new(BaseRule[T])
//
//	if err := copier.Copy(rule, request); err != nil {
//		r.log.Error("Failed to convert rule DTO:", err)
//		return err
//	}
//
//	if err := r.service.Create(ctx.Context(), rule); err != nil {
//		r.log.Error("Failed to create rule:", err)
//		return err
//	}
//
//	response := new(ResponseDto)
//
//	if err := copier.Copy(response, rule); err != nil {
//		r.log.Error("Failed to build response:", err)
//		return err
//	}
//
//	return ctx.JSON(response)
//}

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
