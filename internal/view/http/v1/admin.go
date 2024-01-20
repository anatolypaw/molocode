package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"molocode/internal/app/ctxlogger"
	"molocode/internal/app/entity"
	"net/http"
)

type IAdminUsecase interface {
	AddGood(context.Context, entity.Good) error
	GetAllGoods(context.Context) ([]entity.Good, error)
}

// Добавляет продукт
// метод POST
func AddGood(usecase IAdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Считываем body для использования в логгере
		body_bytes, _ := io.ReadAll(r.Body)
		body := string(body_bytes)
		l := ctxlogger.LoggerFromContext(r.Context())
		l.Info("Добавление продукта", "function", "v1.AddGood", "req_body", body)

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		good_dto := good_dto{}
		decoder := json.NewDecoder(bytes.NewReader(body_bytes))
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good_dto)
		if err != nil {
			l.Error("Json decoder", "error", err)

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedGood := entity.Good{
			Gtin:            good_dto.Gtin,
			Desc:            good_dto.Desc,
			StoreCount:      good_dto.StoreCount,
			GetCodeForPrint: good_dto.GetCodeForPrint,
			AllowProduce:    good_dto.AllowProduce,
			Upload:          good_dto.Upload,
		}

		// Добавляем продукт в хранилище
		err = usecase.AddGood(r.Context(), mappedGood)
		if err != nil {
			l.Warn("Продукт не добавлен", "error", err, "body", body_bytes)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, toResponse(true, "Продукт добавлен", nil))
	}
}

// Возвращает все продукты из базы
// метод POST
func GetAllGoods(usecase IAdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Logger
		l := ctxlogger.LoggerFromContext(r.Context())
		l.Info("Запрос всех продуктов")

		// Получаем продукты
		goods, err := usecase.GetAllGoods(r.Context())
		if err != nil {
			l.Error("Ошибка получения продуктов из базы", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
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
				Upload:          good.Upload,
				CreatedAt:       good.CreatedAt,
			})
		}

		resp_body := toResponse(true, "Успешно", mappedGoods)
		l.Info("Запрос всех продуктов", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
