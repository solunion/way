package profile

import "go.uber.org/fx"

var Module = fx.Module("profile",
	fx.Provide(
		fx.Annotate(
			newProfileRepository,
			fx.As(new(ProfileRepository)),
		),
	),
)
