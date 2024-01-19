package storage

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

func GetForPrint(address string) {
	const op = "storage.GetForPrint: "
	fmt.Print(op)

	start := time.Now()

	count := 10000

	client := &fasthttp.Client{}

	for i := 0; i < count; i++ {
		json := `{
			"gtin": "00000000000000",
			"terminal":"test"
		}`

		req := fasthttp.AcquireRequest()
		req.SetRequestURI(address + "/v1/getCodeForPrint")
		req.Header.SetMethod("GET")
		req.SetBodyString(json)

		resp := fasthttp.AcquireResponse()
		err := client.Do(req, resp)
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode() != 200 {
			fmt.Println(string(resp.Body()))
		}

	}
	end := time.Now()

	diff := end.UnixMilli() - start.UnixMilli()

	fmt.Printf("\n%s: PASS - получено %d кодов, MS: %d, OPS: %.2f\n", op, count, diff, float64(count)/(float64(diff)/1000))
	fmt.Printf("========== %s PASSED ==========\n", op)

}
