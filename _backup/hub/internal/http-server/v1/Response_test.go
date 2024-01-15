package v1

import (
	"errors"
	"fmt"
	"molocode/entity"
	"testing"
)

func TestResponse(t *testing.T) {
	tests := []struct {
		err  error
		data any
	}{
		{nil, nil},
		{nil, 5},
		{errors.New("error"), nil},
		{errors.New("error"), entity.Code{}},
	}

	for i, test := range tests {
		result := test
		fmt.Println(i, result)

	}
}
