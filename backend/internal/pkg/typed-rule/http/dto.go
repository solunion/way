package http

type CreateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Method      string  `json:"method"`
	Path        string  `json:"path"`
}

type Response struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Method      string  `json:"method"`
	Path        string  `json:"path"`
}
