package rule

import (
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/http"
	"go.uber.org/fx"
)

var Module = fx.Module("rule",
	fx.Provide(
		// Repository
		generic.NewRepository,

		// Service
		NewService,
		http.NewService,

		// REST API
		NewRest,
		http.NewRest,
	),
	fx.Invoke(
		registerHandlers,
	),
)

func registerHandlers(app *fiber.App, rest *Rest) {
	app.Post("/rules", rest.Create)
	app.Get("/rules/:type?", rest.GetAll)
}
