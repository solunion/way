package rule

type HttpRuleValue struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type CreateHttpRequest struct {
	CreateRequest
	HttpRuleValue
}

func (r *CreateHttpRequest) Value() HttpRuleValue {
	return HttpRuleValue{
		Method: r.Method,
		Path:   r.Path,
	}
}

type HttpResponse struct {
	Response
	HttpRuleValue
}
