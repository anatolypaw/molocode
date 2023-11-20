package models

import (
	"fmt"
	"regexp"
	"time"
)

// Информация о том, когда и откуда был загружен код
type LoadedInfo struct {
	Source string    // Откуда загружен, например с сервера "server main"
	Time   time.Time // Время получения кода
}

// Информация о том, когда и где он был нанесен на продукцию или считан с неё
type ProducedInfo struct {
	Terminal string    // Имя линии фасовки, где он был нанесен или произведен
	Time     time.Time // Время, когда код был нанесен или считан на линии
	ProdDate string    `bson:",omitempty json:",omitempty"` // Дата производства продукта 2023-10-09
}

type Code struct {
	Id       uint64         `bson:",omitempty json:",omitempty""` // id кода, присваивается только для кодов, которые идут на печать
	Serial   string         // Серийный номер, формат честного знака
	Crypto   string         // Криптохвост, формат честного знака
	Loaded   LoadedInfo     `bson:",omitempty json:",omitempty""` // Информация об источнике поступления кода
	Produced []ProducedInfo `bson:",omitempty json:",omitempty"`  // Информация о его выпуске на линии фасовки
}

// Продукт, gtin для каждого уникален.
type Good struct {
	Gtin        string    `bson:""` // gtin продукта
	Description string    `bson:""` // описание продукта
	StoreCount  uint      `bson:""` // сколько хранить кодов
	Get         bool      `bson:""` // флаг, получать коды из 1с
	Upload      bool      `bson:""` // флаг, выгружать коды в 1с
	PrintType   string    `bson:""` // тип нанесения кода. Самостоятельно наносим или уже нанесен print,  printed
	Avaible     bool      `bson:""` // флаг, выдавать ли кода на терминал
	ShelfLife   uint      `bson:""` // срок годности продукта
	Created     time.Time `bson:""` // Дата создания продукта
	Codes       []Code    `bson:"" json:""`
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

	// Проверяем корректность gtin
	err := ValidateGtin(g.Gtin)
	if err != nil {
		return err
	}
	return nil
}

// Проверяют корректность серийного номера, криптохвоста, источника
func (code *Code) Validate() error {
	const op = "entity.Good.ValidateCode"

	if len(code.Serial) != 6 {
		return fmt.Errorf("%s: Некорректная длинна серийного номера", op)
	}

	if len(code.Crypto) != 4 {
		return fmt.Errorf("%s: Некорректная длинна криптохвоста", op)
	}

	if code.Loaded.Source == "" {
		return fmt.Errorf("%s: Не указан источник кода", op)
	}

	return nil
}
