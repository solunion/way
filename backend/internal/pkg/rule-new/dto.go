package rule_new

import "encoding/json"

type CreateRequest struct {
	ID          string          `json:"id"`
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Value       json.RawMessage `json:"value"`
}

type Response struct {
	ID          string          `json:"id"`
	Type        Type            `json:"type"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Value       json.RawMessage `json:"value"`
}
