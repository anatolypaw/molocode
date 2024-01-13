package httpClient

import (
	"encoding/json"
	"fmt"
	"gatexch/model"
	"log"
	"net/http"
)

// Запрашивает из внешнего  продукт и сколько кодов требуется для него запросить в 1с
func GetNewCodesFromGate(address string, gtin string, limit uint) ([]model.Mark, error) {
	const op = "httpClient.GetNewCodesFromGate"

	// URL to REST API endpoint
	apiURL := "http://" + address + "/exchangemarks/hs/api/getmarks?gtin=" + gtin + "&limit=" + fmt.Sprint(limit)
	log.Print(apiURL)

	// Выполняем HTTP GET запрос
	httpResp, err := http.Get(apiURL)
	if err != nil {
		return []model.Mark{}, fmt.Errorf("%s: %s", op, err)
	}
	defer httpResp.Body.Close()

	// Проверяем, что код ответа сервера 200
	if httpResp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", httpResp.Status)
		return []model.Mark{}, fmt.Errorf("%s: %s", op, err)
	}

	// Ожидаемая структура ответа
	type ResponseModel struct {
		PageSize int          `json:"page_size"`
		Marks    []model.Mark `json:"marks"`
	}
	var responeResult ResponseModel

	decoder := json.NewDecoder(httpResp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&responeResult)
	if err != nil {
		// возвращаем содержимое body
		return []model.Mark{}, fmt.Errorf("%s: %s", op, err)
	}

	// КМ закодирован в BASE64. Декодируем, и проверяем,
	// что gtin запрашиваемого кода совпадает с фактическим

	return responeResult.Marks, nil
}
