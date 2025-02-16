package rule

import "encoding/json"

type CreateRequestDto struct {
	Dto
	ID string `json:"-"`
}

type UpdateRequestDto struct {
	Dto
}

type Dto struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
	Value       json.RawMessage `json:"value"`
}

type ResponseDto struct {
	Dto
}
