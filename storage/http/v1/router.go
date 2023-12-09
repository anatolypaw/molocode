package v1

import (
	"net/http"
	"storage/mongodb"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router(storage *mongodb.Storage) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/v1/good/addOne", AddGood(storage))
	r.Get("/v1/good/getAll", GetGood(storage))

	r.Post("/v1/code/addOne", Test())
	r.Post("/v1/code/setProduced", Test())
	return r
}
