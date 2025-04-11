package http

import "github.com/solunion/way/backend/internal/pkg/typed-rule/generic"

type HttpRule struct {
	generic.Generic
	Method string
	Path   string
}
