package entity

import (
	"fmt"
	"regexp"
	"time"
)

// Продукт, gtin для каждого уникален.
type Good struct {
	Gtin            string    `bson:"_id" json:"gtin"`                              // gtin продукта
	Desc            string    `bson:"desc" json:"desc"`                             // описание продукта
	StoreCount      uint      `bson:"store_count" json:"store_count"`               // сколько хранить кодов
	GetCodeForPrint bool      `bson:"get_code_for_print" json:"get_code_for_print"` // флаг, что разрешено получение кодов для нанесения
	AllowProduce    bool      `bson:"allow_produce" json:"allow_produce"`           // флаг, что разрешено производство
	Upload          bool      `bson:"upload" json:"upload"`                         // флаг, выгружать коды в 1с
	ShelfLife       uint      `bson:"shelf_life" json:"shelf_life"`                 // срок годности продукта. Это нужно для того, что бы на линии фасовки вычислять конечную дату и печатать её на упаковке
	Created         time.Time `bson:"created" json:"created"`                       // Дата создания продукта
}

// Проверяет корректность всех полей
func (g *Good) Validate() error {
	// Gtin
	if err := ValidateGtin(g.Gtin); err != nil {
		return err
	}
	// Description

	if err := ValidateDescription(g.Desc); err != nil {
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
