package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/storage"
	"molocode/entity"
	"net/http"
)

// Добавляет код в базу полученый из ГИС МТ для нанесения
// метод POST
func AddCodesForPrint(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Принимает json структуру
		type reqModel struct {
			SourceName string `json:"source_name"`
			Gtin       string `json:"gtin"`
			Serial     string `json:"serial"`
			Crypto     string `json:"crypto"`
		}

		var m reqModel

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в структуре
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, entity.ToResponse(false, err.Error(), nil))
			return
		}

		// Добавляем код
		err = s.AddCodeForPrint(m.Gtin, m.Serial, m.Crypto, m.SourceName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, entity.ToResponse(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, entity.ToResponse(true, "Код успешно добавлен для печати", m))
	}
}
