package http

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	fx.Provide(registerFiber),
)

func registerFiber() *fiber.App {
	app := fiber.New()
	return app
}
