package http

import (
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
)

type HttpRuleValue struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
type HttpRule struct {
	generic.Rule[HttpRuleValue]
}

func (r *HttpRule) Type() generic.Type {
	return generic.Http
}

//func (r *HttpRule) FromModel(rule *generic.Rule) error {
//	r.Rule = *rule
//
//	if err := json.Unmarshal(rule.Value, &r.HttpRuleValue); err != nil {
//		return err
//	}
//
//	return nil
//}
