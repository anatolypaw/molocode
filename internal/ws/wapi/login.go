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

		var user structs.User

		//Считываем из тела запроса логин и пароль
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			err := fmt.Sprintf("%s: %s", op, err)
			http.Error(w, err, http.StatusBadRequest)
			return
		}
		
		// Запрашиваем пользователя из базы
		user, err = s.GetUserByLoginPass(user.Login, user.Password)
		if err != nil {
			http.Error(w, "login/password incorrect", http.StatusBadRequest)
			return 
		}

		//Если пользователь с парой логин/пароль найден, то возвращаем JWT токен
		_, token, err := tokenAuth.Encode(map[string]interface{}{"login": user.Login, "role": user.Role})
		if err != nil {
			err := fmt.Sprintf("%s: %s", op, err)
			http.Error(w, err, http.StatusBadRequest)
			return
		}

		log.Printf("%s User %s succseful authorized", op, user.Login)
		fmt.Fprint(w, token)
}
}