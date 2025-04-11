package http

import "github.com/solunion/way/backend/internal/pkg/typed-rule/generic"

type CreateRequest struct {
	Response
	ID string `json:"-"`
}

type Response struct {
	generic.Response
	Method string `json:"method"`
	Path   string `json:"path"`
}
