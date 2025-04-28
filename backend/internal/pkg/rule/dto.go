package rule

import "encoding/json"

type CreateRequest struct {
	Response
	ID    string `json:"-"`
	Value json.RawMessage
}
type Response struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
