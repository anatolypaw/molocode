package main

import (
	"fmt"
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

	for i := 0; i < 3; i++ {
		err = storage.AddGood(fmt.Sprintf("%d", i), "hellso")
		if err != nil {
			log.Fatal(err)
		}
	}

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
	// err = srv.ListenAndServe()
	// log.Fatal(err)
}
