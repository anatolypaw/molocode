package main

import (
	"log"
	"storage/internal/storage"
)

func main() {
		//Инициализируем базу данных
		storage, err := storage.New("mongodb://localhost:27017/", "molocode")
		if err != nil {
			log.Fatal(err)
		}
		log.Print("storage ready")

		_ = storage

}

