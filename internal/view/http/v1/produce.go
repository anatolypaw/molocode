package v1

import (
	"encoding/json"
	"fmt"
	"molocode/internal/app/ctxlogger"
	"molocode/internal/app/usecase/usecase_produce"
	"net/http"
)

// Возвращает код для печати
func GetCodeForPrint(usecase usecase_produce.ProduceUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.GetCodeForPrint")
		reqId := ctxlogger.GetReqID(r.Context())

		// - Получаем информацию о запрашиваемом коде из body
		type Req_json struct {
			Gtin         string `json:"gtin"`
			TerminalName string `json:"terminal_name"`
		}

		req_json := Req_json{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&req_json)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// Запрашиваем КМ для печати
		codeForPrint, err := usecase.GetCodeForPrint(r.Context(), req_json.Gtin, req_json.TerminalName)
		if err != nil {
			l.Error("Ошибка получения кода для печати", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// MAPPING
		type CodeForPrint_json struct {
			Gtin    string `json:"gtin"`
			Serial  string `json:"serial"`
			Crypto  string `json:"crypto"`
			PrintId uint64 `json:"print_id"`
		}

		codeForPrint_json := CodeForPrint_json{
			Gtin:    codeForPrint.Code.Gtin,
			Serial:  codeForPrint.Code.Serial,
			Crypto:  codeForPrint.Code.Crypto,
			PrintId: codeForPrint.PrintId,
		}

		resp_body := toResponse(reqId, true, "Успешно", codeForPrint_json)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}

// Отмечает напечатанный код произведенным
func ProducePrinted(usecase usecase_produce.ProduceUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.ProducePrinted")
		reqId := ctxlogger.GetReqID(r.Context())

		// - Получаем информацию о коде из body
		type Req_json struct {
			Gtin         string `json:"gtin"`
			Serial       string `json:"serial"`
			TerminalName string `json:"terminal_name"`
			ProdDate     string `json:"prod_date"`
		}

		req_json := Req_json{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&req_json)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// Запрашиваем КМ для печати
		codeForPrint, err := usecase.GetCodeForPrint(r.Context(), req_json.Gtin, req_json.TerminalName)
		if err != nil {
			l.Error("Ошибка получения кода для печати", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		resp_body := toResponse(reqId, true, "Успешно", nil)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
