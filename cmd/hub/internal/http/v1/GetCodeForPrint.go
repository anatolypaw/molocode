package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"molocode/cmd/hub/internal/storage"
	"net/http"
)

// Возвращает код для печати
func GetCodeForPrint(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.GetCodeForPrint"

		// Принимает json структуру
		type model struct {
			Gtin     string
			Terminal string
		}

		var m model

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в структуре
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			log.Print(err)
			fmt.Fprint(w, err)
			return
		}
		// Получаем продукты из хранилища
		result, err := s.GetCodeForPrint(m.Gtin, m.Terminal)
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
