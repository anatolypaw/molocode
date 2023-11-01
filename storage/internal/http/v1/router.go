package v1

import (
	"net/http"
	"storage/internal/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router(storage *storage.Connection) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/v1/goods", AddGood(storage))
	return r
}
