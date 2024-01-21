package v1

import (
	"context"
	"encoding/json"
	"fmt"
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

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.AddGood")

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		good_dto := good_dto{}
		decoder := json.NewDecoder(r.Body)
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
			l.Warn("Продукт не добавлен", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		resp_body := toResponse(true, "Успешно", nil)
		l.Info("Продукт добавлен", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}

// Возвращает все продукты из базы
// метод POST
func GetAllGoods(usecase IAdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Подготовка логгера
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("func", "v1.GetAllGoods")

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
		l.Info("Успешно", "resp_body", resp_body)
		fmt.Fprint(w, resp_body)
	}
}
