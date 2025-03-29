package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

func main() {
	mux := http.NewServeMux()

	todos := make([]TodoItem, 0)

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {
		var todoItem TodoItem
		if err := json.NewDecoder(r.Body).Decode(&todoItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todos = append(todos, todoItem)
		w.WriteHeader(http.StatusCreated)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
