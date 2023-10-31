package entity

import (
	"fmt"
	"regexp"
)

// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin        string `bson:"gtin,omitempty"`        // gtin продукта
	Description string `bson:"description,omitempty"` // описание продукта
	StoreCount  int    `bson:"store_count,omitempty"` // сколько хранить кодов
	Get         bool   `bson:"get,omitempty"`         // флаг, получать коды из 1с
	Upload      bool   `bson:"upload,omitempty"`      // флаг, выгружать коды в 1с
	Avaible     bool   `bson:"avaible,omitempty"`     // флаг, выдавать ли кода на терминал
	ShelfLife   int    `bson:"shelf_life,omitempty"`  // срок годности продукта
}

// Создает объект продукт, проверяет корректность gtin
func New(gtin string) (Good, error) {
	const op = "entity.Good.New"

	// Проверяем корректность gtin
	matched, err := regexp.MatchString(`^0\d{13}$`, gtin)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		return Good{}, err
	}

	if !matched {
		return Good{}, fmt.Errorf("%s: Некорректный формат gtin '%s'", op, gtin)
	}

	g := Good{
		Gtin: gtin,
	}
	return g, nil

}

// Установить описание
func (g *Good) EditDescription(desc string) {
	const op = "entity.Good.EditDescription"
	_ = op
	g.Description = desc
}
