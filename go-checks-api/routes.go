package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) handlers() http.Handler {
	mux := chi.NewMux()

	mux.Post("/validate", validate)

	return mux
}
