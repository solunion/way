package rule

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("rule",
	fx.Provide(
		// Repository
		newHttpRepository,
		// Handlers per i diversi tipi di regole con annotazioni
		AsRuleHandler(NewHttpHandler),
		AsRuleHandler(NewRouteHandler),
		fx.Annotate(
			NewService,
			fx.ParamTags(``, ``, `group:"handlers"`),
		),
		// REST API
		newHttpRest,
	),
	fx.Invoke(
		registerHandlers,
	),
)

func AsRuleHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(RuleHandler)),
		fx.ResultTags(`group:"handlers"`),
	)
}

func registerHandlers(app *fiber.App, rest *Rest) {
	app.Get("/rules", rest.GetAll)
	app.Post("/rules", rest.Create)
}
