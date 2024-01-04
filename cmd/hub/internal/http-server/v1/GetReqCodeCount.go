package v1

import (
	"fmt"
	"hub/internal/storage"
	"net/http"
)

// Возвращает продукт и сколько кодов требуется для этого продукта до нормы
// для получения из 1с
func GetReqCodeCount(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Получаем продукты и требуемое количество кодов из хранилища
		result, err := s.GetReqCodeCount()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, Response(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, Response(true, "", result))
	}
}
