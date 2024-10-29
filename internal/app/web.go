package main

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/solunion/way/internal/pkg/config"
	"github.com/solunion/way/internal/pkg/web"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		config.Module,
		web.Module,
		fx.Invoke(runWebApp),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
	app.Run()
}

func runWebApp(lc fx.Lifecycle, log *zap.SugaredLogger, web *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("starting web server...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping web server...")
			return nil
		},
	})
}
