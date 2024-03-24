package users

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/users", index)
	mux.HandleFunc("POST /api/v1/users", store)
	mux.HandleFunc("GET /api/v1/users/:id", get)
	mux.HandleFunc("PUT /api/v1/users/:id", update)
}

func get(w http.ResponseWriter, r *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {

}

func store(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request) {

}
func destroy(w http.ResponseWriter, r *http.Request) {

}
