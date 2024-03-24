package routes

import (
	"net/http"
	"todo/routes/api"
)

func RegisterRoutes(mux *http.ServeMux) {
	api.RegisterRoutes(mux)
}
