package config

import "context"

// FIXME: remove it after context received with request

func ContextConfiguration() context.Context {
	return context.Background()
}
