package v1

import (
	"encoding/json"
	"fmt"
	"molocode/internal/entity"
	"net/http"
)

type IAdminUsecase interface {
	AddGood(entity.Good) error
	GetAllGoods() ([]entity.Good, error)
}

// Добавляет продукт, проверяя корректность GTIN  и отсутсвие записи с таким gtin
// метод POST
func AddGood(u IAdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

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
		err = u.AddGood(mappedGood)
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
func GetAllGoods(u IAdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Получаем продукты
		goods, err := u.GetAllGoods()
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
