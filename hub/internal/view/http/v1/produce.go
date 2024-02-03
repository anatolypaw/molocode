package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/ctxlogger"
	"hub/internal/usecase/produce"

	"net/http"
)

// Возвращает код для печати
func GetCodeForPrint(u produce.ProduceUsecase) http.HandlerFunc {
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

		req := Req_json{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&req)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// Запрашиваем КМ для печати
		code, err := u.GetCodeForPrint(r.Context(),
			req.Gtin,
			req.TerminalName)
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
			Gtin:    code.Code.Gtin,
			Serial:  code.Code.Serial,
			Crypto:  code.Code.Crypto,
			PrintId: code.PrintId,
		}

		resp_body := toResponse(reqId, true, "Успешно", codeForPrint_json)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}

// Отмечает напечатанный код произведенным
func ProducePrinted(usecase produce.ProduceUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.ProducePrinted")
		reqId := ctxlogger.GetReqID(r.Context())

		// - Получаем информацию о коде из body
		type Req struct {
			Gtin         string `json:"gtin"`
			Serial       string `json:"serial"`
			TerminalName string `json:"terminal_name"`
			ProdDate     string `json:"prod_date"`
		}

		var req Req
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&req)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// - Обращаемся к usecase
		err = usecase.ProducePrinted(r.Context(),
			req.Gtin,
			req.Serial,
			req.TerminalName,
			req.ProdDate)

		if err != nil {
			l.Error("Ошибка", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		resp_body := toResponse(reqId, true, "Успешно", nil)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
