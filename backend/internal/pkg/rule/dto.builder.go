package rule

import (
	"encoding/json"
	"errors"
	"fmt"
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
	response := new(Response)

	if marshal, err := json.Marshal(rule); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(marshal, response); err != nil {
			return nil, err
		}

		if marshalValue, err := json.Marshal(rule.Value); err != nil {
			return nil, err
		} else {
			switch rule.Type {
			case Http:
				http := new(HttpResponse)
				http.Response = *response
				if err := json.Unmarshal(marshalValue, &http.HttpRuleValue); err != nil {
					return nil, err
				}
				return http, nil

			case Route:
				route := new(RouteResponse)
				route.Response = *response
				if err := json.Unmarshal(marshalValue, &route.RouteRuleValue); err != nil {
					return nil, err
				}
				return route, nil
			default:
				return nil, errors.New("unknown rule type")
			}
		}
	}
}
