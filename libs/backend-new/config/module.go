package config

import "go.uber.org/fx"

var Module = fx.Module("config",
	fx.Provide(ViperConfiguration),
	fx.Provide(DatabaseConnection),
	fx.Provide(LoggerConfiguration),
)
