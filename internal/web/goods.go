package web

import (
	"fmt"
	"net/http"
)

// Возвращает список продуктов из базы
func GetGoods(w http.ResponseWriter, r *http.Request) {
	const op = "api.web.good"
	_ = op
	fmt.Fprintf(w, `
	[
		{
		  "gtin": 460173924506,
		  "desc": "Biospan",
		  "storeCount": 2487,
		  "get": true,
		  "upload": true,
		  "awaible": false,
		  "shelfLife": 7
		},
		{
		  "gtin": 460414550200,
		  "desc": "Ginkogene",
		  "storeCount": 8247,
		  "get": false,
		  "upload": true,
		  "awaible": false,
		  "shelfLife": 4
		},
		{
		  "gtin": 460753859634,
		  "desc": "Medesign",
		  "storeCount": 9865,
		  "get": false,
		  "upload": false,
		  "awaible": false,
		  "shelfLife": 10
		},
		{
		  "gtin": 460641064678,
		  "desc": "Micronaut",
		  "storeCount": 8404,
		  "get": false,
		  "upload": false,
		  "awaible": false,
		  "shelfLife": 10
		},
		{
		  "gtin": 460686603342,
		  "desc": "Snorus",
		  "storeCount": 4396,
		  "get": false,
		  "upload": true,
		  "awaible": false,
		  "shelfLife": 12
		}
	  ]
	`)
}
