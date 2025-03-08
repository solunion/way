package rule

type HttpRule struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
