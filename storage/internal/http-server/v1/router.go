package v1

import (
	"net/http"
	"storage/internal/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router(storage *storage.Storage) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Работа с продуктами
	r.Post("/v1/addGood", AddGood(storage))
	r.Get("/v1/getAllGoods", GetGoods(storage))

	// Работа с кодами при производстве
	r.Get("/v1/getCodeForPrint", GetCodeForPrint(storage))
	r.Post("/v1/setCodeProduced", SetCodeProduced(storage))

	// Получение и выгрузка в 1с
	r.Post("/v1/addCodeForPrint", AddCodeForPrint(storage))
	r.Get("/v1/getReqCodeCount", GetReqCodeCount(storage)) // Возвращает gtin и сколько кодов не хватает

	return r
}
