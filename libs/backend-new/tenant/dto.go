package tenant

type CreateRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResponseDto struct {
	CreateRequestDto
	ID string `json:"id"`
}
