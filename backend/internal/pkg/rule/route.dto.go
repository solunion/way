package rule

type RouteRuleValue struct {
	Path string `json:"path"`
}

type CreateRouteRequest struct {
	CreateRequest
	RouteRuleValue
}

type RouteResponse struct {
	Response
	RouteRuleValue
}
