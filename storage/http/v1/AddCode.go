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
// Принимает json
/*	{
		"sourceName":"1c service",
		"gtin": "04600000000000",
		"serial": "abcdef",
		"crypto": "1234"
	}
*/

func AddCode(s *mongodb.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.AddCode"

		type reqModel struct {
			SourceName string
			Gtin       string
			Serial     string
			Crypto     string
		}

		var rm reqModel

		// Декодируем полученный json
		// Разрешить только поля, укаказанные в структуре
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&rm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			log.Print(err)
			fmt.Fprint(w, err)
			return
		}
		// Добавляем код
		err = s.AddCode(rm.Gtin, rm.Serial, rm.Crypto, rm.SourceName)
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
