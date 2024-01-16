package main

import (
	"log"
	v1 "molocode/internal/controller/http/v1"
	"molocode/internal/service/storeservice"
	"molocode/internal/service/storeservice/mongo"
	"molocode/internal/usecase/admin_usecase"
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

	mongostore, err := mongo.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Panic(err)
	}
	storeService := storeservice.NewStoreService(mongostore)
	admUsecase := admin_usecase.New(storeService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/v1/addGood", v1.AddGood(admUsecase))
	router.Get("/v1/getAllGoods", v1.GetAllGoods(admUsecase))

	//Запускаем сервер веб интерфейса
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
