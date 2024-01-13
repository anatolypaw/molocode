package httpClient

import (
	"encoding/json"
	"fmt"
	"molocode/entity"
	"net/http"
)

// Запрашивает из сервиса hub продукт и сколько кодов требуется для него запросить в 1с
func GetReqCodeCountFromHub(address string) ([]entity.CodeReq, error) {
	const op = "httpClient.GetReqCodeCountFromHub"

	// URL to REST API endpoint
	apiURL := "http://" + address + "/v1/getReqCodeCount"

	// Выполняем HTTP GET запрос
	httpResp, err := http.Get(apiURL)
	if err != nil {
		return []entity.CodeReq{}, fmt.Errorf("%s: %s", op, err)
	}
	defer httpResp.Body.Close()

	// Check if the response status code is OK (200)
	if httpResp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", httpResp.Status)
		return []entity.CodeReq{}, fmt.Errorf("%s: %s", op, err)
	}

	type Response struct {
		Ok   bool
		Desc string           // Описание результата запроса
		Data []entity.CodeReq // Ожиданемое содержимое ответа
	}

	var result Response

	decoder := json.NewDecoder(httpResp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&result)
	if err != nil {
		return []entity.CodeReq{}, fmt.Errorf("%s: %s", op, err)
	}

	if !result.Ok {
		return []entity.CodeReq{}, fmt.Errorf("%s: %s", op, result.Desc)
	}

	return result.Data, nil
}
