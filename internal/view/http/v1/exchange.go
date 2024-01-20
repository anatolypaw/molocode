package v1

import (
	"context"
	"fmt"
	"molocode/internal/app/ctxlogger"
	"molocode/internal/app/usecase/usecase_exchange"
	"net/http"
)

type IExchangeUsecase interface {
	GetGoodsReqCodes(ctx context.Context) ([]usecase_exchange.CodeReq, error)
}

type codeReq_json struct {
	Gtin     string `json:"gtin"`
	Desc     string `json:"desc"`
	Required uint   `json:"required"`
}

// Добавляет продукт
// метод POST
func GetGoodsReqCodes(usecase IExchangeUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Logger
		l := ctxlogger.LoggerFromContext(r.Context())
		l.Info("Запрос требуемого количества кодов", "func", "v1.GetGoodsReqCodes")

		// Получаем список продуктов и требуемое количество кодов
		codereq, err := usecase.GetGoodsReqCodes(r.Context())
		if err != nil {
			l.Warn("Ошибка запроса продуктов", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
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
		l.Info("Получены продукты, требующие загрузки кодов", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
