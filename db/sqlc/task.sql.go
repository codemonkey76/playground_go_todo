// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: task.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (name, owner_id, assignee_id)
VALUES ($1, $2, $3)
RETURNING id, name, completed, owner_id, assignee_id, created_at, updated_at
`

type CreateTaskParams struct {
	Name       string        `json:"name"`
	OwnerID    int64         `json:"owner_id"`
	AssigneeID sql.NullInt64 `json:"assignee_id"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.Name, arg.OwnerID, arg.AssigneeID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.OwnerID,
		&i.AssigneeID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :one
DELETE
FROM tasks
WHERE id = $1
RETURNING id, name, completed, owner_id, assignee_id, created_at, updated_at
`

func (q *Queries) DeleteTask(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRowContext(ctx, deleteTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.OwnerID,
		&i.AssigneeID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTaskById = `-- name: GetTaskById :one
SELECT id, name, completed, owner_id, assignee_id, created_at, updated_at
FROM tasks
WHERE id = $1
`

func (q *Queries) GetTaskById(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskById, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.OwnerID,
		&i.AssigneeID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, name, completed, owner_id, assignee_id, created_at, updated_at
FROM tasks
WHERE (id > $1 or $1 IS NULL)
AND ($2 IS NULL OR name ILIKE $2)
ORDER BY id ASC
LIMIT $3
`

type ListTasksParams struct {
	ID      int32       `json:"id"`
	Column2 interface{} `json:"column_2"`
	Limit   int32       `json:"limit"`
}

func (q *Queries) ListTasks(ctx context.Context, arg ListTasksParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks, arg.ID, arg.Column2, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Completed,
			&i.OwnerID,
			&i.AssigneeID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :one
UPDATE tasks 
SET name = $2, completed = $3, owner_id = $4, assignee_id = $5, updated_at = now()
WHERE id = $1
RETURNING id, name, completed, owner_id, assignee_id, created_at, updated_at
`

type UpdateTaskParams struct {
	ID         int32         `json:"id"`
	Name       string        `json:"name"`
	Completed  bool          `json:"completed"`
	OwnerID    int64         `json:"owner_id"`
	AssigneeID sql.NullInt64 `json:"assignee_id"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTask,
		arg.ID,
		arg.Name,
		arg.Completed,
		arg.OwnerID,
		arg.AssigneeID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.OwnerID,
		&i.AssigneeID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
