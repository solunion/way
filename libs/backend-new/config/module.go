package config

import "go.uber.org/fx"

var Module = fx.Module("config",
	fx.Provide(ContextConfiguration),
	fx.Provide(EnvironmentConfiguration),
	fx.Provide(DatabaseConnection),
	fx.Provide(LoggerConfiguration),
)
