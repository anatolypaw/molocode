package entity

import (
	"errors"
	"regexp"
)

type Gtin string
type Serial string
type Crypto string

func (g Gtin) Validate() error {
	matched, err := regexp.MatchString(`^0\d{13}$`, string(g))
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("некорректный формат gtin")
	}
	return nil
}

func (s Serial) Validate() error {
	if len(s) != 6 {
		return errors.New("некорректная длинна serial")
	}
	return nil
}

func (s Crypto) Validate() error {
	if len(s) != 4 {
		return errors.New("некорректная длинна crypto")
	}
	return nil
}
