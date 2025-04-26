package db

import (
	"fmt"
	"proyecto/internal/config/env"
)

type Dsn struct {
	user string
	password string
	host string
	port string
	dbName string
}

func GetDSN() string {
	if env.Get("APP_ENV") == "local" {
		dsn := Dsn {
			user: env.Get("DB_USER"),
			password: env.Get("DB_PASSWORD"),
			host: env.Get("DB_HOST"),
			port: env.Get("DB_PORT"),
			dbName: env.Get("DB_DATABASE"),
		}
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dsn.user, dsn.password, dsn.host, dsn.port, dsn.dbName)
	}

	return env.Get("DATABASE_URL")
}
