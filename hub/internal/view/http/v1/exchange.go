package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/ctxlogger"
	"hub/internal/entity"
	"hub/internal/usecase/exchange"

	"net/http"
)

// Добавляет продукт
// метод POST
func GetGoodsReqCodes(u exchange.ExchangeUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.GetGoodsReqCodes")
		reqId := ctxlogger.GetReqID(r.Context())

		// Получаем список продуктов и требуемое количество кодов
		codereq, err := u.GetGoodsReqCodes(r.Context())
		if err != nil {
			l.Warn("Ошибка запроса продуктов", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		type codeReq_json struct {
			Gtin     string `json:"gtin"`
			Desc     string `json:"desc"`
			Required uint   `json:"required"`
		}

		// MAPPING
		mappedCodeReq := []codeReq_json{}
		for _, ths := range codereq {
			mappedCodeReq = append(mappedCodeReq, codeReq_json{
				Gtin:     ths.Gtin,
				Desc:     ths.Desc,
				Required: ths.Required,
			})
		}

		resp_body := toResponse(reqId, true, "Успешно", mappedCodeReq)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}

// Доабвляет код для печати
func AddCodeForPrint(usecase exchange.ExchangeUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.AddCodeForPrint")
		reqId := ctxlogger.GetReqID(r.Context())

		// - Получаем код из body
		type AddCode_json struct {
			Gtin       string `json:"gtin"`
			Serial     string `json:"serial"`
			Сrypto     string `json:"crypto"`
			SourceName string `json:"source_name"`
		}

		code_json := AddCode_json{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&code_json)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedCode := entity.Code{
			Gtin:   code_json.Gtin,
			Serial: code_json.Serial,
			Crypto: code_json.Сrypto,
		}
		// Добавляем продукт
		err = usecase.AddCodeForPrint(r.Context(), mappedCode, code_json.SourceName)
		if err != nil {
			l.Warn("Ошибка добавления кода для печати", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		resp_body := toResponse(reqId, true, "Успешно", nil)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
