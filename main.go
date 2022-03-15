package main

import (
	"fmt"
	"log"
	"net/http"
	UserController "teste-db/src/controller/user-controller"
	"teste-db/src/database"
	"teste-db/src/database/migrations"

	"github.com/gorilla/mux"
)

func main() {
	database.StartDB()
	r := mux.NewRouter()
	db := database.GetDatabase()
	migrations.RunMigrations(db)
	UserController.ExecuteAll(r)
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is running on port 8080")
}
