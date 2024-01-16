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

func ValidateSerial(serial string) error {
	if len(serial) != 6 {
		return errors.New("некорректная длинна serial")
	}
	return nil
}

func ValidateCrypto(crypto string) error {
	if len(crypto) != 4 {
		return errors.New("некорректная длинна crypto")
	}
	return nil
}
