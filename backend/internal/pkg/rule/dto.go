package rule

type CreateRequest struct {
	Response
	ID    string                 `json:"-"`
	Value map[string]interface{} `json:"value"`
}

type Response struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
