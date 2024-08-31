package tenant

import "go.uber.org/fx"

var Module = fx.Module("tenant",
	fx.Provide(
		fx.Annotate(
			newTenantRepository,
			fx.As(new(TenantRepository)),
		),
	),
)
