package http_router

import (
	"net/http"
	v1 "storage/internal/http/v1"
	"storage/internal/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router(storage *storage.Connection) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/v1/goods", v1.AddGood(storage))
	return r
}
