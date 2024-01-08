package v1

import (
	"fmt"
	"net/http"
)

func Mock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.v1.Mock"

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n%s", op, r.URL)
	}
}
