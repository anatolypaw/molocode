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

	r.Post("/v1/addGood", AddGood(storage))
	r.Get("/v1/getAllGoods", GetGood(storage))

	r.Post("/v1/addCode", AddCode(storage))
	r.Post("/v1/setCodeProduced", Test())
	return r
}
