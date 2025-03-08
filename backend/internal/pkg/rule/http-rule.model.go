package rule

type HttpRule struct {
	BaseRule
	Method string `json:"method"`
	Path   string `json:"path"`
}
