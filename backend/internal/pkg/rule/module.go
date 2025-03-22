package rule

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("rule",
	fx.Provide(
		// Repository
		newRepository,

		// Service
		newService,

		// REST API
		newRest,
	),
	fx.Invoke(
		registerHandlers,
	),
)

func registerHandlers(app *fiber.App, rest *Rest) {
	app.Get("/rules", rest.GetAll)
	app.Post("/rules", rest.Create)
}
