package rule

import (
	"encoding/json"
	"fmt"
)

// RouteHandler implementa RuleHandler per le regole Route
type RouteHandler struct{}

// NewRouteHandler crea una nuova istanza di RouteHandler
func NewRouteHandler() (RuleHandler, *RouteRuleHandlerKey) {
	return &RouteHandler{}, &RouteRuleHandlerKey{}
}

// Type restituisce il tipo di regola gestito da questo handler
func (h *RouteHandler) Type() Type {
	return Route
}

// ToEntity converte i dettagli JSON in un'entità RouteRule
func (h *RouteHandler) ToEntity(data json.RawMessage, common CommonRuleDTO) (Rule, error) {
	var details RouteRuleDetails
	if err := json.Unmarshal(data, &details); err != nil {
		return nil, fmt.Errorf("invalid Route rule details: %w", err)
	}

	// Converti i dettagli in un'entità RouteRule
	description := common.Description
	return &RouteRule{
		BaseRule: BaseRule{
			ID:           common.ID,
			Name:         common.Name,
			Description:  &description,
			Type:         Route,
			TypeInString: Route.String(),
		},
		Path: details.RouteName, // Utilizziamo RouteName come Path
	}, nil
}

// ToDTO converte un'entità RouteRule in dettagli JSON
func (h *RouteHandler) ToDTO(rule Rule) (json.RawMessage, error) {
	routeRule, ok := rule.(*RouteRule)
	if !ok {
		return nil, fmt.Errorf("invalid rule type for Route handler: expected *RouteRule")
	}

	details := RouteRuleDetails{
		RouteName: routeRule.Path,
		Priority:  0, // Valore di default, potrebbe essere esteso nel modello
	}

	return json.Marshal(details)
}

// Validate valida i dettagli JSON per una regola Route
func (h *RouteHandler) Validate(data json.RawMessage) error {
	var details RouteRuleDetails
	if err := json.Unmarshal(data, &details); err != nil {
		return fmt.Errorf("invalid Route rule details: %w", err)
	}

	if details.RouteName == "" {
		return fmt.Errorf("Route rule requires a route name")
	}

	return nil
}
