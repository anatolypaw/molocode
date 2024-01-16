package main

import (
	"fmt"
	"log"
	"molocode/internal/adapters/hubstore_mongo"
	"molocode/internal/domain/service/hubservice"
)

func main() {
	storage, err := hubstore_mongo.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Panic(err)
	}

	hub := hubservice.NewHubService(storage)

	goods, _ := hub.GetAllGoods()
	fmt.Printf("%v", goods)

}
