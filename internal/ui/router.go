package ui

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
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

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
	
		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)
	
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		  _, claims, _ := jwtauth.FromContext(r.Context())
		  w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})


	  })
	 
		// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome anonymous"))
		})
	})

//	fs := http.FileServer(http.Dir("./www/build/"))
//	uiRouter.Handle("/*", fs)

//	uiRouter.Get("/goods", ui.GetGoods)

	return r
}