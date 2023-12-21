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

	// Работа с продуктами
	r.Post("/v1/addGood", AddGood(storage))
	r.Get("/v1/getAllGoods", GetGoods(storage))

	// Работа с кодами
	r.Post("/v1/addCodeForPrint", AddCodeForPrint(storage))
	r.Get("/v1/getCodeForPrint", GetCodeForPrint(storage))

	r.Post("/v1/setCodeProduced", Test())
	return r
}
