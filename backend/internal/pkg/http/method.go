package http

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	PATCH   HttpMethod = "PATCH"
	DELETE  HttpMethod = "DELETE"
	OPTIONS HttpMethod = "OPTIONS"
)
