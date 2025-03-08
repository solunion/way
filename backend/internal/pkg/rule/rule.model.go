package rule

import (
	"encoding/json"
	"fmt"
)

type Rule interface {
	GetType() Type
	GetValue() GenericRule[map[string]interface{}]
}

type GenericRule[T HttpRule | RouteRule | map[string]interface{}] struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Value       T       `json:"value"`
	Type        Type    `json:"type"`
}

type HttpRule struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (r GenericRule[T]) GetType() Type {
	return r.Type
}

func (r GenericRule[T]) GetValue() GenericRule[map[string]interface{}] {
	var value map[string]interface{}

	switch r.Type {
	case Http:
		fmt.Println("Convert value to HttpRule")
		http, ok := any(r.Value).(map[string]interface{})
		if !ok {
			// FIXME: handle it
			fmt.Printf("Failed to convert value to HttpRule for ID: %s\n", r.ID)
		}
		value = map[string]interface{}{
			"method": http["method"],
			"path":   http["path"],
		}
	case Route:
		fmt.Println("Convert value to RouteRule")
		route, ok := any(r.Value).(map[string]interface{})
		if !ok {
			// FIXME: handle it
			fmt.Printf("Failed to convert value to RouteRule for ID: %s\n", r.ID)
		}
		fmt.Printf("Route: %s\n", route["route"])

		value = map[string]interface{}{
			"route": route["route"],
		}
	}

	result := GenericRule[map[string]interface{}]{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Value:       value,
		Type:        r.Type,
	}

	return result
}

//func (r *HttpRule) FromValue(value []byte) error {
//	return json.Unmarshal(value, r)
//}

type RouteRule struct {
	Route string `json:"route"`
}

func FromEntity[T map[string]interface{}](entity RuleDao, rule *GenericRule[T]) error {
	var err error

	rule.ID = entity.ID.String()
	rule.Name = entity.Name
	rule.Description = entity.Description
	rule.Type = entity.Type
	fmt.Println("Inside FromEntity: entity.Value: ", string(entity.Value))

	err = json.Unmarshal(entity.Value, &rule.Value)

	fmt.Println("Inside FromEntity: rule.Value: ", rule.Value)

	return err
}
