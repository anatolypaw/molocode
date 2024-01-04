package v1

import (
	"encoding/json"
	"molocode/entity"
)

// Возвращает JSON c данными и ошибкой, если она есть
func Response(ok bool, desc string, data any) string {

	result := entity.Response{}
	result.Ok = ok
	result.Desc = desc
	result.Data = data

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}

	return string(resultJson)
}
