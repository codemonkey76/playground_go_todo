package seeder

import (
	"context"
	"log"
	"todo/core"
	db "todo/db/sqlc"
)

func Seed(args []string) {
	log.Println("Seeding the database")

	pg := core.ConnectDb()
	defer pg.Close()
	ctx := context.Background()
	queries := db.New(pg)

	if len(args) < 1 {
		SeedAll(queries, ctx)
		return
	}

	switch args[0] {
	case "users":
		SeedUsers(queries, ctx)
	case "all":
		SeedAll(queries, ctx)
	}

	log.Println("Finished Seeding")
}

func SeedAll(queries *db.Queries, ctx context.Context) {
	log.Println("Seeding All Models")
	SeedUsers(queries, ctx)
}
