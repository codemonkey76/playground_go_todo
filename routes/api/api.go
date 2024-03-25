package api

import (
	"net/http"
	"todo/routes/api/tasks"
	"todo/routes/api/users"
)

func RegisterRoutes(mux *http.ServeMux) {
	users.RegisterRoutes(mux)
	tasks.RegisterRoutes(mux)
}
