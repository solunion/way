package http

import (
	"encoding/json"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
)

type HttpRuleValue struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
type HttpRule struct {
	generic.Rule
	HttpRuleValue
}

func (r *HttpRule) Type() generic.Type {
	return generic.Http
}

func (r *HttpRule) Value() json.RawMessage {
	value := &HttpRuleValue{
		Method: r.Method,
		Path:   r.Path,
	}

	if raw, err := json.Marshal(value); err == nil {
		return raw
	}

	return nil
}
