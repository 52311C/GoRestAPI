package handlers

import (
	"encoding/json"
	"go-todo-api/database"
	"go-todo-api/models"
	"net/http"
)

// GET /items
func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name FROM items")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ID, &item.Name)
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// POST /items
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	err := database.DB.QueryRow(
		"INSERT INTO items(name) VALUES($1) RETURNING id",
		item.Name,
	).Scan(&item.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// DELETE /items?id=1
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := database.DB.Exec("DELETE FROM items WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
