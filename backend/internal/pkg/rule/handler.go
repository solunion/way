package rule

import "encoding/json"

// RuleHandler definisce l'interfaccia per i gestori di regole specifiche per tipo
type RuleHandler interface {
	// Type restituisce il tipo di regola gestito da questo handler
	Type() Type

	// ToEntity converte i dettagli JSON in un'entità Rule
	ToEntity(details json.RawMessage, common CommonRuleDTO) (Rule, error)

	// ToDTO converte un'entità Rule in dettagli JSON
	ToDTO(rule Rule) (json.RawMessage, error)

	// Validate valida i dettagli JSON per questo tipo di regola
	Validate(details json.RawMessage) error
}
