package config

import "go.uber.org/fx"

var Module = fx.Module("config",
	fx.Provide(ContextConfiguration),
	fx.Provide(DotEnvConfiguration),
	fx.Provide(DatabaseConnection),
	fx.Provide(LoggerConfiguration),
)
