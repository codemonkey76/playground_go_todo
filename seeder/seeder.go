package seeder

import (
	"context"
	"log"
	"todo/core"
	db "todo/db/sqlc"
)

func Seed(usersOnly bool) {
	log.Println("Seeding the database")

	pg := core.ConnectDb()
	defer pg.Close()
	ctx := context.Background()
	queries := db.New(pg)

	if usersOnly {
		SeedUsers(queries, ctx)
	} else {
		SeedAll(queries, ctx)
	}

	log.Println("Finished Seeding")
}

func Usage() {
	log.Println("Usage: seed [options]")
	log.Println("Options:")
	log.Println("  -users-only    Seed only users")
}

func SeedAll(queries *db.Queries, ctx context.Context) {
	log.Println("Seeding All Models")
	SeedUsers(queries, ctx)
}
