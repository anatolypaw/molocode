package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"molocode/internal/app/entity"
	"net/http"
)

type IAdminUseCase interface {
	AddGood(context.Context, entity.Good) error
	GetAllGoods(context.Context) ([]entity.Good, error)
}

// Добавляет продукт
// метод POST
func AddGood(ctx context.Context, usecase IAdminUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		ctx = entity.WithNewReqID(ctx)

		good := good_dto{}

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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
		err = usecase.AddGood(ctx, mappedGood)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, toResponse(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, toResponse(true, "Продукт добавлен", nil))
	}
}

// Возвращает все продукты из базы
// метод POST
func GetAllGoods(ctx context.Context, usecase IAdminUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		ctx = entity.WithNewReqID(ctx)

		// Получаем продукты
		goods, err := usecase.GetAllGoods(ctx)
		if err != nil {
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
