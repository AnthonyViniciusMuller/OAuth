package main

import (
	"fmt"
	"net/http"

	"github.com/AnthonyViniciusMuller/OAuth/cmd/rest_server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Mount("/v1", handler.V1Router())

	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080", router)
}
