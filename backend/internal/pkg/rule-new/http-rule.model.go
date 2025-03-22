package rule_new

import (
	"encoding/json"
)

type HttpRuleValue struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type HttpRule struct {
	BasicRule
	HttpRuleValue
}

func (r *HttpRule) GetType() Type {
	return Http
}

func (r *HttpRule) GetValue() (json.RawMessage, error) {
	value := HttpRuleValue{
		Method: r.Method,
		Path:   r.Path,
	}
	return json.Marshal(value)
}

func (r *HttpRule) Validate() error {
	// TODO: implement me
	value := new(HttpRuleValue)
	return json.Unmarshal(r.Value, &value)
}
