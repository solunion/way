package rule

import (
	"encoding/json"
	"fmt"
)

// HttpHandler implementa RuleHandler per le regole HTTP
type HttpHandler struct{}

// NewHttpHandler crea una nuova istanza di HttpHandler
func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

// Type restituisce il tipo di regola gestito da questo handler
func (h *HttpHandler) Type() Type {
	return Http
}

// ToEntity converte i dettagli JSON in un'entità HttpRule
func (h *HttpHandler) ToEntity(data json.RawMessage, common CommonRuleDTO) (Rule, error) {
	var details HttpRuleDetails
	if err := json.Unmarshal(data, &details); err != nil {
		return nil, fmt.Errorf("invalid HTTP rule details: %w", err)
	}

	description := common.Description
	return &HttpRule{
		BaseRule: BaseRule{
			ID:           common.ID,
			Name:         common.Name,
			Description:  &description,
			Type:         Http,
			TypeInString: Http.String(),
		},
		Method: details.Method,
		Path:   details.Path,
	}, nil
}

// ToDTO converte un'entità HttpRule in dettagli JSON
func (h *HttpHandler) ToDTO(rule Rule) (json.RawMessage, error) {
	httpRule, ok := rule.(*HttpRule)
	if !ok {
		return nil, fmt.Errorf("invalid rule type for HTTP handler: expected *HttpRule")
	}

	details := HttpRuleDetails{
		Method: httpRule.Method,
		Path:   httpRule.Path,
	}

	return json.Marshal(details)
}

// Validate valida i dettagli JSON per una regola HTTP
func (h *HttpHandler) Validate(data json.RawMessage) error {
	var details HttpRuleDetails
	if err := json.Unmarshal(data, &details); err != nil {
		return fmt.Errorf("invalid HTTP rule details: %w", err)
	}

	if details.Method == "" {
		return fmt.Errorf("HTTP rule requires a method")
	}

	if details.Path == "" {
		return fmt.Errorf("HTTP rule requires a path")
	}

	return nil
}
