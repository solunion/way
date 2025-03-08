package rule

type RouteRule struct {
	BaseRule
	Path string `json:"route"`
}
