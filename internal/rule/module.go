package rule

import "go.uber.org/fx"

var Module = fx.Module("rule",
	fx.Provide(
		fx.Annotate(
			newRuleRepository,
			fx.As(new(RuleRepository)),
		),
	),
)
