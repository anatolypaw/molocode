package main

import (
	"log"
	"net/http"
	v1 "storage/http/v1"
	"storage/mongodb"
	"time"
)

func main() {
	log.Println("Starting storage service")

	//Подключаемся к хранилищу
	storage, err := mongodb.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Storage ready")

	//Запускаем сервер веб интерфейса
	s := &http.Server{
		Addr:         ":80",
		Handler:      v1.Router(storage),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Сервер веб интерфейса %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
