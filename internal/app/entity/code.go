package entity

import (
	"errors"
	"time"
)

type Code struct {
	Gtin   string
	Serial string
	Crypto string
}

type CodeForPrint struct {
	Code    Code
	PrintId uint64
}

type FullCode struct {
	Code
	SourceInfo   SourceInfo
	PrintInfo    PrintInfo
	ProducedInfo []ProducedInfo
	UploadInfo   UploadInfo
}

func (code *Code) Validate() error {
	err := ValidateGtin(code.Gtin)
	if err != nil {
		return err
	}

	err = ValidateSerial(code.Serial)
	if err != nil {
		return err
	}

	err = ValidateCrypto(code.Crypto)
	if err != nil {
		return err
	}

	return nil
}

// Когда и откуда был загружен код
type SourceInfo struct {
	Name string    // Откуда загружен, например с сервера "server main"
	Time time.Time // Время получения кода
}

// Когда и где код пошел в выпуск продукции. т.е. был  связан с единицей продукции
type ProducedInfo struct {
	Terminal string    // Имя линии фасовки, где он был нанесен или считан камерой
	Time     time.Time // Время, когда код был нанесен или считан на линии
	ProdDate string    // Дата производства продукта 2023-10-09
	Discard  bool      // True - операция отбраковки кода
}

// Информация, связанная с печатью
type PrintInfo struct {
	Avaible      bool      // Флаг, что код доступен для печати
	UploadTime   time.Time // Время выдачи кода из базы
	TerminalName string    // Имя линии, куда передан код
	PrintID      uint32    // Уникальный номер для кода, присваивается при выдаче кода из БД
}

// Информация о выгрузке во внешнюю систему
type UploadInfo struct {
	Time   time.Time
	Status string
}

func ValidateSerial(serial string) error {
	if len(serial) != 6 {
		return errors.New("некорректная длинна serial")
	}
	return nil
}

func ValidateCrypto(crypto string) error {
	if len(crypto) != 4 {
		return errors.New("некорректная длинна crypto")
	}
	return nil
}
