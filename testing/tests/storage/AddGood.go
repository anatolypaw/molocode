package storage

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func AddGood(address string) {
	const op = "storage.addGood"
	var tests = []struct {
		testdesc string
		status   int
		json     string
	}{
		{"Продукт создается первый раз	", http.StatusCreated, `{"gtin": "00000000000000", "description":"Описание"}`},
		{"Дубль продукта не создается 	", http.StatusBadRequest, `{"gtin": "00000000000000", "description":"Описание"}`},
		{"Отсутствует описание        	", http.StatusBadRequest, `{"gtin": "00000000000000", "description":""}`},
		{"Короткий gtin               	", http.StatusBadRequest, `{"gtin": "00000000000000", "description":"Описание"}`},
		{"Длинный gtin             	  	", http.StatusBadRequest, `{"gtin": "000000000000000", "description":"Описание"}`},
		{"Недопустимые символы gtin		", http.StatusBadRequest, `{"gtin": "0000000000000T", "description":"Описание"}`},
		{"Битый json					", http.StatusBadRequest, `{"gtin": "00000000000000, "description":"Описание"`},
	}

	for i, test := range tests {

		var body bytes.Buffer
		body.WriteString(test.json)

		resp, err := http.Post(address+"/v1/goods", "application/json", &body)
		if err != nil {
			fmt.Printf("%s: %s", op, err)
			os.Exit(1)
		}

		respb, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == test.status {
			fmt.Printf("%s: #%d - PASS - %s - %s\n", op, i, test.testdesc, string(respb))
		} else {
			fmt.Printf("%s: #%d - ERROR - %s \nresp: %v\n", op, i, test.testdesc, string(respb))
			fmt.Printf("========== %s ABORTED ==========\n", op)
			return
		}

	}

	fmt.Printf("========== %s PASSED ==========\n", op)

}
