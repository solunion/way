package configs

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func DatabaseConnection(config *Config) *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Database().User,
		config.Database().Pass,
		config.Database().Host,
		config.Database().Port,
		config.Database().Name,
		config.Database().SSLMode,
	)

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return bun.NewDB(conn, pgdialect.New())
}
