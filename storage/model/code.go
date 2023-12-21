package model

import (
	"fmt"
	"time"
)

type Code struct {
	Serial       string         `bson:"_id" json:",omitempty"` // Серийный номер, формат честного знака. Уникален для каждого кода с этим GTIN
	Crypto       string         `bson:"" json:",omitempty"`    // Криптохвост, формат честного знака
	SourceInfo   SourceInfo     `bson:"" json:",omitempty"`    // Информация об источнике поступления кода
	PrintInfo    PrintInfo      `bson:"" json:",omitempty"`    // Информация, связанная с печатью
	ProducedInfo []ProducedInfo `bson:"" json:",omitempty"`    // Информация о его выпуске на линии фасовки
	UploadInfo   UploadInfo     `bson:"" json:",omitempty"`    // Информация о выгрузке в 1с
}

// Информация о коде, передаваемая на терминал для нанесения
type CodeForPrint struct {
	Gtin    string
	Serial  string
	Crypto  string
	PrintId uint64
}

// Когда и откуда был загружен код
type SourceInfo struct {
	Name string    `bson:"" json:",omitempty"` // Откуда загружен, например с сервера "server main"
	Time time.Time `bson:"" json:",omitempty"` // Время получения кода
}

// Когда и где код пошел в выпуск продукции. т.е. был  связан с единицей продукции
type ProducedInfo struct {
	Terminal string    `bson:"" json:",omitempty"` // Имя линии фасовки, где он был нанесен или считан камерой
	Time     time.Time `bson:"" json:",omitempty"` // Время, когда код был нанесен или считан на линии
	ProdDate string    `bson:"" json:",omitempty"` // Дата производства продукта 2023-10-09
	Discard  bool      `bson:"" json:",omitempty"` // True - операция отбраковки кода
}

// Информация, связанная с печатью
type PrintInfo struct {
	PrintId uint64 `bson:"" json:",omitempty"` // Уникальный номер для кода.
}

// Информация о выгрузке в 1с
type UploadInfo struct {
	Time   time.Time `bson:"" json:",omitempty"`
	Status string    `bson:"" json:",omitempty"`
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
