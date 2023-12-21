package model

import (
	"fmt"
	"regexp"
	"time"
)

// Продукт, gtin для каждого уникален.
type Good struct {
	Gtin           string    `bson:"_id"` // gtin продукта
	Description    string    `bson:""`    // описание продукта
	StoreCount     uint      `bson:""`    // сколько хранить кодов
	AcceptForPrint bool      `bson:""`    // флаг, что разрешено получение кодов для нанесения
	SendForPrint   bool      `bson:""`    // флаг, что продукт доступен для печати
	Upload         bool      `bson:""`    // флаг, выгружать коды в 1с
	Avaible        bool      `bson:""`    // флаг, выдавать ли кода на терминал
	ShelfLife      uint      `bson:""`    // срок годности продукта. Это нужно для того, что бы на линии фасовки вычислять конечную дату и печатать её на упаковке
	Created        time.Time `bson:""`    // Дата создания продукта
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

// Проверка корректности описания продукта
func ValidateDescription(description string) error {
	const op = "entity.Good.ValidateDescription"

	// Проверяем корректность описания
	if len(description) == 0 {
		return fmt.Errorf("%s: Отсутствует описание", op)
	}

	return nil
}
