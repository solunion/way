package rule

import (
	"context"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/solunion/way/backend/internal/pkg/common/handlers"
	"go.uber.org/zap"
	"strings"
)

type Rest struct {
	handlers.Rest[Rule]
	service *Service
	log     *zap.SugaredLogger
}

func NewRest(service *Service, log *zap.SugaredLogger) *Rest {
	return &Rest{service: service, log: log}
}

func (r *Rest) Create(ctx fiber.Ctx) error {
	r.log.Debug("Rule - Create: API called...")

	request, err := r.buildCreateRequest(ctx)

	if err != nil {
		r.log.Error("Failed to bind body request:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rule := new(Rule)

	if err := copier.Copy(rule, request); err != nil {
		r.log.Error("Failed to build request model:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.service.Create(ctx.Context(), rule); err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := r.buildResponse(rule)

	if err != nil {
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

	rules := make([]Rule, 0)

	if err := r.service.GetAll(requestCtx, &rules); err != nil {
		r.log.Error("Failed to find all rules:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := make([]map[string]interface{}, 0)

	for _, rule := range rules {
		item, err := r.buildResponse(&rule)

		if err != nil {
			r.log.Error("Failed to build response:", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response = append(response, item)
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Rest) GetById(ctx fiber.Ctx) error {
	r.log.Debug("Rule - GetById API called...")

	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		r.log.Error("Failed to parse id:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rule := new(Rule)

	if err := r.service.GetByID(ctx.Context(), rule, id); err != nil {
		r.log.Error("Failed to find rule by id:", err)
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := r.buildResponse(rule)

	if err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Rest) Update(ctx fiber.Ctx) error {
	r.log.Debug("Rule - Update API called...")
	request, err := r.buildUpdateRequest(ctx)

	if err != nil {
		r.log.Error("Failed to bind body request:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rule := new(Rule)

	if err := copier.Copy(rule, request); err != nil {
		r.log.Error("Failed to build request model:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := r.service.Update(ctx.Context(), rule); err != nil {
		r.log.Error("Failed to create rule:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := r.buildResponse(rule)

	if err != nil {
		r.log.Error("Failed to build response:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Rest) buildCreateRequest(ctx fiber.Ctx) (*CreateRequest, error) {
	ruleType := &struct {
		Type string `json:"type"`
	}{}

	if err := ctx.Bind().Body(ruleType); err != nil {
		return nil, err
	}

	request := new(CreateRequest)
	var err error

	switch strings.ToUpper(ruleType.Type) {
	case "HTTP":
		httpReq := new(CreateHttpRequest)

		err = ctx.Bind().Body(httpReq)
		if err != nil {
			return nil, err
		}

		*request = httpReq.CreateRequest
		err = mapstructure.Decode(httpReq.HttpRuleValue, &request.Value)
	case "ROUTE":
		routeReq := new(CreateRouteRequest)

		err = ctx.Bind().Body(routeReq)
		if err != nil {
			return nil, err
		}

		*request = routeReq.CreateRequest
		err = mapstructure.Decode(routeReq.RouteRuleValue, &request.Value)
	default:
		return nil, fmt.Errorf("unknown rule type '%s'", ruleType)
	}

	return request, err
}

func (r *Rest) buildUpdateRequest(ctx fiber.Ctx) (*UpdateRequest, error) {
	var err error

	id := ctx.Params("id")

	ruleType := &struct {
		Type string `json:"type"`
	}{}

	if err := ctx.Bind().Body(ruleType); err != nil {
		return nil, err
	}

	request := new(UpdateRequest)

	switch strings.ToUpper(ruleType.Type) {
	case "HTTP":
		httpReq := new(UpdateHttpRequest)

		err = ctx.Bind().Body(httpReq)
		if err != nil {
			return nil, err
		}

		*request = httpReq.UpdateRequest
		err = mapstructure.Decode(httpReq.HttpRuleValue, &request.Value)
	case "ROUTE":
		routeReq := new(UpdateRouteRequest)

		err = ctx.Bind().Body(routeReq)
		if err != nil {
			return nil, err
		}

		*request = routeReq.UpdateRequest
		err = mapstructure.Decode(routeReq.RouteRuleValue, &request.Value)
	default:
		return nil, fmt.Errorf("unknown rule type '%s'", ruleType)
	}

	request.ID = id

	return request, err
}

func (r *Rest) buildResponse(rule *Rule) (map[string]interface{}, error) {
	mapResult := make(map[string]interface{})

	common := Common{
		ID:          rule.ID.String(),
		Type:        rule.Type.String(),
		Name:        rule.Name,
		Description: rule.Description,
	}

	var err error

	switch rule.Type {
	case Http:
		val := new(HttpRuleValue)
		err = mapstructure.Decode(rule.Value, val)

		if err != nil {
			return nil, err
		}
		httpResponse := &HttpResponse{
			Response: Response{common},
			Value: HttpRuleValue{
				Method: val.Method,
				Path:   val.Path,
			},
		}

		err = mapstructure.Decode(httpResponse, &mapResult)
	case Route:
		val := new(RouteRuleValue)
		err = mapstructure.Decode(rule.Value, val)
		if err != nil {
			return nil, err
		}
		routeResponse := &RouteResponse{
			Response: Response{common},
			Value: RouteRuleValue{
				Path: val.Path,
			},
		}
		err = mapstructure.Decode(routeResponse, &mapResult)
	default:
		return nil, fmt.Errorf("unhandled rule type '%s'", rule.Type)
	}

	return mapResult, err
}
