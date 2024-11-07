package handler

import (
	"net/http"

	"github.com/AnthonyViniciusMuller/OAuth/cmd/rest_server/handler/auth"
	"github.com/go-chi/chi/v5"
)

var authHandler = auth.New()

func V1Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Get("/authorize", authHandler.Authorize)
	router.Get("/token", authHandler.Token)

	return router
}
