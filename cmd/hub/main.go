package main

import (
	"context"
	"log"
	"molocode/internal/app/storage/mongostore"
	"molocode/internal/app/usecase/usecase_admin"
	v1 "molocode/internal/view/http/v1"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	mstore, err := mongostore.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Panic("%v", err)
	}

	admUseCase := usecase_admin.New(mstore)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	ctx := context.Background()
	router.Post("/v1/addGood", v1.AddGood(ctx, &admUseCase))
	router.Get("/v1/getAllGoods", v1.GetAllGoods(ctx, &admUseCase))

	s := &http.Server{
		Addr:         ":80",
		Handler:      router,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Printf("Сервер веб интерфейса %s", s.Addr)
	log.Fatal("%v", s.ListenAndServe())
}
