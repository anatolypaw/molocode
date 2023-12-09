package model

import (
	"fmt"
	"regexp"
	"time"
)

// Когда и откуда был загружен код
type SourceInfo struct {
	Source string    // Откуда загружен, например с сервера "server main"
	Time   time.Time // Время получения кода
}

// Когда и куда был загружен код
type ReceiveInfo struct {
	Receiver string // Имя линии, получившей код
	Time     time.Time
}

// Информация о выгрузке в 1с
type UploadInfo struct {
	Time   time.Time
	Status string
}

// Когда и где код пошел в выпуск продукции. т.е. был  связан с единицей продукции
type ProducedInfo struct {
	Terminal string    // Имя линии фасовки, где он был нанесен или считан камерой
	Time     time.Time // Время, когда код был нанесен или считан на линии
	ProdDate string    `bson:",omitempty" json:",omitempty"` // Дата производства продукта 2023-10-09
	Discard  bool      // True - операция отбраковки кода
}

type Code struct {
	Serial       string         // Серийный номер, формат честного знака. Уникален для каждого кода с этим GTIN
	Crypto       string         // Криптохвост, формат честного знака
	SourceInfo   SourceInfo     `bson:",omitempty" json:",omitempty"` // Информация об источнике поступления кода
	ReceiveInfo  ReceiveInfo    `bson:",omitempty" json:",omitempty"` // Информация о получателе кода
	ProducedInfo []ProducedInfo `bson:",omitempty" json:",omitempty"` // Информация о его выпуске на линии фасовки
	UploadInfo   UploadInfo     `bson:",omitempty" json:",omitempty"` // Информация о выгрузке в 1с
}

// Продукт, gtin для каждого уникален.
type Good struct {
	Gtin        string    `bson:""` // gtin продукта
	Description string    `bson:""` // описание продукта
	StoreCount  uint      `bson:""` // сколько хранить кодов
	Get         bool      `bson:""` // флаг, получать коды из 1с
	Upload      bool      `bson:""` // флаг, выгружать коды в 1с
	Avaible     bool      `bson:""` // флаг, выдавать ли кода на терминал
	ShelfLife   uint      `bson:""` // срок годности продукта. Это нужно для того, что бы на линии фасовки вычислять конечную дату и печатать её на упаковке
	Created     time.Time `bson:""` // Дата создания продукта
	Codes       []Code    `bson:"" json:""`
}

// Проверяет корректность всех полей
func (g *Good) Validate() error {
	// Gtin
	err := ValidateGtin(g.Gtin)
	if err != nil {
		return err
	}

	// Description
	err = ValidateDescription(g.Description)
	if err != nil {
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

// Проверка корректности описания
func ValidateDescription(description string) error {
	const op = "entity.Good.ValidateDescription"

	// Проверяем корректность описания
	if len(description) == 0 {
		return fmt.Errorf("%s: Отсутствует описание", op)
	}

	return nil
}

// Проверяют корректность серийного номера, криптохвоста
func (code *Code) Validate() error {
	const op = "entity.Good.ValidateCode"

	if len(code.Serial) != 6 {
		return fmt.Errorf("%s: Некорректная длинна серийного номера", op)
	}

	if len(code.Crypto) != 4 {
		return fmt.Errorf("%s: Некорректная длинна криптохвоста", op)
	}

	return nil
}
