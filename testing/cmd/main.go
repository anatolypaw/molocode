package main

import (
	"testing/tests/storage"
)

var addr = "http://localhost:80"

func main() {
	storage.CleanMongo("mongodb://localhost:27017/", "molocode")
	storage.AddGood(addr)
	storage.AddMoreCodes(addr)
}
