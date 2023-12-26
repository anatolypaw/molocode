package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"storage/storage"
)

// Возвращает продукт и сколько кодов требуется для этого продукта до нормы
// для получения из 1с
func GetReqCodeCount(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.GetReqCodeCount"

		// Отдает json структуру
		type model struct {
			Gtin     string
			Terminal string
		}

		var m model
		_ = m

		// Получаем продукты и требуемое количество кодов из хранилища
		result, err := s.GetReqCodeCount()
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
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resultJson))
	}
}
