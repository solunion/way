package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/backend/tenant"
	"libs/backend-new/http"

	//"github.com/solunion/way/internal/pkg/http"
	//"github.com/solunion/way/internal/pkg/tenant"
	"github.com/solunion/way/backend/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		config.Module,
		http.Module,
		tenant.Module,
		fx.Invoke(runWebApp),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
	app.Run()
}

func runWebApp(lc fx.Lifecycle, log *zap.SugaredLogger, config *config.Config, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error = nil
			go func() {
				log.Info("starting web server...")
				err = app.Listen(fmt.Sprintf("%s:%d", config.Web.Host, config.Web.Port)) //fiber.ListenConfig{EnablePrefork: true},

			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping web server...")
			return app.Shutdown()
		},
	})
}
