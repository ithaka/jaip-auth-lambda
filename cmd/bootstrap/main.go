package main

import (
	"github.com/JSTOR-Labs/online-pep/auth-lambda/pkg/routes"
	"github.com/JSTOR-Labs/online-pep/middleware"
	"github.com/akrylysov/algnhsa"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RealIP)
	r.Use(middleware.JSTORUUID)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	r.Get("/", routes.AuthJSTORHandler)

	algnhsa.ListenAndServe(r, nil)
}
