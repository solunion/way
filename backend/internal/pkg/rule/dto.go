package rule

import "encoding/json"

// CommonRuleDTO contiene i campi comuni a tutti i tipi di rule
type CommonRuleDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// RuleDetails è un DTO generico che può contenere qualsiasi tipo di dettagli specifici
type RuleDetails[T any] struct {
	Type    Type          `json:"type"`
	Common  CommonRuleDTO `json:"common"`
	Details T             `json:"details"`
}

// HttpRuleDetails contiene i dettagli specifici per le regole HTTP
type HttpRuleDetails struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

// RouteRuleDetails contiene i dettagli specifici per le regole Route
type RouteRuleDetails struct {
	RouteName string `json:"routeName"`
	Priority  int    `json:"priority"`
}

// CreateRuleRequest è il DTO utilizzato per creare una nuova regola
type CreateRuleRequest struct {
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Details     json.RawMessage `json:"details"`
}

// UpdateRuleRequest è il DTO utilizzato per aggiornare una regola esistente
type UpdateRuleRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Details     json.RawMessage `json:"details"`
}

// RuleResponse è il DTO utilizzato per restituire una regola
type RuleResponse struct {
	ID          string          `json:"id"`
	Type        Type            `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Details     json.RawMessage `json:"details"`
}
