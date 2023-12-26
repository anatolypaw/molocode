package model

import (
	"fmt"
	"regexp"
	"time"
)

// Продукт, gtin для каждого уникален.
type Good struct {
	Gtin            string    `bson:"_id"`             // gtin продукта
	Description     string    `bson:"Description"`     // описание продукта
	StoreCount      uint      `bson:"StoreCount"`      // сколько хранить кодов
	GetCodeForPrint bool      `bson:"GetCodeForPrint"` // флаг, что разрешено получение кодов для нанесения
	AllowProduce    bool      `bson:"AllowProduce"`    // флаг, что разрешено производство
	Upload          bool      `bson:"Upload"`          // флаг, выгружать коды в 1с
	ShelfLife       uint      `bson:"ShelfLife"`       // срок годности продукта. Это нужно для того, что бы на линии фасовки вычислять конечную дату и печатать её на упаковке
	Created         time.Time `bson:"Created"`         // Дата создания продукта
}

// Проверяет корректность всех полей
func (g *Good) Validate() error {
	// Gtin
	if err := ValidateGtin(g.Gtin); err != nil {
		return err
	}
	// Description

	if err := ValidateDescription(g.Description); err != nil {
		return err
	}

	return nil
}

// Проверка корректности gtin
func ValidateGtin(gtin string) error {
	const op = "model.Good.ValidateGtin"

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

// Проверка корректности описания продукта
func ValidateDescription(description string) error {
	const op = "model.Good.ValidateDescription"

	// Проверяем корректность описания
	if len(description) == 0 {
		return fmt.Errorf("%s: Отсутствует описание", op)
	}

	return nil
}
