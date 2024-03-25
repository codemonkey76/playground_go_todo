package seeder

import (
	"context"
	"database/sql"
	"log"
	db "todo/db/sqlc"
)

func SeedTasks(queries *db.Queries, ctx context.Context) {
	log.Println("Seeding Tasks")
	addTask(queries, ctx, "Task 1")
}

func addTask(queries *db.Queries, ctx context.Context, name string) {
	task, err := queries.CreateTask(ctx, db.CreateTaskParams{
		Name:       name,
		OwnerID:    1,
		AssigneeID: sql.NullInt64{},
	})

	if err != nil {
		log.Fatalf("Error creating task: %v", err)
		return
	}
	log.Println("Task Created: ", task)
}
