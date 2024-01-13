package router

import (
	v1 "hub/internal/http-server/v1"
	"hub/internal/storage"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router(storage *storage.Storage) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Работа с продуктами
	r.Post("/v1/addGood", v1.AddGood(storage))
	r.Get("/v1/getAllGoods", v1.GetAllGoods(storage))

	// Работа с кодами при производстве
	r.Get("/v1/getCodeForPrint", v1.GetCodeForPrint(storage))
	r.Post("/v1/setCodeProduced", v1.SetCodeProduced(storage))

	// Получение и выгрузка в 1с
	r.Post("/v1/addCodeForPrint", v1.AddCodeForPrint(storage))
	r.Get("/v1/getReqCodeCount", v1.GetReqCodeCount(storage)) // Возвращает gtin и сколько кодов не хватает

	r.Get("/test", v1.Test())

	return r
}
