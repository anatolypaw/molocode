package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"storage/entity"
	"storage/internal/storage"
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
		const op = "http.v1.AddGood"

		good := entity.Good{}

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в entity.Good
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			log.Print(err)
			fmt.Fprint(w, err)
			return
		}

		// Добавляем продукт в хранилище
		result, err := s.AddGood(good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		// Хранилище возвращает информацию о продукте, которая была добавлена
		// Преобразуем ответ хранилища в json и передаем клиенту
		resultJson, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resultJson))
	}
}
