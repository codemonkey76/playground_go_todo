-- name: CreateTask :one
INSERT INTO tasks (name, owner_id, assignee_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTaskById :one
SELECT *
FROM tasks
WHERE id = $1;

-- name: UpdateTask :one
UPDATE tasks 
SET name = $2, completed = $3, owner_id = $4, assignee_id = $5, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteTask :one
DELETE
FROM tasks
WHERE id = $1
RETURNING *;

-- name: ListTasks :many
SELECT *
FROM tasks
WHERE (id > $1 or $1 IS NULL)
AND ($2 IS NULL OR name ILIKE $2)
ORDER BY id ASC
LIMIT $3;

