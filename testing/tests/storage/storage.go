package storage

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// Запускает все тесты
func Test(address string) {
	addGood(address)

}

func addGood(address string) {
	const op = "storage.addGood"
	var tests = []struct {
		testdesc string
		status   int
		json     string
	}{
		{"Продукт создается первый раз	", 201, `{"gtin": "00000000000000", "desc":"Описание"}`},
		{"Дубль продукта не создается 	", 400, `{"gtin": "00000000000000", "desc":"Описание"}`},
		{"Отсутствует описание        	", 400, `{"gtin": "00000000000000", "desc":""}`},
		{"Короткий gtin               	", 400, `{"gtin": "00000000000000", "desc":"Описание"}`},
		{"Длинный gtin             	  	", 400, `{"gtin": "000000000000000", "desc":"Описание"}`},
		{"Недопустимые символы gtin		", 400, `{"gtin": "0000000000000T", "desc":"Описание"}`},
		{"Битый json					", 400, `{"gtin": "0000000000000T", "desc":"Описание"`},
	}

	for i, test := range tests {

		var body bytes.Buffer
		body.WriteString(test.json)

		resp, err := http.Post(address+"/v1/goods", "application/json", &body)
		if err != nil {
			fmt.Printf("%s: %s", op, err)
			os.Exit(1)
		}

		if resp.StatusCode == test.status {
			fmt.Printf("%s: #%d - PASS - %s \n", op, i, test.testdesc)
		} else {
			fmt.Printf("%s: #%d - ERROR - %s \nresp: %v\n", op, i, test.testdesc, resp)
			fmt.Printf("========== %s ABORTED ==========\n", op)
			return
		}

	}

	fmt.Printf("========== %s PASSED ==========\n", op)

}
