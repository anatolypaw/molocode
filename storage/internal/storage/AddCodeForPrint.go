package storage

import (
	"context"
	"fmt"
	"storage/entity"
	"time"
)

// Добавляет код к указанному по gtin продукту для последующей печати
// Добавляет код, только если в свойствах этого продукта разрешено получение кодов для нанесения
func (con *Storage) AddCodeForPrint(gtin, serial, crypto, sourceName string) error {
	const op = "storage.AddCodeForPrint"

	// Валидируем данные о коде
	if err := entity.ValidateSerialCrypto(serial, crypto); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем, существует ли такой продукт в БД
	good, err := con.GetGood(gtin)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем, разрешено ли получение кодов для этого продукта
	if !good.GetCodeForPrint {
		return fmt.Errorf("%s: %s", op, "Для этого продукта запрещено получение кодов для нанесения")
	}

	// TODO проверять, не превышено ли количество требуемых кодов для этого продукта

	var newCode entity.Code
	newCode.Serial = serial
	newCode.Crypto = crypto
	newCode.SourceInfo.Name = sourceName
	newCode.SourceInfo.Time = time.Now()
	newCode.PrintInfo.Avaible = true
	newCode.ProducedInfo = []entity.ProducedInfo{}

	_, err = con.db.Collection(gtin).InsertOne(context.TODO(), newCode)
	if err != nil {
		return err
	}

	return err
}
