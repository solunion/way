package rule

import (
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/gofiber/fiber/v3"
	"strings"
)

func buildRequest(ctx fiber.Ctx) (any, error) {
	ruleType := &struct {
		Type string `json:"type"`
	}{}

	if err := ctx.Bind().Body(ruleType); err != nil {
		return nil, err
	}

	var request any

	switch strings.ToUpper(ruleType.Type) {
	case "HTTP":
		request = new(CreateHttpRequest)
	case "ROUTE":
		request = new(CreateRouteRequest)
	default:
		return nil, fmt.Errorf("unknown rule type '%s'", ruleType)
	}

	return request, ctx.Bind().Body(request)
}

func buildResponse(rule *Rule[any]) (any, error) {
	common := Response{
		ID:          rule.ID.String(),
		Type:        rule.Type.String(),
		Name:        rule.Name,
		Description: rule.Description,
	}

	switch rule.Type {
	case Http:
		var val HttpRuleValue
		err := mapstructure.Decode(rule.Value, &val)
		if err != nil {
			return nil, err
		}
		return &HttpResponse{
			Response: common,
			HttpRuleValue: HttpRuleValue{
				Method: val.Method,
				Path:   val.Path,
			},
		}, nil
	case Route:
		var val RouteRuleValue
		err := mapstructure.Decode(rule.Value, &val)
		if err != nil {
			return nil, err
		}
		return &RouteResponse{
			Response: common,
			RouteRuleValue: RouteRuleValue{
				Path: val.Path,
			},
		}, nil
	default:
		return nil, errors.New("unhandled rule type")
	}
}
