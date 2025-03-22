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
		fx.Annotated{
			Target: NewHttpHandler,
			Name:   "http_handler",
		},
		fx.Annotated{
			Target: NewRouteHandler,
			Name:   "route_handler",
		},
		// Service che utilizza gli handler
		fx.Annotate(
			NewService,
			fx.ParamTags(``, ``, `name:"http_handler"`, `name:"route_handler"`),
		),
		// REST API
		newHttpRest,
	),
	fx.Invoke(
		registerHandlers,
	),
)

func registerHandlers(app *fiber.App, rest *Rest) {
	app.Get("/rules", rest.GetAll)
	app.Post("/rules", rest.Create)
}
