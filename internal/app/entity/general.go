package entity

import (
	"errors"
	"regexp"
)

func ValidateGtin(gtin string) error {
	matched, err := regexp.MatchString(`^0\d{13}$`, gtin)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("некорректный формат gtin")
	}
	return nil
}
