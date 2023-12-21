package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"storage/mongodb"
)

// Добавляет код в базу полученый из ГИС МТ для нанесения
// метод POST
func AddCodeForPrint(s *mongodb.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.AddCode"

		// Принимает json структуру
		type reqModel struct {
			SourceName string
			Gtin       string
			Serial     string
			Crypto     string
		}

		var m reqModel

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
		// Добавляем код
		err = s.AddCodeForPrint(m.Gtin, m.Serial, m.Crypto, m.SourceName)
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
