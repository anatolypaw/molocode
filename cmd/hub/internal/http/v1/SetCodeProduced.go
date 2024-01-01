package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"molocode/cmd/hub/internal/storage"
	"net/http"
)

// Устанавливает код произведенным
func SetCodeProduced(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.SetCodeProduced"

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
			err = fmt.Errorf("%s: %w", op, err)
			log.Print(err)
			fmt.Fprint(w, err)
			return
		}

		// Передаем в хранилище
		err = s.SetCodeProduced(m.Gtin, m.Serial, m.Crypto, m.Terminal, m.Proddate, m.Discard)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	}
}
