package main

import (
	"fmt"
	"log"
	"molocode/internal/adapters/hubstore"
	"molocode/internal/domain/hubservice"
)

func main() {
	storage, err := hubstore.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Panic(err)
	}

	hub := hubservice.New(storage)

	goods, _ := hub.GetAllGoods()
	fmt.Printf("%v", goods)

}
