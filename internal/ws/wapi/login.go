package wapi

import (
	"encoding/json"
	"fmt"
	"log"
	"molocode/internal/storage"
	"molocode/internal/structs"
	"net/http"

	"github.com/go-chi/jwtauth"
)

// Получает login и password, проверяет пароль в базе и считывает роль, возвращает
// JWT токен с логином и ролью
func Login(s *storage.Storage, tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		const op = "ws.wapi.Login"

		var u structs.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			err := fmt.Sprintf("%s: %s", op, err)
			http.Error(w, err, http.StatusBadRequest)
			return
		}

		log.Print(u)
}
}