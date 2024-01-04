package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/storage"
	"molocode/entity"
	"net/http"
)

// Добавляет продукт, проверяя корректность GTIN  и отсутсвие записи с таким gtin
// метод POST
// Принимает json
/*	{
		"gtin": "04600000000000",
		"description": "Молоко 3,5%"
	}
*/

func AddGood(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		good := entity.Good{}

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Response(false, err.Error(), nil))
			return
		}

		// Добавляем продукт в хранилище
		result, err := s.AddGood(good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Response(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, Response(true, "Продукт добавлен", result))
	}
}
