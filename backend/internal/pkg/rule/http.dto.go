package rule

type HttpRuleValue struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (v HttpRuleValue) GetRuleType() Type {
	return Http
}

type CreateHttpRequest struct {
	CreateRequest
	HttpRuleValue
}

type HttpResponse struct {
	Response
	HttpRuleValue
}
