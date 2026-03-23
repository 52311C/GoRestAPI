package main

import (
	"fmt"
	"go-todo-api/database"
	"go-todo-api/routes"
	"log"
	"net/http"
)

func main() {
	database.Connect()
	routes.Setup()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
