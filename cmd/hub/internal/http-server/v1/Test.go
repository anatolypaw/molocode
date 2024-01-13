package v1

import (
	"fmt"
	"molocode/entity"
	"net/http"
)

func Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		type Test struct {
			Name  string `json:"name"`
			Count uint   `json:"count"`
		}

		result := Test{
			Name:  "golang",
			Count: 99999,
		}

		fmt.Fprint(w, entity.ToResponse(true, "TEST", result))
	}
}
