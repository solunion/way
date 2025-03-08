package rule

import (
	"encoding/json"
	"fmt"
)

//type Rule interface {
//  Type() Type
//  ToValue() []byte
//}

type GenericRule[T HttpRule | RouteRule] struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Value       T
	Type        Type
}

type HttpRule struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (r *HttpRule) Type() Type {
	return Http
}

func (r *HttpRule) ToValue() []byte {
	var value = map[string]interface{}{
		"method": r.Method,
		"path":   r.Path,
	}

	result, err := json.Marshal(value)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func (r *HttpRule) FromValue(value []byte) error {
	return json.Unmarshal(value, r)
}

type RouteRule struct {
	Prova string `json:"path"`
}

func FromEntity[T HttpRule | RouteRule](entity RuleDao, rule *GenericRule[T]) error {
	var err error

	rule.ID = entity.ID.String()
	rule.Name = entity.Name
	rule.Description = entity.Description
	rule.Type = entity.Type

	switch entity.Type {
	case Http:
		var value = new(HttpRule)
		err = json.Unmarshal(entity.Value, value)
	case Route:
		var value = new(RouteRule)
		err = json.Unmarshal(entity.Value, value)
	}

	return err
}
