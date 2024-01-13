package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gatexch/model"
	"molocode/entity"
	"net/http"
)

// Передает в хаб код для печати
func UploadCodeToHub(address string, name string, code model.Code) error {
	const op = "httpClient.UploadCodeToHub"

	// URL to REST API endpoint
	apiURL := "http://" + address + "/v1/addCodeForPrint"

	// Передаем json структуру
	type reqModel struct {
		SourceName string `json:"source_name"`
		Gtin       string `json:"gtin"`
		Serial     string `json:"serial"`
		Crypto     string `json:"crypto"`
	}

	payload := reqModel{
		SourceName: name,
		Gtin:       code.Gtin,
		Serial:     code.Serial,
		Crypto:     code.Crypto,
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	// Выполняем HTTP POST запрос
	httpResp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadJson))
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	defer httpResp.Body.Close()

	var result entity.Response

	decoder := json.NewDecoder(httpResp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&result)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	if !result.Ok {
		return fmt.Errorf("%s: %s", op, result.Desc)
	}

	return nil
}
