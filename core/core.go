package core

import (
	"database/sql"
	_ "github.com/lib/pq"
	"todo/env"
)

func ConnectDb() (*sql.DB, error) {
	connStr := env.GetConnectionString()
	return sql.Open("postgres", connStr)
}
