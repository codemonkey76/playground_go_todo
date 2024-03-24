package controllers

import db "todo/db/sqlc"

type TaskController struct{}

func (tc *TaskController) Get(id int) (db.Task, error) {
}

func (tc *TaskController) Post() {
	// Create a new task
}

func (tc *TaskController) Put() {
	// Update a task
}

func (tc *TaskController) Delete() {
	// Delete a task
}
