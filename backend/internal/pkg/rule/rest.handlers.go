package rule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/gofiber/fiber/v3"
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

	request, err := r.buildRequest(ctx)

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

	response := make([]RuleResponse, 0)

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

func (r *Rest) buildRequest(ctx fiber.Ctx) (*CreateRequest, error) {
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

func (r *Rest) buildResponse(rule *Rule) (RuleResponse, error) {
	var response RuleResponse

	common := Response{
		ID:          rule.ID.String(),
		Type:        rule.Type.String(),
		Name:        rule.Name,
		Description: rule.Description,
	}

	switch rule.Type {
	case Http:
		val := new(HttpRuleValue)
		err := mapstructure.Decode(rule.Value, val)
		if err != nil {
			return nil, err
		}
		response = &HttpResponse{
			Response: common,
			HttpRuleValue: HttpRuleValue{
				Method: val.Method,
				Path:   val.Path,
			},
		}
	case Route:
		val := new(RouteRuleValue)
		err := mapstructure.Decode(rule.Value, val)
		if err != nil {
			return nil, err
		}
		response = &RouteResponse{
			Response: common,
			RouteRuleValue: RouteRuleValue{
				Path: val.Path,
			},
		}
	default:
		return nil, errors.New("unhandled rule type")
	}

	return response, copier.Copy(response, rule)
}
