package v1

import (
	"fmt"
	"net/http"
)

// Возвращает  ok
func Test(w http.ResponseWriter, r *http.Request) {
	const op = "ts.v1.Test"
	fmt.Fprint(w, op)
}
