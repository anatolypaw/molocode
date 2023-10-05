package main

import (
	"log"
	"molocode/internal/storage"
	"molocode/internal/web"
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

	//Инициализируем роутер для api web интерфейса
	webRouter := chi.NewRouter()
	//middleware
	webRouter.Use(middleware.Logger)
	webRouter.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("./www/build/"))
	webRouter.Handle("/*", fs)

	webRouter.Get("/goods", web.GetGoods)

	webSrv := &http.Server{
		Addr:         ":80",
		Handler:      webRouter,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//Инициализируем роутер для api работы с терминалами
	apiRouter := chi.NewRouter()
	//middleware
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.Recoverer)

	apiRouter.Get("/v1/goods", web.GetGoods)

	apiSrv := &http.Server{
		Addr:         ":3000",
		Handler:      apiRouter,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("starting web interface server on %s", webSrv.Addr)
		log.Fatal(webSrv.ListenAndServe())
	}()

	log.Printf("starting terminal api server on %s", apiSrv.Addr)
	log.Fatal(apiSrv.ListenAndServe())

}
