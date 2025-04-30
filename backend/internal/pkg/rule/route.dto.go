package rule

type RouteRuleValue struct {
	Path string `json:"path"`
}

func (v RouteRuleValue) GetRuleType() Type {
	return Route
}

type CreateRouteRequest struct {
	CreateRequest
	RouteRuleValue
}

type RouteResponse struct {
	Response Response       `mapstructure:",squash"`
	Value    RouteRuleValue `mapstructure:",squash"`
}
