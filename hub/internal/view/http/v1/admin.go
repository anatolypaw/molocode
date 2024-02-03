package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/ctxlogger"
	"hub/internal/entity"
	"hub/internal/usecase/admin"

	"net/http"
)

// Добавляет продукт
// метод POST
func AddGood(u admin.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.AddGood")
		reqId := ctxlogger.GetReqID(r.Context())

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		good_dto := good_dto{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good_dto)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedGood := entity.Good{
			Gtin:            good_dto.Gtin,
			Desc:            good_dto.Desc,
			StoreCount:      good_dto.StoreCount,
			GetCodeForPrint: good_dto.GetCodeForPrint,
			AllowProduce:    good_dto.AllowProduce,
			AllowPrint:      good_dto.AllowPrint,
			Upload:          good_dto.Upload,
		}

		// Добавляем продукт в хранилище
		err = u.AddGood(r.Context(), mappedGood)
		if err != nil {
			l.Warn("Продукт не добавлен", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		resp_body := toResponse(reqId, true, "Успешно", nil)
		l.Info("Продукт добавлен", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}

// Возвращает все продукты из базы
// метод POST
func GetAllGoods(usecase admin.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.GetAllGoods")
		reqId := ctxlogger.GetReqID(r.Context())

		// Получаем продукты
		goods, err := usecase.GetAllGoods(r.Context())
		if err != nil {
			l.Error("Ошибка получения продуктов из базы", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(reqId, false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedGoods := []good_dto{}
		for _, good := range goods {
			mappedGoods = append(mappedGoods, good_dto{
				Gtin:            string(good.Gtin),
				Desc:            good.Desc,
				StoreCount:      good.StoreCount,
				GetCodeForPrint: good.GetCodeForPrint,
				AllowProduce:    good.AllowProduce,
				AllowPrint:      good.AllowPrint,
				Upload:          good.Upload,
				CreatedAt:       good.CreatedAt,
			})
		}

		resp_body := toResponse(reqId, true, "Успешно", mappedGoods)
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
