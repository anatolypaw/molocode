package main

import (
	"log"
	"molocode/internal/storage"
	"molocode/internal/ui"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	//Инициализируем базу данных
	storage, err := storage.NewMongodb("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("storage is init")
	defer storage.Close()

	//Запускаем сервер веб интерфейса
	go func() {
		s := &http.Server{
			Addr:         ":80",
			Handler:      ui.Router(),
			IdleTimeout:  1 * time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		log.Printf("Запуск сервера веб интерфейса %s", s.Addr)
		log.Fatal(s.ListenAndServe())
	}()

	//Инициализируем роутер для api работы с терминалами
	apiRouter := chi.NewRouter()
	//middleware
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.Recoverer)

	//apiRouter.Get("/v1/goods", web.GetGoods)

	apiSrv := &http.Server{
		Addr:         ":3000",
		Handler:      apiRouter,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//Запускаем сервер работы с терминалами
	log.Printf("starting terminal api server on %s", apiSrv.Addr)
	log.Fatal(apiSrv.ListenAndServe())
}
