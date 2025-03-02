package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DbOptions struct {
	Host     string
	Db       string
	User     string
	Password string
}

func (db DbOptions) Connection() *bun.DB {
	// Create a PostgreSQL connection using pgdriver
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		db.User, db.Password, db.Host, db.Db)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun DB instance
	conn := bun.NewDB(sqldb, pgdialect.New())

	// Test the connection
	ctx := context.Background()
	if err := conn.PingContext(ctx); err != nil {
		panic(err)
	} else {
		fmt.Println("Connection has been established with database")
	}

	return conn
}
