package seeder

import (
	"context"
	"fmt"
	"log"
	"todo/core"
	db "todo/db/sqlc"
	"todo/ui"
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
	ui.Colorize(ui.ColorYellow, "Description:\n")
	fmt.Printf("    Seed the database with some data\n\n")
	ui.Colorize(ui.ColorYellow, "Usage:\n")
	fmt.Print("    -users-only    Seed only the users table\n\n")
}

func SeedAll(queries *db.Queries, ctx context.Context) {
	log.Println("Seeding All Models")
	SeedUsers(queries, ctx)
}
