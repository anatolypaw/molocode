package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/storage"
	"molocode/entity"
	"net/http"
)

// Устанавливает код произведенным
func SetCodeProduced(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Принимает json структуру
		type model struct {
			Gtin     string `json:"gtin"`
			Serial   string `json:"serial"`
			Crypto   string `json:"crypto"`
			Terminal string `json:"terminal"`
			Proddate string `json:"prod_date"`
			Discard  bool   `json:"discard"`
		}

		var m model

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

		// Передаем в хранилище
		err = s.SetCodeProduced(m.Gtin, m.Serial, m.Crypto, m.Terminal, m.Proddate, m.Discard)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, entity.ToResponse(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, entity.ToResponse(true, "", nil))
	}
}
