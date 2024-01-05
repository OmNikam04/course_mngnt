package main

import (
	"fmt"
	"net/http"

	"student-courses-with-transactions/internal/repository"
	"student-courses-with-transactions/internal/usecases"

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

	// connecting to db
	db, err := repository.ConnectDb()

	if err != nil {
		fmt.Printf("Error while coonecting to database %v", err)
	}

	// Registering all usecases(services)
	dao := repository.NewDAO(db)
	usecases.NewCourseService(dao)


	// setting chi router
	app := chi.NewRouter()
	setUpRoutes(app)
	
	http.ListenAndServe(":3000", app)
}
