package tenant

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("tenant",
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
	app.Post("/tenants", rest.Create)
	app.Get("/tenants", rest.GetAll)
	app.Get("/tenants/:id", rest.GetById)
	app.Put("/tenants/:id", rest.Update)
	app.Delete("/tenants/:id", rest.Delete)
}
