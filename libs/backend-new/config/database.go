package config

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/alexlast/bunzap"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func DatabaseConnection(lc fx.Lifecycle, config *Config, logger *zap.Logger) *bun.DB {

	dsn := config.Database().Uri

	if dsn == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			config.Database().User,
			config.Database().Pass,
			config.Database().Host,
			config.Database().Port,
			config.Database().Name,
			config.Database().SSLMode,
		)
	}

	fmt.Printf("Database connection string: %s\n", dsn)

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, pgdialect.New())

	db.AddQueryHook(bunzap.NewQueryHook(bunzap.QueryHookOptions{
		Logger: logger,
		//SlowDuration: 200 * time.Millisecond, // Omit to log all operations as debug
	}))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			if db != nil {
				return db.Close()
			}
			return nil
		},
	})

	return db
}
