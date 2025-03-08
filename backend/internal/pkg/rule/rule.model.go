package rule

import (
	"encoding/json"
	"fmt"
)

type Rule interface {
	Type() string
	Value() GenericRule[map[string]interface{}]
}

type GenericRule[T HttpRule | RouteRule | map[string]interface{}] struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   *string `json:"description"`
	InternalValue T       `json:"value"`
	InternalType  Type    `json:"type"`
}

func (r GenericRule[T]) Type() string {
	return r.InternalType.String()
}

func (r GenericRule[T]) Value() GenericRule[map[string]interface{}] {
	var finalValue map[string]interface{}
	value, ok := any(r.InternalValue).(map[string]interface{})
	if !ok {
		// FIXME: handle it
		fmt.Printf("Failed to convert finalValue to map[string]interface{} for ID: %s\n", r.ID)
	}

	switch r.InternalType {
	case Http:
		fmt.Println("Convert finalValue to HttpRule")
		finalValue = map[string]interface{}{
			"method": value["method"],
			"path":   value["path"],
		}
	case Route:
		fmt.Println("Convert finalValue to RouteRule")
		finalValue = map[string]interface{}{
			"route": value["route"],
		}
	}

	result := GenericRule[map[string]interface{}]{
		ID:            r.ID,
		Name:          r.Name,
		Description:   r.Description,
		InternalValue: finalValue,
		InternalType:  r.InternalType,
	}

	return result
}

func FromEntity[T map[string]interface{}](entity RuleDao, rule *GenericRule[T]) error {
	var err error

	rule.ID = entity.ID.String()
	rule.Name = entity.Name
	rule.Description = entity.Description
	rule.InternalType = entity.Type
	fmt.Println("Inside FromEntity: entity.InternalValue: ", string(entity.Value))

	err = json.Unmarshal(entity.Value, &rule.InternalValue)

	fmt.Println("Inside FromEntity: rule.InternalValue: ", rule.InternalValue)

	return err
}
