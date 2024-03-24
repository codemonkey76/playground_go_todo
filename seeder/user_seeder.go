package seeder

import (
	"context"
	"database/sql"
	"log"
	"todo/auth"
	db "todo/db/sqlc"
)

func SeedUsers(queries *db.Queries, ctx context.Context) {
	log.Println("Seeding Users")
	addUser(queries, ctx, "Shane Poppleton", "shane@bjja.com.au", "password")
}

func addUser(queries *db.Queries, ctx context.Context, name string, email string, password string) {
	hashedPassword, err := auth.HashPassword(password)

	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
		return
	}

	user, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:          name,
		Email:         email,
		Password:      hashedPassword,
		RememberToken: sql.NullString{},
	})

	if err != nil {
		log.Fatalf("Error creating user: %v", err)
		return
	}
	log.Println("User Created: ", user)
}
