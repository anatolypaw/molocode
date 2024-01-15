package main

import (
	"encoding/json"
	"fmt"
	"molocode/entity"
)

func main() {
	fmt.Print(Response(true, "", nil))
}

// Возвращает JSON c данными и ошибкой, если она есть
func Response(ok bool, desc string, data any) string {

	result := entity.Response{}
	result.Ok = ok
	result.Data = data

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}

	return string(resultJson)
}
