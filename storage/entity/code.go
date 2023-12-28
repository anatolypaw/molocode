package entity

import (
	"fmt"
	"time"
)

type Code struct {
	Serial       string         `bson:"_id"`          // Серийный номер, формат честного знака. Уникален для каждого кода с этим GTIN
	Crypto       string         `bson:"Crypto"`       // Криптохвост, формат честного знака
	SourceInfo   SourceInfo     `bson:"Sourceinfo"`   // Информация об источнике поступления коданформация о линии, которая получила код
	PrintInfo    PrintInfo      `bson:"PrintInfo"`    // Информация, связанная с печатью
	ProducedInfo []ProducedInfo `bson:"ProducedInfo"` // Информация о его выпуске на линии фасовки
	UploadInfo   UploadInfo     `bson:"UploadInfo"`   // Информация о выгрузке в 1с
}

// Когда и откуда был загружен код
type SourceInfo struct {
	Name string    `bson:"Name"` // Откуда загружен, например с сервера "server main"
	Time time.Time `bson:"Time"` // Время получения кода
}

// Когда и где код пошел в выпуск продукции. т.е. был  связан с единицей продукции
type ProducedInfo struct {
	Terminal string    `bson:"Terminal"` // Имя линии фасовки, где он был нанесен или считан камерой
	Time     time.Time `bson:"Time"`     // Время, когда код был нанесен или считан на линии
	ProdDate string    `bson:"ProdDate"` // Дата производства продукта 2023-10-09
	Discard  bool      `bson:"Discard"`  // True - операция отбраковки кода
}

// Информация, связанная с печатью
type PrintInfo struct {
	Avaible    bool      `bson:"Avaible"`    // Флаг, что код доступен для печати
	UploadTime time.Time `bson:"UploadTime"` // Время выдачи кода из базы
	Terminal   string    `bson:"Terminal"`   // Имя линии, куда передан код
	PrintID    uint32    `bson:"PrintID"`    // Уникальный номер для кода, присваивается при выдаче кода из БД
}

// Информация о выгрузке в 1с
type UploadInfo struct {
	Time   time.Time `bson:"Time"`
	Status string    `bson:"Status"`
}

func ValidateSerial(serial string) error {
	const op = "model.Code.ValidateSerial"
	if len(serial) != 6 {
		return fmt.Errorf("%s: Некорректная длинна серийного номера", op)
	}

	return nil
}

func ValidateCrypto(crypto string) error {
	const op = "model.Code.ValidateCrypto"
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
