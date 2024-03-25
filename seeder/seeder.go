package seeder

import (
	"context"
	"fmt"
	"log"
	"todo/database"
	"todo/db/sqlc"
	"todo/ui"
)

func Seed(usersOnly bool) {
	log.Println("Seeding the database")

	pg := database.GetDB()
	defer pg.Close()

	ctx := context.Background()
	queries := sqlc.New(pg)

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

func SeedAll(queries *sqlc.Queries, ctx context.Context) {
	log.Println("Seeding All Models")
	SeedUsers(queries, ctx)
	SeedTasks(queries, ctx)
}
