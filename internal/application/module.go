package application

import "go.uber.org/fx"

var Module = fx.Module("application",
	fx.Provide(
		fx.Annotate(
			newApplicationRepository,
			fx.As(new(ApplicationRepository)),
		),
	),
)
