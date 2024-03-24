// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

type Permission struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type PermissionRole struct {
	ID           int32        `json:"id"`
	PermissionID int64        `json:"permission_id"`
	RoleID       int64        `json:"role_id"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

type Role struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type RoleUser struct {
	ID        int32        `json:"id"`
	RoleID    int64        `json:"role_id"`
	UserID    int64        `json:"user_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Step struct {
	ID        int32        `json:"id"`
	TaskID    int64        `json:"task_id"`
	Name      string       `json:"name"`
	Completed bool         `json:"completed"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Task struct {
	ID         int32         `json:"id"`
	OwnerID    int64         `json:"owner_id"`
	AssigneeID sql.NullInt64 `json:"assignee_id"`
	Name       string        `json:"name"`
	Completed  bool          `json:"completed"`
	CreatedAt  sql.NullTime  `json:"created_at"`
	UpdatedAt  sql.NullTime  `json:"updated_at"`
}

type Team struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type TeamUser struct {
	ID        int32        `json:"id"`
	TeamID    int64        `json:"team_id"`
	UserID    int64        `json:"user_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type User struct {
	ID              int32          `json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	EmailVerifiedAt sql.NullTime   `json:"email_verified_at"`
	Password        string         `json:"password"`
	RememberToken   sql.NullString `json:"remember_token"`
	CreatedAt       sql.NullTime   `json:"created_at"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
}