package rule

import (
	"github.com/gofiber/fiber/v3"

	"go.uber.org/fx"
)

var Module = fx.Module("rule",
	fx.Provide(
		// Repository
		NewRepository,

		// Service
		NewService,

		// REST API
		NewRest,
	),
	fx.Invoke(
		registerHandlers,
	),
)

func registerHandlers(app *fiber.App, rest *Rest) {
	app.Post("/rules", rest.Create)
	app.Get("/rules/types/:type?", rest.GetAll)
	app.Get("/rules/:id", rest.GetById)
}
