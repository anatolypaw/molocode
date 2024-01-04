package v1

import (
	"encoding/json"
	"fmt"
	"hub/internal/storage"
	"net/http"
)

// Устанавливает код произведенным
func SetCodeProduced(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Принимает json структуру
		type model struct {
			Gtin     string
			Serial   string
			Crypto   string
			Terminal string
			Proddate string
			Discard  bool
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

		// Передаем в хранилище
		err = s.SetCodeProduced(m.Gtin, m.Serial, m.Crypto, m.Terminal, m.Proddate, m.Discard)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Response(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, Response(true, "", nil))
	}
}
