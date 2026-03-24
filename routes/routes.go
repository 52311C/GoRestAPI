package routes

import (
	"encoding/json"
	"go-todo-api/handlers"
	"net/http"
)

func Setup() {
	// Serve a simple browser UI from ./static at the root path.
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetItems(w, r)
		case http.MethodPost:
			handlers.CreateItem(w, r)
		case http.MethodPut:
			handlers.UpdateItem(w, r)
		case http.MethodDelete:
			handlers.DeleteItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
