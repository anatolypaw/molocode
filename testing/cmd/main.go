package main

import (
	"testing/tests/storage"
)

var addr = "http://127.0.0.1"

func main() {
	//	storage.CleanMongo("mongodb://localhost:27017/", "molocode")
	storage.AddGood(addr)
	storage.AddMoreCodes(addr)
}
