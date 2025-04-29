package rule

import (
	"errors"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"strings"
)

func buildRequest(ruleType string) (any, error) {
	switch strings.ToUpper(ruleType) {
	case "HTTP":
		return new(CreateHttpRequest), nil
	case "ROUTE":
		return new(CreateRouteRequest), nil
	default:
		return nil, fmt.Errorf("unknown rule type '%s'", ruleType)
	}
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
