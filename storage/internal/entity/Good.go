package entity

import (
	"fmt"
	"regexp"
)

type Code struct {
	Serial string // Серийный номер, формат честного знака
	Crypto string // Криптохвост, формат честного знака

}

// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin        string `bson:""` // gtin продукта
	Description string `bson:""` // описание продукта
	StoreCount  uint   `bson:""` // сколько хранить кодов
	Get         bool   `bson:""` // флаг, получать коды из 1с
	Upload      bool   `bson:""` // флаг, выгружать коды в 1с
	Avaible     bool   `bson:""` // флаг, выдавать ли кода на терминал
	ShelfLife   uint   `bson:""` // срок годности продукта
	Codes       []Code `json:",omitempty"`
}

// Проверка корректности gtin как отдельная функция
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

// Проверяет корректность внутри структуры gtin
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
