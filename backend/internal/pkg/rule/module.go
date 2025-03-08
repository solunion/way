package rule

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("rule",
	fx.Provide(
		newHttpRepository,
		newHttpService[HttpRule],
		newHttpService[RouteRule],
		newHttpRest[HttpRule],
		newHttpRest[RouteRule],
	),
	fx.Invoke(
		registerHandlers[HttpRule],
		registerHandlers[RouteRule],
	),
)

func registerHandlers[T HttpRule | RouteRule](app *fiber.App, rest *Rest[T]) {
	app.Get("/rules", rest.GetAll)
	app.Post("/rules", rest.Create)
}
