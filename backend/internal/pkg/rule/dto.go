package rule

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
	Value       string `json:"value"`
	Type        string `json:"type"`
}

type ResponseDto struct {
	Dto
}
