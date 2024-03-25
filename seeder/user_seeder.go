package seeder

import (
	"context"
	"database/sql"
	"log"
	"todo/auth"
	"todo/db/sqlc"
)

func SeedUsers(queries *sqlc.Queries, ctx context.Context) {
	log.Println("Seeding Users")
	addUser(queries, ctx, "Shane Poppleton", "shane@bjja.com.au", "password")
	log.Println("Finished Seeding Users")
}

func addUser(queries *sqlc.Queries, ctx context.Context, name string, email string, password string) {
	hashedPassword, err := auth.HashPassword(password)

	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return
	}

	user, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name:          name,
		Email:         email,
		Password:      hashedPassword,
		RememberToken: sql.NullString{},
	})

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return
	}
	log.Println("User Created: ", user)
}
