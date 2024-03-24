-- name: CreateUser :one
INSERT INTO users (name, email, password, remember_token)
VALUES ($1, $2, $3, $4)
RETURNING id, name, email, email_verified_at, created_at, updated_at;

-- name: GetUserById :one
SELECT id, name, email, email_verified_at, created_at, updated_at 
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, name, email, email_verified_at, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users 
SET name = $2, email = $3, email_verified_at = $4, password = $5, remember_token = $6, updated_at = now()
WHERE id = $1
RETURNING id, name, email, email_verified_at, created_at, updated_at;

-- name: DeleteUser :one
DELETE
FROM users
WHERE id = $1
RETURNING id, name, email, email_verified_at, created_at, updated_at;

-- name: ListUsers :many
SELECT id, name, email, email_verified_at, created_at, updated_at
FROM users
WHERE (id > $1 or $1 IS NULL)
AND ($2 IS NULL OR name ILIKE $2)
AND ($3 IS NULL OR email ILIKE $3)
ORDER BY id ASC
LIMIT $4;

