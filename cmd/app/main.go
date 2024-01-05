package main

import (
	"net/http"
	"fmt"
	"student-courses-with-transactions/database"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// loading envs
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	
	database.ConnectDb()

	app := chi.NewRouter()
	setUpRoutes(app)
	
	http.ListenAndServe(":3000", app)
}
