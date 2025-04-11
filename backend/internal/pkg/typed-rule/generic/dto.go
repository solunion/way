package generic

type Response struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
