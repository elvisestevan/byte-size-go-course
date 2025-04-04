package transport

import (
	"byte-size-go-course/internal/todo"
	"encoding/json"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoService *todo.Service) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(todoService.GetAll()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("GET /todo/search", func(w http.ResponseWriter, r *http.Request) {
		term := r.URL.Query().Get("term")
		if term == "" {
			http.Error(w, "Missing search term", http.StatusBadRequest)
			return
		}
		results := todoService.Search(term)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {
		var todoItem TodoItem
		if err := json.NewDecoder(r.Body).Decode(&todoItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := todoService.Add(todoItem.Item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("DELETE /todo", func(w http.ResponseWriter, r *http.Request) {
		var todoItem TodoItem
		if err := json.NewDecoder(r.Body).Decode(&todoItem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := todoService.Delete(todoItem.Item)
		if err != nil {
			http.Error(w, "Todo item not found", http.StatusNotFound)
		}

	})

	return &Server{mux: mux}
}

func (s *Server) Serve() error {

	if err := http.ListenAndServe(":8080", s.mux); err != nil {
		return err
	}
	return nil
}
