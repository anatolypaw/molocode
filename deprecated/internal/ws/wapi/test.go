package wapi

import (
	"fmt"
	"net/http"
)

// Возвращает  ok
func Test(w http.ResponseWriter, r *http.Request) {
	const op = "ws.wapi.Test"
	fmt.Fprint(w, op)
}
