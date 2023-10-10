package wapi

import (
	"fmt"
	"net/http"
)

//
type User struct{
	Name string `json:"name"`
	Pass string `json:"pass"`
}


func Login(w http.ResponseWriter, r *http.Request){
	const op = "ws.wapi.Login"
	fmt.Fprint(w, op)

}