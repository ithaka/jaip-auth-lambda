package main

import (
	"main/pkg/routes"

	"github.com/akrylysov/algnhsa"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func main() {
	r := chi.NewRouter()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	r.Get("/", routes.AuthJSTORHandler)

	algnhsa.ListenAndServe(r, nil)
}
