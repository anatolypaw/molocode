package model

import (
	"fmt"
	"regexp"
	"time"
)

// Когда и откуда был загружен код
type SourceInfo struct {
	Name string    `bson:"n" json:",omitempty"` // Откуда загружен, например с сервера "server main"
	Time time.Time `bson:"t" json:",omitempty"` // Время получения кода
}

// Когда и где код пошел в выпуск продукции. т.е. был  связан с единицей продукции
type ProducedInfo struct {
	Terminal string    `bson:"tr" json:",omitempty"` // Имя линии фасовки, где он был нанесен или считан камерой
	Time     time.Time `bson:"tm" json:",omitempty"` // Время, когда код был нанесен или считан на линии
	ProdDate string    `bson:"pd" json:",omitempty"` // Дата производства продукта 2023-10-09
	Discard  bool      `bson:"d" json:",omitempty"`  // True - операция отбраковки кода
}

// Информация, связанная с печатью
type PrintInfo struct {
	ReadyForPrint bool   `bson:"r" json:",omitempty"`
	PrintId       uint64 `bson:"i" json:",omitempty"` // Уникальный номер для кода.
}

// Информация о выгрузке в 1с
type UploadInfo struct {
	Time   time.Time `bson:"tm" json:",omitempty"`
	Status string    `bson:"s" json:",omitempty"`
}

type Code struct {
	Serial       string         `bson:"sn" json:",omitempty"`  // Серийный номер, формат честного знака. Уникален для каждого кода с этим GTIN
	Crypto       string         `bson:"cn" json:",omitempty"`  // Криптохвост, формат честного знака
	SourceInfo   SourceInfo     `bson:"si" json:",omitempty"`  // Информация об источнике поступления кода
	PrintInfo    PrintInfo      `bson:"pti" json:",omitempty"` // Информация, связанная с печатью
	ProducedInfo []ProducedInfo `bson:"pdi" json:",omitempty"` // Информация о его выпуске на линии фасовки
	UploadInfo   UploadInfo     `bson:"ui" json:",omitempty"`  // Информация о выгрузке в 1с
}

// Продукт, gtin для каждого уникален.
type Good struct {
	Gtin        string    `bson:"_id"` // gtin продукта
	Description string    `bson:""`    // описание продукта
	StoreCount  uint      `bson:""`    // сколько хранить кодов
	Get         bool      `bson:""`    // флаг, получать коды из 1с
	Upload      bool      `bson:""`    // флаг, выгружать коды в 1с
	Avaible     bool      `bson:""`    // флаг, выдавать ли кода на терминал
	ShelfLife   uint      `bson:""`    // срок годности продукта. Это нужно для того, что бы на линии фасовки вычислять конечную дату и печатать её на упаковке
	Created     time.Time `bson:""`    // Дата создания продукта
	Codes       []Code    `bson:"" json:",omitempty"`
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

// Проверка корректности описания
func ValidateDescription(description string) error {
	const op = "entity.Good.ValidateDescription"

	// Проверяем корректность описания
	if len(description) == 0 {
		return fmt.Errorf("%s: Отсутствует описание", op)
	}

	return nil
}

func ValidateSerial(serial string) error {
	const op = "entity.Good.ValidateSerial"
	if len(serial) != 6 {
		return fmt.Errorf("%s: Некорректная длинна серийного номера", op)
	}

	return nil
}

func ValidateCrypto(crypto string) error {
	const op = "entity.Good.ValidateCrypto"
	if len(crypto) != 4 {
		return fmt.Errorf("%s: Некорректная длинна криптохвоста", op)
	}

	return nil
}

// Проверяют корректность серийного номера, криптохвоста
func (code *Code) ValidateSerialCrypto() error {
	if err := ValidateSerial(code.Serial); err != nil {
		return err
	}

	if err := ValidateCrypto(code.Crypto); err != nil {
		return err
	}
	return nil
}
