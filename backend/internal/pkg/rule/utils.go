package rule

import (
	"fmt"
	"github.com/go-viper/mapstructure/v2"
)

func GetTypedValue[T RuleValues](rule *Rule, value *T) error {
	if value == nil {
		value = new(T)
	}

	deferenceValue := *value

	if rule.Type != deferenceValue.GetRuleType() {
		return fmt.Errorf("rule type mismatch. expected: %s, got: %s", deferenceValue.GetRuleType(), rule.Type)
	}

	return mapstructure.Decode(rule.Value, &value)
}
