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

		NewProva,
	),
	fx.Invoke(
		registerHandlers,
	),
)

func registerHandlers(app *fiber.App, rest *Rest, prova *Prova) {
	app.Post("/rules", rest.Create)
	app.Get("/rules/:type?", rest.GetAll)
	app.Get("/prova", prova.Get)
}
