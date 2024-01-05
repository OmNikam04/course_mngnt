package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func setUpRoutes(app *chi.Mux){
	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
}