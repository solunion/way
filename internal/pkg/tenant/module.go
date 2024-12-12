package tenant

import (
	"go.uber.org/fx"
)

var Module = fx.Module("tenant",
	fx.Provide(
		newTenantRepository,
	),
	fx.Invoke(
		registerTenantHttp,
	),
)
