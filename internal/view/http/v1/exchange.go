package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"molocode/internal/app/ctxlogger"
	"molocode/internal/app/entity"
	"molocode/internal/app/usecase/usecase_exchange"
	"net/http"
)

type IExchangeUsecase interface {
	GetGoodsReqCodes(ctx context.Context) ([]usecase_exchange.CodeReq, error)
	AddCodeForPrint(ctx context.Context, code entity.Code, source string) error
}

// Добавляет продукт
// метод POST
func GetGoodsReqCodes(usecase IExchangeUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.GetGoodsReqCodes")

		// Получаем список продуктов и требуемое количество кодов
		codereq, err := usecase.GetGoodsReqCodes(r.Context())
		if err != nil {
			l.Warn("Ошибка запроса продуктов", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
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

		resp_body := toResponse(true, "Успешно", mappedCodeReq)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}

// Доабвляет код для печати
func AddCodeForPrint(usecase IExchangeUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.AddCodeForPrint")

		// - Получаем код из body
		type AddCode_json struct {
			Gtin       string `json:"gtin"`
			Serial     string `json:"serial"`
			Srypto     string `json:"crypto"`
			SourceName string `json:"source_name"`
		}

		code_json := AddCode_json{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&code_json)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedCode := entity.Code{
			Gtin:   code_json.Gtin,
			Serial: code_json.Serial,
			Crypto: code_json.Srypto,
		}
		// Добавляем продукт
		err = usecase.AddCodeForPrint(r.Context(), mappedCode, code_json.SourceName)
		if err != nil {
			l.Warn("Ошибка добавления кода для печати", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		resp_body := toResponse(true, "Успешно", nil)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
