package tasks

import (
	"encoding/json"
	"net/http"
	"todo/database"
	"todo/db/sqlc"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/tasks", index)
	mux.HandleFunc("POST /api/v1/tasks", store)
	mux.HandleFunc("GET /api/v1/tasks/:id", get)
	mux.HandleFunc("PUT /api/v1/tasks/:id", update)
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get task"))
}

func index(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()

	queries := sqlc.New(db)

	tasks, err := queries.ListTasks(r.Context(), sqlc.ListTasksParams{
		ID:      0,
		Column2: nil,
		Limit:   100,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jData, err := json.Marshal(tasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func store(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request) {

}
func destroy(w http.ResponseWriter, r *http.Request) {

}
