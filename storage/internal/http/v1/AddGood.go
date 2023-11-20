package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"storage/internal/domain/models"
	"storage/internal/storage/mongodb"
)

// Добавляет продукт, проверяя корректность GTIN  и отсутсвие записи с таким gtin
// метод POST
// Принимает json
/*	{
		"gtin": "04600000000000",
		"desc": "Молоко 3,5%"
	}
*/

func AddGood(s *mongodb.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.AddGood"

		// Устанавливаем время создания
		good := models.Good{}

		// Декодируем полученный в теле json
		// Разрешить только поля, укаказанные в entity.Good
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		// Добавляем продукт в хранилище
		res, err := s.AddGood(good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		// Преобразуем ответ хранилища в json и передаем клиенту
		resJson, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, string(resJson))
	}
}
