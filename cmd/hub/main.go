package main

import (
	"log"
	"molocode/internal/domain/usecase/usecase_admin"
	"molocode/internal/storage/mongo"

	v1 "molocode/internal/view/http/v1"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	mongoStore, err := mongo.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Panic(err)
	}

	admUseCase := usecase_admin.NewAdminUseCase(mongoStore)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/v1/addGood", v1.AddGood(&admUseCase))
	router.Get("/v1/getAllGoods", v1.GetAllGoods(&admUseCase))

	s := &http.Server{
		Addr:         ":80",
		Handler:      router,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Printf("Сервер веб интерфейса %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
