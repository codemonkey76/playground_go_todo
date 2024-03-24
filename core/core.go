package core

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"todo/env"
)

func ConnectDb() *sql.DB {
	connStr := env.GetConnectionString()
	pg, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	return pg
}
