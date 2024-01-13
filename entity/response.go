package entity

import "encoding/json"

type Response struct {
	Ok   bool   `json:"ok"`
	Desc string `json:"desc"` // Описание результата
	Data any    `json:"data"`
}

// Возвращает JSON c данными и ошибкой, если она есть
func ToResponse(ok bool, desc string, data any) string {

	result := Response{}
	result.Ok = ok
	result.Desc = desc
	result.Data = data

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}

	return string(resultJson)
}
