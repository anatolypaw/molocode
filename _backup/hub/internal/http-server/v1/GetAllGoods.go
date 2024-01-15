package v1

import (
	"fmt"
	"hub/internal/storage"
	"molocode/entity"
	"net/http"
)

// Возвращает все продукты из базы
func GetAllGoods(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// Получаем продукты из хранилища
		result, err := s.GetGoods()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, entity.ToResponse(false, err.Error(), nil))
			return
		}

		fmt.Fprint(w, entity.ToResponse(true, "", result))
	}
}
