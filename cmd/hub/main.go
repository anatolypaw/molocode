package main

import (
	"fmt"
	"log"
	"molocode/internal/service/storeservice"
	"molocode/internal/service/storeservice/mongo"
)

func main() {
	mongostore, err := mongo.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Panic(err)
	}

	storeservice := storeservice.NewStoreService(mongostore)

	goods, _ := storeservice.GetAllGoods()
	fmt.Printf("%v", goods)

}
