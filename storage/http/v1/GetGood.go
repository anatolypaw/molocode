package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"storage/mongodb"
)

// Возвращает все продукты из базы
func GetGood(s *mongodb.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.GetGood"

		// Получаем продукты из хранилища
		result, err := s.GetGood()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		// Преобразуем ответ хранилища в json и передаем клиенту
		resultJson, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, string(resultJson))
	}
}
