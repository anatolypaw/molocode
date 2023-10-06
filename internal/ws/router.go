package ws

import (
	"fmt"
	"molocode/internal/ws/wapi"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

var tokenAuth *jwtauth.JWTAuth

func Router() http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil) // В проде установить секрет

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

	r := chi.NewRouter()

	//middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Защищенные wapi маршруты
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
	
		r.Get("/wapi/admin", func(w http.ResponseWriter, r *http.Request) {
		  _, claims, _ := jwtauth.FromContext(r.Context())
		  w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		}) 
		

	  })

	fs := http.FileServer(http.Dir("./www/build/"))

	// Маршруты с редиректом на авторизацию 
	r.Group(func(r chi.Router) {
		r.Use(UnloggedInRedirector)
		r.Get("/", fs.ServeHTTP)
		
	})

	// Публичные маршруты
	r.Group(func(r chi.Router) {
		r.Get("/*", fs.ServeHTTP)
		r.Get("/wapi/login", wapi.Test)
	})

	return r
}


func UnloggedInRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := jwtauth.FromContext(r.Context())

		if token == nil || jwt.Validate(token) != nil {
		http.Redirect(w, r, "/login.html", http.StatusFound)
			}

			next.ServeHTTP(w, r)
	})
}