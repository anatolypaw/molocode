package main

import (
	"log"
	"net/http"
	http_router "storage/internal/controller/http"
	"storage/internal/storage"
	"storage/internal/usecase"
	"time"
)

func main() {

	usecase.AddGood("00000000000000", "Молоко 55%")
	usecase.SetDescription("00000000000000", "lalalasdlalsd")
	log.Println("Starting app")

	//Подключаемся к хранилищу
	storage, err := storage.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Storage ready")

	//Запускаем сервер веб интерфейса
	s := &http.Server{
		Addr:         ":80",
		Handler:      http_router.Router(storage),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Сервер веб интерфейса %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
