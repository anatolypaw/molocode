package ts // Terminal Server

import (
	v1 "molocode/internal/ts/v1"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router() http.Handler {
	//Инициализируем роутер для api работы с терминалами
	r := chi.NewRouter()

	//middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/v1/test", v1.Test)
	return r
}