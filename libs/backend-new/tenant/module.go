package tenant

import (
	"go.uber.org/fx"
)

var Module = fx.Module("tenant",
	fx.Provide(
		newRepository,
		newService,
	),
	fx.Invoke(
		newRest,
	),
)
