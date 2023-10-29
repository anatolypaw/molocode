package v1

import (
	"fmt"
	"net/http"
	"storage/internal/storage"
)

// Добавляет продукт
func AddGood(s *storage.Storage) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		const op = "http.v1"
		res := s.AddGood("0460123123123", "Описание")
		fmt.Fprint(w, res)
	}
} 