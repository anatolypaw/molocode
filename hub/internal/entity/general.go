package entity

import (
	"errors"
	"regexp"
)

func ValidateGtin(gtin string) error {
	r := regexp.MustCompile(`^0\d{13}$`)
	if !r.MatchString(gtin) {
		return errors.New("некорректный формат gtin")
	}
	return nil
}
