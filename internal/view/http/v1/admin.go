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

		// Logger
		const op = "v1.AddGood"
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("function", op)

		good := good_dto{}

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			l.Warn("Json decoder", "error", err)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		// MAPPING
		mappedGood := entity.Good{
			Gtin:            good.Gtin,
			Desc:            good.Desc,
			StoreCount:      good.StoreCount,
			GetCodeForPrint: good.GetCodeForPrint,
			AllowProduce:    good.AllowProduce,
			Upload:          good.Upload,
		}

		// Добавляем продукт в хранилище
		err = usecase.AddGood(r.Context(), mappedGood)
		if err != nil {
			l.Warn("Ошибка добавления продукта", "error", err)
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
		const op = "v1.GetAllGoods"
		l := ctxlogger.LoggerFromContext(r.Context())
		l = l.With("function", op)

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

		fmt.Fprint(w, toResponse(true, "Успешно", mappedGoods))
	}
}
