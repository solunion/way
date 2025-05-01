package rule

type HttpRuleValue struct {
	Method string `mapstructure:"method" json:"method"`
	Path   string `mapstructure:"path" json:"path"`
}

func (v HttpRuleValue) GetRuleType() Type {
	return Http
}

type CreateHttpRequest struct {
	CreateRequest
	HttpRuleValue
}

type UpdateHttpRequest struct {
	UpdateRequest
	HttpRuleValue
}

type HttpResponse struct {
	Response Response      `mapstructure:",squash"`
	Value    HttpRuleValue `mapstructure:",squash"`
}
