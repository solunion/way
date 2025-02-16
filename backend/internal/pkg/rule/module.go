package rule

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("rule",
	fx.Provide(
		newRepository,
		newService,
		newRest,
	),
	fx.Invoke(
		registerHandlers,
	),
)

func registerHandlers(app *fiber.App, rest *Rest) {
	app.Get("/rules", rest.GetAll)
}
