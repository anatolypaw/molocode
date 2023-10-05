package main

import (
	"log"
	"molocode/internal/storage"
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

	//Инициализируем роутер
	router := chi.NewRouter()

	//middleware
	//router.Use(middleware.RequestID)
	//router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("./www/build/"))
	router.Handle("/*", fs)

	router.Get("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	srv := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
