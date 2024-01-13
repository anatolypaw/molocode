package httpClient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gatexch/model"
	"log"
	"net/http"
	"regexp"
)

// Запрашивает из внешнего  продукт и сколько кодов требуется для него запросить в 1с
func GetNewCodesFromGate(address string, gtin string, limit uint) ([]model.Code, error) {
	const op = "httpClient.GetNewCodesFromGate"

	// URL to REST API endpoint
	apiURL := "http://" + address + "/exchangemarks/hs/api/getmarks?gtin=" + gtin + "&limit=" + fmt.Sprint(limit)
	log.Print(op, " ", apiURL)

	// Выполняем HTTP GET запрос
	httpResp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer httpResp.Body.Close()

	// Проверяем, что код ответа сервера 200
	if httpResp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", httpResp.Status)
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	// Ожидаемая структура ответа
	type RespCode struct {
		Code string `json:"code"`
	}

	type ResponseModel struct {
		PageSize int        `json:"page_size"`
		Marks    []RespCode `json:"marks"`
	}
	var responeResult ResponseModel

	decoder := json.NewDecoder(httpResp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&responeResult)
	if err != nil {
		// возвращаем содержимое body
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	// Резальтат
	var resultMarks []model.Code

	// КМ закодирован в BASE64. Декодируем, и проверяем,
	// что gtin запрашиваемого кода совпадает с фактическим
	// Делаем это для каждой марки
	// var marks []entity.Mark
	for _, b64mark := range responeResult.Marks {
		markDec, err := base64.StdEncoding.DecodeString(b64mark.Code)
		if err != nil {
			fmt.Printf("%s %s\n", b64mark.Code, err)
			continue
		}
		// Разибраем декодированый код
		// 0104607009780870215s=hQH93bWg6
		re := regexp.MustCompile(`^01(0\d{13})21(.{6}).?93(.{4})$`)
		submatches := re.FindSubmatch(markDec)
		if len(submatches) != 4 {
			log.Printf("%s ошибка парсинга КМ", markDec)
			continue
		}
		resGtin := string(submatches[1])
		resSerial := string(submatches[2])
		resCrypto := string(submatches[3])

		if resGtin != gtin {
			log.Printf("%s: gtin не совпадает с запрошенным", markDec)
			continue
		}

		resultMarks = append(resultMarks, model.Code{Gtin: gtin, Serial: resSerial, Crypto: resCrypto})
	}

	return resultMarks, nil
}
