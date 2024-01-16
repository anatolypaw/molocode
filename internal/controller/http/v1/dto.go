package v1

import (
	"encoding/json"
	"time"
)

type response_dto struct {
	Ok   bool   `json:"ok"`
	Desc string `json:"desc"` // Описание результата
	Data any    `json:"data"`
}

// Возвращает JSON c данными и ошибкой, если она есть
func toResponse(ok bool, desc string, data any) string {
	result := response_dto{}
	result.Ok = ok
	result.Desc = desc
	result.Data = data

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}
	return string(resultJson)
}

type good_dto struct {
	Gtin            string    `json:"gtin"`
	Desc            string    `json:"desc"`
	StoreCount      uint      `json:"store_count"`
	GetCodeForPrint bool      `json:"get_code_for_print"`
	AllowProduce    bool      `json:"alllow_produce"`
	Upload          bool      `json:"upload"`
	CreatedAt       time.Time `json:",omitempty"`
}
