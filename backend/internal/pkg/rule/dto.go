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
	Method      string `json:"method"`
	Path        string `json:"path"`
}

type ResponseDto struct {
	Dto
}
