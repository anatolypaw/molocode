package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"storage/internal/adapter/storage"
	"storage/internal/usecase"
)

// Добавляет продукт, метод POST
// Принимает json
/*	{
		"gtin": "04600000000000",
		"desc": "Молоко 3,5%"
	}
*/
func AddGood(s *storage.Connection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.AddGood"

		var good struct {
			Gtin string
			Desc string
		}

		// Декодируем полученный в теле json
		err := json.NewDecoder(r.Body).Decode(&good)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}

		err = usecase.AddGood(s, good.Gtin, good.Desc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("%s: %w", op, err)
			fmt.Fprint(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "OK")
	}
}