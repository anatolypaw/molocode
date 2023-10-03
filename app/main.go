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
	storage, err := storage.NewSqlite3("./data/sqlite3/main.db")
	if err != nil {
		log.Panic(err)
	}
	log.Print("storage is init")
	defer storage.Close()

	//Инициализируем роутер
	router := chi.NewRouter()

	//middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	srv := &http.Server{
		Addr:         ":3000",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
