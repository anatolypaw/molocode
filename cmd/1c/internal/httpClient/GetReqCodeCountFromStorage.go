package httpClient

import (
	"encoding/json"
	"fmt"
	"molocode/entity"
	"net/http"
)

// Запрашивает из storage продукт и сколько кодов требуется для него запросить в 1с
func GetReqCodeCountFromStorage(address string) ([]entity.CodeReq, error) {
	const op = "http.Client.GetReqCodeCountFromStorage"

	// URL to REST API endpoint
	apiURL := "http://" + address + "/v1/getReqCodeCount"

	// Выполняем HTTP GET запрос
	response, err := http.Get(apiURL)
	if err != nil {
		return []entity.CodeReq{}, fmt.Errorf("%s: %s", op, err)
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", response.Status)
		return []entity.CodeReq{}, fmt.Errorf("%s: %s", op, err)
	}

	// Декодируем JSON  в структуру
	reqCodeCount := []entity.CodeReq{}
	err = json.NewDecoder(response.Body).Decode(&reqCodeCount)
	if err != nil {
		return []entity.CodeReq{}, fmt.Errorf("%s: %s", op, err)
	}

	return reqCodeCount, nil
}
