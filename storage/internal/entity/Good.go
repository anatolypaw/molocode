package entity

import (
	"fmt"
	"regexp"
)

// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin        string `bson:",omitempty"` // gtin продукта
	Description string `bson:",omitempty"` // описание продукта
	StoreCount  uint   `bson:",omitempty"` // сколько хранить кодов
	Get         bool   `bson:",omitempty"` // флаг, получать коды из 1с
	Upload      bool   `bson:",omitempty"` // флаг, выгружать коды в 1с
	Avaible     bool   `bson:",omitempty"` // флаг, выдавать ли кода на терминал
	ShelfLife   uint   `bson:",omitempty"` // срок годности продукта
}

// Проверка корректности gtin
func ValidateGtin(gtin string) error {
	const op = "entity.Good.ValidateGtin"

	// Проверяем корректность gtin
	matched, err := regexp.MatchString(`^0\d{13}$`, gtin)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		return err
	}

	if !matched {
		return fmt.Errorf("%s: Некорректный формат gtin '%s'", op, gtin)
	}

	return nil
}

// Проверяет корректность gtin
func (g *Good) ValidateGtin() error {
	const op = "entity.Good.ValidateGtin"

	// Проверяем корректность gtin
	err := ValidateGtin(g.Gtin)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		return err
	}
	return nil
}
