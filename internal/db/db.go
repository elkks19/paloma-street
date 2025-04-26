package db

import (
	"proyecto/internal/config/env"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	_ "github.com/uptrace/bun/driver/pgdriver"
)

func Init() *bun.DB {
	rawConn, err := sql.Open("pg", GetDSN())
	if err != nil {
		panic(err)
	}

	// NOTE: CONFIGURE BASED ON USAGE
	rawConn.SetMaxIdleConns(10)
	rawConn.SetMaxOpenConns(50)
	rawConn.SetConnMaxIdleTime(30 * time.Minute)
	rawConn.SetConnMaxLifetime(0)

	conn := bun.NewDB(rawConn, pgdialect.New())
	if env.Get("APP_DEBUG") == "true" {
		conn.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}

	return conn
}
