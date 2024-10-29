package web

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/internal/pkg/configs"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("web",
	fx.Provide(registerFiber),
)

func registerFiber(lc fx.Lifecycle, log *zap.SugaredLogger, config *configs.Config) *fiber.App {
	app := fiber.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error = nil
			go func() {
				err = app.Listen(fmt.Sprintf("%s:%d", config.Web.Host, config.Web.Port)) //fiber.ListenConfig{EnablePrefork: true},

			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
