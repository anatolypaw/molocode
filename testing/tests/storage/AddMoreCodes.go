package storage

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func req() {

}

func AddMoreCodes(address string) {
	const op = "storage.addMoreCodes"
	fmt.Println(op)

	start := time.Now()

	json := fmt.Sprintf(`{
		"sourceName":"1c service",
		"gtin": "00000000000000",
		"serial": "%s",
		"crypto": "%s"
	}`, RandStringBytes(6), RandStringBytes(4))

	count := 10_000
	for i := 0; i < count; i++ {

		var body bytes.Buffer
		body.WriteString(json)

		resp, err := http.Post(address+"/v1/addCode", "application/json", &body)
		if err != nil {
			fmt.Printf("%s: %s", op, err)
			os.Exit(1)
		}

		if resp.StatusCode != http.StatusOK {
			respbyte, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s ERROR: код ответа %d - %s\n", op, resp.StatusCode, string(respbyte))
			return
		}

	}
	end := time.Now()

	diff := end.UnixMilli() - start.UnixMilli()

	fmt.Printf("%s: PASS - добавлено %d кодов, MS: %d, OPS: %.2f\n", op, count, diff, float64(count)/(float64(diff)/1000))
	fmt.Printf("========== %s PASSED ==========\n", op)

}
