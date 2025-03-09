package rule

type CreateHttpRequestDto struct {
	Dto
	ID     string `json:"-"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

type CreateRequestDto struct {
	Type string `json:"type"`
}

type UpdateRequestDto struct {
	Dto
}

type Dto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Value       string `json:"value"`
	Type string `json:"type"`
}

type ResponseDto struct {
	Dto
}
