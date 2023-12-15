package storage

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
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

	count := 10_000
	for i := 0; i < count; i++ {
		json := fmt.Sprintf(`{
			"sourceName":"1c service",
			"gtin": "00000000000000",
			"serial": "%6d",
			"crypto": "%s"
		}`, i, RandStringBytes(4))

		var body bytes.Buffer
		body.WriteString(json)

		resp, err := http.Post(address+"/v1/addCode", "application/json", &body)
		if err != nil {
			fmt.Printf("%s: %s", op, err)
			i--
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			respbyte, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s ERROR: код ответа %d - %s\n", op, resp.StatusCode, string(respbyte))
			i--
		}

	}
	end := time.Now()

	diff := end.UnixMilli() - start.UnixMilli()

	fmt.Printf("%s: PASS - добавлено %d кодов, MS: %d, OPS: %.2f\n", op, count, diff, float64(count)/(float64(diff)/1000))
	fmt.Printf("========== %s PASSED ==========\n", op)

}
