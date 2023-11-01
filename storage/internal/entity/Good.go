package entity

import (
	"fmt"
	"regexp"
)

// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin        string // gtin продукта
	Description string // описание продукта
	StoreCount  uint   // сколько хранить кодов
	Get         bool   // флаг, получать коды из 1с
	Upload      bool   // флаг, выгружать коды в 1с
	Avaible     bool   // флаг, выдавать ли кода на терминал
	ShelfLife   uint   // срок годности продукта
}

// Создает объект продукт, проверяет корректность gtin
func (g *Good) ValidateGtin() error {
	const op = "entity.Good.New"

	// Проверяем корректность gtin
	matched, err := regexp.MatchString(`^0\d{13}$`, g.Gtin)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		return err
	}

	if !matched {
		return fmt.Errorf("%s: Некорректный формат gtin '%s'", op, g.Gtin)
	}

	return nil

}
