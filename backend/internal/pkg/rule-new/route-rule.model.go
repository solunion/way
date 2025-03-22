package rule_new

import "encoding/json"

type RouteRuleValue struct {
	Route string `json:"route"`
}

type RouteRule struct {
	BasicRule
	RouteRuleValue
}

func (r *RouteRule) GetType() Type {
	return Http
}

func (r *RouteRule) GetValue() (json.RawMessage, error) {
	value := RouteRuleValue{
		Route: r.Route,
	}
	return json.Marshal(value)
}
