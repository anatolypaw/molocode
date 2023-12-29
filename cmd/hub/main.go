package main

import (
	"log"
	v1 "molocode/internal/hub/http/v1"
	"molocode/internal/hub/storage"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting hub service")

	//Подключаемся к хранилищу
	storage, err := storage.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Fatalln(err)
	}
	defer storage.Close()
	log.Println("Storage ready")

	//Запускаем сервер веб интерфейса
	s := &http.Server{
		Addr:         ":80",
		Handler:      v1.Router(storage),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	log.Printf("Сервер веб интерфейса %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
