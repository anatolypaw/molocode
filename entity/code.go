package entity

import (
	"fmt"
	"time"
)

type Code struct {
	Serial       string         `bson:"_id" json:"serial"`                  // Серийный номер, формат честного знака. Уникален для каждого кода с этим GTIN
	Crypto       string         `bson:"crypto" json:"crypto"`               // Криптохвост, формат честного знака
	SourceInfo   SourceInfo     `bson:"source_info" json:"source_info"`     // Информация об источнике поступления коданформация о линии, которая получила код
	PrintInfo    PrintInfo      `bson:"print_info" json:"print_info"`       // Информация, связанная с печатью
	ProducedInfo []ProducedInfo `bson:"produced_info" json:"produced_info"` // Информация о его выпуске на линии фасовки
	UploadInfo   UploadInfo     `bson:"upload_info" json:"upload_info"`     // Информация о выгрузке в 1с
}

// Когда и откуда был загружен код
type SourceInfo struct {
	Name string    `bson:"name" json:"name"` // Откуда загружен, например с сервера "server main"
	Time time.Time `bson:"time" json:"time"` // Время получения кода
}

// Когда и где код пошел в выпуск продукции. т.е. был  связан с единицей продукции
type ProducedInfo struct {
	Terminal string    `json:"terminal"`  // Имя линии фасовки, где он был нанесен или считан камерой
	Time     time.Time `json:"time"`      // Время, когда код был нанесен или считан на линии
	ProdDate string    `json:"prod_date"` // Дата производства продукта 2023-10-09
	Discard  bool      `json:"discard"`   // True - операция отбраковки кода
}

// Информация, связанная с печатью
type PrintInfo struct {
	Avaible    bool      `bson:"avaible" json:"avaible"`         // Флаг, что код доступен для печати
	UploadTime time.Time `bson:"upload_time" json:"upload_time"` // Время выдачи кода из базы
	Terminal   string    `bson:"terminal" json:"terminal"`       // Имя линии, куда передан код
	PrintID    uint32    `bson:"print_id" json:"print_id"`       // Уникальный номер для кода, присваивается при выдаче кода из БД
}

// Информация о выгрузке в 1с
type UploadInfo struct {
	Time   time.Time `bson:"time" json:"time"`
	Status string    `bson:"status" json:"status"`
}

func ValidateSerial(serial string) error {
	const op = "entity.Code.ValidateSerial"
	if len(serial) != 6 {
		return fmt.Errorf("%s: Некорректная длинна серийного номера", op)
	}

	return nil
}

func ValidateCrypto(crypto string) error {
	const op = "entity.Code.ValidateCrypto"
	if len(crypto) != 4 {
		return fmt.Errorf("%s: Некорректная длинна криптохвоста", op)
	}

	return nil
}

// Проверяют корректность серийного номера, криптохвоста
func ValidateSerialCrypto(serial, crypto string) error {
	if err := ValidateSerial(serial); err != nil {
		return err
	}

	if err := ValidateCrypto(crypto); err != nil {
		return err
	}
	return nil
}

func (code *Code) ValidateSerialCrypto() error {
	return ValidateSerialCrypto(code.Serial, code.Crypto)
}
