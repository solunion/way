package http

import (
	"encoding/json"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
)

type HttpRule struct {
	generic.Rule
	Method string
	Path   string
}

func (r *HttpRule) Type() generic.Type {
	return generic.Http
}

func (r *HttpRule) Value() json.RawMessage {
	value := &struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}{
		Method: r.Method,
		Path:   r.Path,
	}

	if raw, err := json.Marshal(value); err == nil {
		return raw
	}

	return nil
}
