package main

import (
	"fmt"
	"testing/tests/storage"
)

var addr = "http://localhost:80"

func main() {
	storage.CleanMongo("mongodb://localhost:27017/", "molocode")
	fmt.Print("Запустите сервис storage и нажмите enter")
	fmt.Scanln()
	storage.AddGood(addr)
}
