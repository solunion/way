package rule

import (
	"encoding/json"
)

type RouteRuleValue struct {
	Path string `json:"path"`
}

type CreateRouteRequest struct {
	CreateRequest
	RouteRuleValue
}

func (r *CreateRouteRequest) Value() (json.RawMessage, error) {
	return json.Marshal(r.RouteRuleValue)
}

type RouteResponse struct {
	Response
	RouteRuleValue
}
