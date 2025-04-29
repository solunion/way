package rule

import (
	"encoding/json"
)

type HttpRuleValue struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type CreateHttpRequest struct {
	CreateRequest
	HttpRuleValue
}

func (r *CreateHttpRequest) Value() (json.RawMessage, error) {
	return json.Marshal(r.HttpRuleValue)
}

type HttpResponse struct {
	Response
	HttpRuleValue
}
