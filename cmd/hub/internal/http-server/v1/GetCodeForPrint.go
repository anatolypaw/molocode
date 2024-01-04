package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/storage"
	"net/http"
)

// Возвращает код для печати
func GetCodeForPrint(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

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
			fmt.Fprint(w, Response(false, err.Error(), nil))
			return
		}
		// Получаем продукты из хранилища
		result, err := s.GetCodeForPrint(m.Gtin, m.Terminal)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Response(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, Response(true, "", result))
	}
}
