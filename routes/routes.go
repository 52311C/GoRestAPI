package routes

import (
	"go-todo-api/handlers"
	"net/http"
)

func Setup() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetItems(w, r)
		case http.MethodPost:
			handlers.CreateItem(w, r)
		case http.MethodDelete:
			handlers.DeleteItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
