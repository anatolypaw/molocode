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
		{"Отсутствует описание          ", http.StatusBadRequest, `{"gtin": "00000000000000", "desc":""}`},
		{"Короткий gtin                 ", http.StatusBadRequest, `{"gtin": "0000000000000", "desc":"Описание"}`},
		{"Длинный gtin                  ", http.StatusBadRequest, `{"gtin": "000000000000000", "desc":"Описание"}`},
		{"Недопустимые символы gtin     ", http.StatusBadRequest, `{"gtin": "0000000000000T", "desc":"Описание"}`},
		{"Битый json                    ", http.StatusBadRequest, `{"gtin": "00000000000000, "desc":"Описание"`},
		{"Некорректное имя ключа        ", http.StatusBadRequest, `{"abcd": "00000000000000", "badkey":"Продукт 1"}`},
		{"Продукт создается первый раз  ", http.StatusOK, `{
															"gtin": "00000000000000", 
															"desc":"Продукт 1", 
															"get_code_for_print": true, 
															"allow_produce": true
															}`},
		{"Дубль продукта не создается   ", http.StatusBadRequest, `{"gtin": "00000000000000", "desc":"Описание"}`},
		{"Создается продукт 2 первый    ", http.StatusOK, `{"gtin": "00000000000002", "desc":"Продукт 2", "get_code_for_print": true}`},
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
			log.Fatal()
			return
		}

	}

	fmt.Printf("========== %s PASSED ==========\n", op)

}
