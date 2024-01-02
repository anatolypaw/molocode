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
		{"Отсутствует описание          ", http.StatusBadRequest, `{"gtin": "00000000000000", "description":""}`},
		{"Короткий gtin                 ", http.StatusBadRequest, `{"gtin": "0000000000000", "description":"Описание"}`},
		{"Длинный gtin                  ", http.StatusBadRequest, `{"gtin": "000000000000000", "description":"Описание"}`},
		{"Недопустимые символы gtin     ", http.StatusBadRequest, `{"gtin": "0000000000000T", "description":"Описание"}`},
		{"Битый json                    ", http.StatusBadRequest, `{"gtin": "00000000000000, "description":"Описание"`},
		{"Некорректное имя ключа        ", http.StatusBadRequest, `{"abcd": "00000000000000", "badkey":"Продукт 1"}`},
		{"Продукт создается первый раз  ", http.StatusOK, `{
															"Gtin": "00000000000000", 
															"description":"Продукт 1", 
															"getcodeforprint": true, 
															"allowproduce": true
															}`},
		{"Дубль продукта не создается   ", http.StatusBadRequest, `{"gtin": "00000000000000", "description":"Описание"}`},
		{"Создается продукт 2 первый    ", http.StatusOK, `{"gtin": "00000000000002", "description":"Продукт 2", "getcodeforprint": true}`},
	}

	for i, test := range tests {

		var body bytes.Buffer
		body.WriteString(test.json)

		resp, err := http.Post(address+"/v1/addGood", "application/json", &body)
		if err != nil {
			fmt.Printf("%s: %s", op, err)
			os.Exit(1)
		}

		respbyte, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == test.status {
			fmt.Printf("%s: #%d - PASS - %s - %s\n", op, i, test.testdesc, string(respbyte))
		} else {
			fmt.Printf("%s: #%d - ERROR - %s \nresp: %v\n", op, i, test.testdesc, string(respbyte))
			fmt.Printf("========== %s FAIL ==========\n", op)
			return
		}

	}

	fmt.Printf("========== %s PASSED ==========\n", op)

}
