package tenant

type CreateRequestDto struct {
	Dto
	ID string `json:"-"`
}

type UpdateRequestDto struct {
	Dto
}

type Dto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResponseDto struct {
	Dto
}
