package storage

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/valyala/fasthttp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func AddMoreCodes(address string) {
	const op = "storage.addMoreCodes: "
	fmt.Print(op)

	start := time.Now()

	count := 1_000_000
	fmt.Printf("%10d", 0)

	client := &fasthttp.Client{}

	for i := 0; i < count; i++ {
		//time.Sleep(1 * time.Second)
		json := fmt.Sprintf(`{
			"sourceName":"1c service",
			"gtin": "00000000000000",
			"serial": "%s",
			"crypto": "%s"
		}`, RandStringBytes(6), RandStringBytes(4))

		req := fasthttp.AcquireRequest()
		req.SetRequestURI(address + "/v1/addCode")
		req.Header.SetMethod("POST")
		req.SetBodyString(json)

		resp := fasthttp.AcquireResponse()
		err := client.Do(req, resp)
		if err != nil {
			fmt.Println(err)
		}

		if i%100 == 0 {
			fmt.Printf("\b\b\b\b\b\b\b\b\b\b%10d", i)
		}

	}
	end := time.Now()

	diff := end.UnixMilli() - start.UnixMilli()

	fmt.Printf("%s: PASS - добавлено %d кодов, MS: %d, OPS: %.2f\n", op, count, diff, float64(count)/(float64(diff)/1000))
	fmt.Printf("========== %s PASSED ==========\n", op)

}
